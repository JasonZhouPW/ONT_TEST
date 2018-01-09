using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System;
using System.Numerics;

namespace ONT_DEx
{
    public class Ont_Fund : SmartContract
    {
        public static readonly byte[] AssetIdKey = { 97, 115, 115, 101, 116, 105, 100 };//assetid
        public static readonly byte[] FundCallerKey = { 102, 117, 110, 100, 99, 97, 108, 108, 101, 114 };//fundcaller
        public static readonly byte[] FundAdminKey = { 102, 117, 110, 100, 97, 100, 109, 105, 110 };//fundadmin
        public static readonly byte[] TotalBalancePrefix = { 116, 111, 116, 97, 108 };//total
        public static readonly byte[] AvailBalancePrefix = { 97, 118, 97, 105, 108 };//avail
        //Error code
        public static readonly int Error_NO = 0;//success
        public static readonly int Error_TriggerType = 1000;
        public static readonly int Error_ParamInvalidate = 1001;
        public static readonly int Error_UnknowOperation = 1002;
        public static readonly int Error_VerifyInvalidate = 1003;
        public static readonly int Error_CallerInvalidate = 1004;
        public static readonly int Error_BalanceNoEnough = 1005;
        public static readonly int Error_LockBalanceInvalidate = 1006;
        public static readonly int Error_AssetInvalidate = 1007;
        public static readonly int Error_WitnessInvalidate = 1008;
        public static readonly int Error_UTXOInvalidate = 1009;
        public static readonly int Error_DulicateInit = 1010;
        public static readonly int Error_NotInit = 1011;

