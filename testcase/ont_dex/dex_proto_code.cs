using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System;
using System.Numerics;

namespace ONT_DEx
{
    public class Ont_Proto : SmartContract
    {
        [Appcall("7f414d02bcecba8af3daaec395181870ab358c72")]
        public static extern object[] Ont_Fund(string operation, object[] args);

        public static readonly byte[] ProtoAdminKey = { 112, 114, 111, 116, 111, 97, 100, 109, 105, 110 };
        public static readonly byte[] ProtoCallerPrefix = { 99, 97, 108, 108, 101, 114 };

        //Error code
        public static readonly int Error_NO = 0;//success
        public static readonly int Error_TriggerType = 2000;
        public static readonly int Error_ParamInvalidate = 2001;
        public static readonly int Error_UnknowOperation = 2002;
        public static readonly int Error_VerifyInvalidate = 2003;
        public static readonly int Error_CallerInvalidate = 2004;
        public static readonly int Error_WitnessInvalidate = 2007;
        public static readonly int Error_DuplicateInit = 2009;
        public static readonly int Error_NotInit = 2010;

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
                    if (args.Length != 2) {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] admin = (byte[])args[0];
                    byte[] caller = (byte[])args[1];
                    return Init(admin,caller);
                }
                if (operation == "onmakeorder")
                {
                    if (args.Length != 3)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] buyer = (byte[])args[0];
                    byte[] seller = (byte[])args[1];
                    BigInteger amount = (BigInteger)args[2];
                    return OnMakeOrder(buyer, seller, amount);
                }
                if (operation == "onordercomplete")
                {
                    if (args.Length != 3)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] buyer = (byte[])args[0];
                    byte[] seller = (byte[])args[1];
                    BigInteger amount = (BigInteger)args[2];
                    return OnOrderComplete(buyer, seller, amount);
                }
                if (operation == "onordercancel")
                {
                    if (args.Length != 3)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] buyer = (byte[])args[0];
                    byte[] seller = (byte[])args[1];
                    BigInteger amount = (BigInteger)args[2];
                    return OnOrderCancel(buyer, seller, amount);
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
                if (operation == "addcaller")
                {
                    if (args.Length != 1)
                {
                    ret[0] = Error_ParamInvalidate;
                    return ret;
                }
                byte[] caller = (byte[])args[0];
                    return AddCaller(caller);
                }
                if (operation == "deletecaller")
                {
                    if (args.Length != 1)
                {
                    ret[0] = Error_ParamInvalidate;
                    return ret;
                }
                byte[] caller = (byte[])args[0];
                    return DeleteCaller(caller);
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

        public static object[] Init(byte[] admin,byte[] caller)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] oldAdmin = (byte[])Storage.Get(Storage.CurrentContext, ProtoAdminKey);
            if (oldAdmin.Length != 0)
            {
                ret[0] = Error_DuplicateInit;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, ProtoAdminKey, admin);
            byte[] zero = { 0 };
            if (caller == zero)
            {
                return ret;
            }
            byte[] callerKey = getKey(ProtoCallerPrefix, caller);
            Storage.Put(Storage.CurrentContext, callerKey, caller);
            return ret;
        }

        public static object[] OnMakeOrder(byte[] buyer, byte[] seller, BigInteger amount)
        {
            object[] ret = new object[3];
            ret[0] = Error_NO;
            byte[] caller = ExecutionEngine.CallingScriptHash;
            ret[1] = caller;
            ret[2] = ExecutionEngine.ExecutingScriptHash;
            if (!checkCallerPermission(caller))
            {
                ret[0] = Error_CallerInvalidate;
                return ret;
            }
            object[] param = { buyer, amount };
            object[] retf = Ont_Fund("lock", param);
            int errorCode = (int)retf[0];
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
            }
            return ret;
        }

        public static object[] OnOrderComplete( byte[] buyer, byte[] seller, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] caller = ExecutionEngine.CallingScriptHash;
            if (!checkCallerPermission(caller))
            {
                ret[0] = Error_CallerInvalidate;
                return ret;
            }
            object[] param1 = { buyer, amount };
            object[] retf = Ont_Fund("unlock", param1);
            int errorCode = (int)retf[0];
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            object[] param2 = { buyer, amount };
            retf = Ont_Fund("payment", param2);
            errorCode = (int)retf[0];
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            object[] param3 = { seller, amount };
            retf = Ont_Fund("receipt", param3);
            errorCode = (int)retf[0];
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            return ret;
        }

        public static object[] OnOrderCancel(byte[] buyer, byte[] seller, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;

            byte[] caller = ExecutionEngine.CallingScriptHash;
            if (!checkCallerPermission(caller))
            {
                ret[0] = Error_CallerInvalidate;
                return ret;
            }
            object[] param = { buyer, amount };
            object[] retf = Ont_Fund("unlock", param);
            int errorCode = (int)retf[0];
            if (errorCode != Error_NO)
            {
                ret[0] = errorCode;
                return ret;
            }
            return ret;
        }

        public static object[] AddCaller(byte[] caller)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;

            byte[] admin = Storage.Get(Storage.CurrentContext, ProtoAdminKey);
            if (admin.Length == 0)
            {
                ret[0] = Error_NotInit;
                return ret;
            }
            if (!Runtime.CheckWitness(admin))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }

            byte[] callerKey = getKey(ProtoCallerPrefix, caller);
            Storage.Put(Storage.CurrentContext, callerKey, caller);
            return ret;
        }

        public static object[] DeleteCaller(byte[] caller)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] admin = Storage.Get(Storage.CurrentContext, ProtoAdminKey);
            if (admin.Length == 0)
            {
                ret[0] = Error_NotInit;
                return ret;
            }
            if (!Runtime.CheckWitness(admin))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            byte[] callerKey = getKey(ProtoCallerPrefix, caller);
            Storage.Delete(Storage.CurrentContext, callerKey);
            return ret;
        }

        public static object[] CheckCallerPermission(byte[] caller)
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = checkCallerPermission(caller);
            return ret;
        }

        private static bool checkCallerPermission(byte[] caller)
        {
            byte[] callerKey = getKey(ProtoCallerPrefix, caller);
            byte[] c = Storage.Get(Storage.CurrentContext, callerKey);
            return c.Length != 0;
        }

        //修改管理员
        public static object[] ChangeAdmin(byte[] newAmdin)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;

            byte[] oldAdmin = Storage.Get(Storage.CurrentContext, ProtoAdminKey);
            if (oldAdmin.Length == 0)
            {
                ret[0] = Error_NotInit;
                return ret;
            }
            //Only admin can change
            if (!Runtime.CheckWitness(oldAdmin))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, ProtoAdminKey, newAmdin);
            return ret;
        }

        public static object[] GetAdmin()
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = Storage.Get(Storage.CurrentContext, ProtoAdminKey);
            return ret;
        }

        private static byte[] getKey(byte[] prefix, byte[] key)
        {
            return prefix.Concat(key);
        }
    }
}