        public static object[] Main(string operation, params object[] args)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (Runtime.Trigger == TriggerType.Verification)
            {
                ret[0] = Error_TriggerType;
                return ret;
            }
            else if (Runtime.Trigger == TriggerType.Application)
            {
                if (operation == "init")
                {
                    if (args.Length != 3)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] assetId = (byte[])args[0];
                    byte[] admin = (byte[])args[1];
                    byte[] caller = (byte[])args[2];
                    return Init(assetId, admin, caller);
                }
                if (operation == "changeadmin")
                {
                    if (args.Length != 1)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] newAdmin = (byte[])args[0];
                    return ChangeAdmin(newAdmin);
                }
                if (operation == "getadmin")
                {
                    return GetAdmin();
                }
                if (operation == "setcaller")
                {
                    if (args.Length != 1)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = (byte[])args[0];
                    return SetCaller(caller);
                }
                if (operation == "getcaller")
                {
                    return GetCaller();
                }
                if (operation == "deposit")
                {
                    return Deposit();
                }
                if (operation == "lock")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = ExecutionEngine.CallingScriptHash;
                    byte[] buyer = (byte[])args[0];
                    BigInteger amount = (BigInteger)args[1];
                    return Lock(caller, buyer, amount);
                }
                if (operation == "unlock")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = ExecutionEngine.CallingScriptHash;
                    byte[] buyer = (byte[])args[0];
                    BigInteger amount = (BigInteger)args[1];
                    return Unlock(caller, buyer, amount);
                }
                if (operation == "balanceof")
                {
                    if (args.Length != 1)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] address = (byte[])args[0];
                    return BalanceOf(address);
                }
                if (operation == "receipt")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = ExecutionEngine.CallingScriptHash;
                    byte[] reciver = (byte[])args[0];
                    BigInteger amount = (BigInteger)args[1];
                    return Receipt(caller, reciver, amount);
                }
                if (operation == "payment")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = ExecutionEngine.CallingScriptHash;
                    byte[] payer = (byte[])args[0];
                    BigInteger amount = (BigInteger)args[1];
                    return Payment(caller, payer, amount);
                }
                if (operation == "checkcallerpermisssion")
                {
                    if (args.Length != 1)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] caller = (byte[])args[0];
                    return CheckCallerPermission(caller);
                }
            }

            ret[0] = Error_UnknowOperation;
            return ret;
        }

        //部署
        public static object[] Init(byte[] assetId, byte[] admin, byte[] caller)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            //Only once
            byte[] oldAdmin = Storage.Get(Storage.CurrentContext, FundAdminKey);
            if (oldAdmin.Length != 0)
            {
                ret[0] = Error_DulicateInit;
                return ret;
            }
            byte[] zero = { 0 };
            if (caller != zero)
            {
                Storage.Put(Storage.CurrentContext, FundCallerKey, caller);
            }
            Storage.Put(Storage.CurrentContext, FundAdminKey, admin);
            Storage.Put(Storage.CurrentContext, AssetIdKey, assetId);
            return ret;
        }

        //修改管理员
        public static object[] ChangeAdmin(byte[] newAmdin)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] oldAdmin = Storage.Get(Storage.CurrentContext, FundAdminKey);
            if (oldAdmin.Length == 0)
            {
                ret[0]=Error_NotInit;
                return ret;
            }
            //Only admin can change
            if (!Runtime.CheckWitness(oldAdmin))
            {
                ret[0]= Error_WitnessInvalidate;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, FundAdminKey, newAmdin);
            return ret;
        }

        public static object[] GetAdmin()
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = Storage.Get(Storage.CurrentContext, FundAdminKey);
            return ret;
        }

        //设置合约调用方
        public static object[] SetCaller(byte[] caller)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] admin = Storage.Get(Storage.CurrentContext, FundAdminKey);
            if (admin.Length == 0)
            {
                ret[0] = Error_NotInit;
                return ret;
            }
            if (!Runtime.CheckWitness(admin))
            {
                ret[0] = Error_WitnessInvalidate;
            }
            Storage.Put(Storage.CurrentContext, FundCallerKey, caller);
            return ret;
        }

        public static object[] GetCaller()
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = Storage.Get(Storage.CurrentContext, FundCallerKey);
            return ret;
        }

        //入金
        public static object[] Deposit()
        {
            object[] ret = new object[3];
            ret[0] = Error_NO;
            byte[] assetId = Storage.Get(Storage.CurrentContext, AssetIdKey);
            if (assetId.Length == 0)
            {
                ret[0] = Error_NotInit;
                return ret;
            }

            Transaction tx = (Transaction)ExecutionEngine.ScriptContainer;
            TransactionOutput reference = tx.GetReferences()[0];
            if (reference.AssetId != assetId)
            {
                ret[0] = Error_AssetInvalidate;
                return ret;
            }
            byte[] sender = reference.ScriptHash;
            byte[] receiver = ExecutionEngine.ExecutingScriptHash;
            TransactionOutput[] outputs = tx.GetOutputs();
            BigInteger amount = 0;

            foreach (TransactionOutput output in outputs)
            {
                if (receiver.Equals(output.ScriptHash))
                {
                    amount += (BigInteger)output.Value;
                }
            }
            if (amount <= 0)
            {
                ret[0] = Error_UTXOInvalidate;
                return ret;
            }
            byte[] totalBalanceKey = getKey(TotalBalancePrefix, sender);
            byte[] availBalanceKey = getKey(AvailBalancePrefix, sender);
            BigInteger totalBalance = Storage.Get(Storage.CurrentContext, totalBalanceKey).AsBigInteger();
            BigInteger availBalance = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            Storage.Put(Storage.CurrentContext, totalBalanceKey, (totalBalance + amount).AsByteArray());
            Storage.Put(Storage.CurrentContext, availBalanceKey, (availBalance + amount).AsByteArray());
            ret[1] = availBalance + amount;
            ret[2] = totalBalance + amount;
            return ret;
        }

        public static object[] Receipt(byte[] caller, byte[] reciver, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (amount <= 0)
            {
                ret[0] = Error_ParamInvalidate;
                return ret;
            }

            int errorCode = checkCallerPermission(caller);
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            byte[] totalBalanceKey = getKey(TotalBalancePrefix, reciver);
            byte[] availBalanceKey = getKey(AvailBalancePrefix, reciver);
            BigInteger totalBalance = Storage.Get(Storage.CurrentContext, totalBalanceKey).AsBigInteger();
            BigInteger availBalance = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            Storage.Put(Storage.CurrentContext, totalBalanceKey, (totalBalance + amount).AsByteArray());
            Storage.Put(Storage.CurrentContext, availBalanceKey, (availBalance + amount).AsByteArray());
            return ret;
        }

        public static object[] Payment(byte[] caller, byte[] payer, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (amount <= 0)
            {
                ret[0] = Error_ParamInvalidate;
                return ret;
            }
            int errorCode = checkCallerPermission(caller);
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }

            byte[] availBalanceKey = getKey(AvailBalancePrefix, payer);
            BigInteger availBalance = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            if (availBalance < amount)
            {
                ret[0] = Error_BalanceNoEnough;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, availBalanceKey, (availBalance - amount).AsByteArray());

            byte[] totalBalanceKey = getKey(TotalBalancePrefix, payer);
            BigInteger totalBalance = Storage.Get(Storage.CurrentContext, totalBalanceKey).AsBigInteger();
            Storage.Put(Storage.CurrentContext, totalBalanceKey, (totalBalance - amount).AsByteArray());
            return ret;
        }

        //锁仓
        public static object[] Lock(byte[] caller, byte[] buyer, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (amount <= 0)
            {
                ret[0] = Error_ParamInvalidate;
                return ret;
            }

            int errorCode = checkCallerPermission(caller);
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            byte[] availBalanceKey = getKey(AvailBalancePrefix, buyer);
            BigInteger availBalance = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            if (availBalance < amount)
            {
                ret[0] = Error_BalanceNoEnough;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, availBalanceKey, (availBalance - amount).AsByteArray());
            return ret;
        }

        //解锁
        public static object[] Unlock(byte[] caller, byte[] buyer, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (amount <= 0)
            {
                ret[0] = Error_ParamInvalidate;
                return ret;
            }

            int errorCode = checkCallerPermission(caller);
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            byte[] totalBalanceKey = getKey(TotalBalancePrefix, buyer);
            BigInteger totalBalance = Storage.Get(Storage.CurrentContext, totalBalanceKey).AsBigInteger();
            byte[] availBalanceKey = getKey(AvailBalancePrefix, buyer);
            BigInteger availBalance = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            if (totalBalance - availBalance < amount)
            {
                //已经解锁过了
                ret[0] = Error_LockBalanceInvalidate;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, availBalanceKey, (availBalance + amount).AsByteArray());
            return ret;
        }

        //总余额
        public static object[] BalanceOf(byte[] address)
        {
            object[] ret = new object[3];
            ret[0] = Error_NO;
            if (!Runtime.CheckWitness(address))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            byte[] availBalanceKey = getKey(AvailBalancePrefix, address);
            ret[1] = Storage.Get(Storage.CurrentContext, availBalanceKey).AsBigInteger();
            byte[] totalBalanceKey = getKey(TotalBalancePrefix, address);
            ret[2] = Storage.Get(Storage.CurrentContext, totalBalanceKey).AsBigInteger();
            return ret;
        }

        public static object[] CheckCallerPermission(byte[] caller)
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            byte[] c = Storage.Get(Storage.CurrentContext, FundCallerKey);
            if (c.Length == 0)
            {
                ret[1] = true;
                return ret;
            }
            ret[1] = (c == caller);
            return ret;
        }

        private static int checkCallerPermission(byte[] caller)
        {
            byte[] c = Storage.Get(Storage.CurrentContext, FundCallerKey);
            if (c.Length == 0)
            {
                return Error_NO;
            }

            if (caller != c)
            {
                return Error_CallerInvalidate;
            }
            return Error_NO;
        }

        private static byte[] getKey(byte[] prefix, byte[] key)
        {
            return prefix.Concat(key);
        }
    }
}
