using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System;
using System.Numerics;

namespace ONT_DEx
{
    public class ONT_P2P : SmartContract
    {
        public static readonly byte[] OrderLockTimeKey = { 111, 114, 100, 101, 114, 108, 111, 99, 107, 116, 105, 109, 101 };//orderlocktime
        public static readonly byte[] P2PAdminKey = { 112, 50, 112, 97, 100, 109, 105, 110 };//p2padmin
        public static readonly byte[] OrderAmountPrefix = { 111, 97, 109, 111, 117, 110, 116 }; //oamount
        public static readonly byte[] OrderSellerPrefix = { 111, 115, 101, 108, 108, 101, 114 };//oseller
        public static readonly byte[] OrderBuyerPrefix = { 111, 98, 117, 121, 101, 114 };//obuyer
        public static readonly byte[] OrderTimePrefix = { 111, 116, 105, 109, 101 };//otime

        //Error code
        public static readonly int Error_NO = 0;//success
        public static readonly int Error_TriggerType = 3000;
        public static readonly int Error_ParamInvalidate = 3001;
        public static readonly int Error_UnknowOperation = 3002;
        public static readonly int Error_VerifyInvalidate = 3003;
        public static readonly int Error_DuplicateOrder = 3004;
        public static readonly int Error_InvalidateOrder = 3005;
        public static readonly int Error_NoUnlockTime = 3006;
        public static readonly int Error_InvalidateOperator = 3007;
        public static readonly int Error_WitnessInvalidate = 3008;
        public static readonly int Error_DuplicateInit = 3009;
        public static readonly int Error_UnInit = 3010;

        [Appcall("d06bb85d525d977d0a9702a08f841528723df0e9")]
        public static extern object[] Ont_Proto(string operation, object[] args);

        public static object[] Main(string operation, params object[] args)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (Runtime.Trigger == TriggerType.Verification)
            {
                ret[0] = Error_TriggerType ;
                return ret;
            }
            else if (Runtime.Trigger == TriggerType.Application)
            {
                if (operation == "init")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] admin = (byte[])args[0];
                    BigInteger lockTime = (BigInteger)args[1];
                    return Init(admin, lockTime);
                }
                if (operation == "makebuyorder")
                {
                    if (args.Length != 6)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] orderSig = (byte[])args[0];
                    byte[] orderId = (byte[])args[1];
                    byte[] buyer = (byte[])args[2];
                    byte[] seller = (byte[])args[3];
                    byte[] buyerPk = (byte[])args[4];
                    BigInteger amount = (BigInteger)args[5];
                    return MakeBuyOrder(orderSig, orderId, buyer, seller, buyerPk, amount);
                }
                if (operation == "buyordercomplete")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] orderId = (byte[])args[0];
                    byte[] buyer = (byte[])args[1];
                    return BuyOrderComplete(orderId, buyer);
                }
                if (operation == "buyordercancel")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] orderId = (byte[])args[0];
                    byte[] buyer = (byte[])args[1];
                    return BuyOrderCancel(orderId, buyer);
                }
                if (operation == "sellertrycloseorder")
                {
                    if (args.Length != 2)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    byte[] orderId = (byte[])args[0];
                    byte[] seller = (byte[])args[1];
                    return SellerTryCloseOrder(orderId, seller);
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
                if (operation == "setorderlocktime")
                {
                    if (args.Length != 1)
                    {
                        ret[0] = Error_ParamInvalidate;
                        return ret;
                    }
                    BigInteger lTime = (BigInteger)args[0];
                    return SetOrderLockTime(lTime);
                }
                if (operation == "getorderlocktime")
                {
                    return GetOrderLockTime();
                }
            }
            ret[0] = Error_UnknowOperation;
            return ret;
        }

        public static object[] Init(byte[] admin, BigInteger lockTime)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            //Only once
            byte[] oldAdmin = Storage.Get(Storage.CurrentContext, P2PAdminKey);
            if (oldAdmin.Length != 0)
            {
                ret[0] = Error_DuplicateInit;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, P2PAdminKey, admin);
            Storage.Put(Storage.CurrentContext, OrderLockTimeKey, lockTime.AsByteArray());
            return ret;
        }

        public static object[] MakeBuyOrder(byte[] orderSig, byte[] orderId, byte[] buyer, byte[] seller, byte[] buyerPk, BigInteger amount)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            //if VerifySignature(orderSig, buyerPk)
            //{
            //    return Error_VerifyInvalidate;
            //}
            if (!Runtime.CheckWitness(seller))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }

            Header header = Blockchain.GetHeader(Blockchain.GetHeight());
            BigInteger orderTime = (BigInteger)header.Timestamp;

            byte[] oBuyerKey = getKey(OrderBuyerPrefix, orderId);
            byte[] oBuyer = Storage.Get(Storage.CurrentContext, oBuyerKey);
            if (oBuyer.Length != 0)
            {
                //不能重复创建订单
                ret[0] = Error_DuplicateOrder;
                return ret;
            }

            byte[] oAmountKey = getKey(OrderAmountPrefix, orderId);
            byte[] oSellerKey = getKey(OrderSellerPrefix, orderId);
            byte[] oTimeKey = getKey(OrderTimePrefix, orderId);

            Storage.Put(Storage.CurrentContext, oAmountKey, amount.AsByteArray());
            Storage.Put(Storage.CurrentContext, oBuyerKey, buyer);
            Storage.Put(Storage.CurrentContext, oSellerKey, seller);
            Storage.Put(Storage.CurrentContext, oTimeKey, orderTime.AsByteArray());

            //object[] param = { buyer, seller, amount };
            //object[] retp = Ont_Proto("onmakeorder", param);
            //int errorCode = (int)retp[0];
            //if (errorCode != Error_NO)
            //{
            //    return retp;
            //}
            return ret;
        }

        public static object[] BuyOrderComplete(byte[] orderId, byte[] buyer)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (!Runtime.CheckWitness(buyer))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            byte[] oBuyerKey = getKey(OrderBuyerPrefix, orderId);
            byte[] oBuyer = Storage.Get(Storage.CurrentContext, oBuyerKey);
            if (oBuyer.Length == 0)
            {
                ret[0] = Error_InvalidateOrder;
                return ret;
            }
            if (oBuyer != buyer)
            {
                ret[0] = Error_InvalidateOperator;
                return ret;
            }
            byte[] oAmountKey = getKey(OrderAmountPrefix, orderId);
            BigInteger oAmount = Storage.Get(Storage.CurrentContext, oAmountKey).AsBigInteger();

            byte[] oSellerKey = getKey(OrderSellerPrefix, orderId);
            byte[] oSeller = Storage.Get(Storage.CurrentContext, oSellerKey);

            object[] param = { oBuyer, oSeller, oAmount };
            object[] retp = Ont_Proto("onordercomplete", param);
            int errorCode = (int)retp[0];
            if (errorCode != Error_NO)
            {
                return retp;
            }

            byte[] oTimeKey = getKey(OrderTimePrefix, orderId);

            //删除买单
            Storage.Delete(Storage.CurrentContext, oBuyerKey);
            Storage.Delete(Storage.CurrentContext, oSellerKey);
            Storage.Delete(Storage.CurrentContext, oAmountKey);
            Storage.Delete(Storage.CurrentContext, oTimeKey);
            return ret;
        }

        public static object[] BuyOrderCancel(byte[] orderId, byte[] buyer)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (!Runtime.CheckWitness(buyer))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            byte[] oBuyerKey = getKey(OrderBuyerPrefix, orderId);
            byte[] oBuyer = Storage.Get(Storage.CurrentContext, oBuyerKey);
            if (oBuyer.Length == 0)
            {
                ret[0] = Error_InvalidateOrder;
                return ret;
            }
            if (oBuyer != buyer)
            {
                ret[0] = Error_InvalidateOperator;
                return ret;
            }
            byte[] oAmountKey = getKey(OrderAmountPrefix, orderId);
            BigInteger oAmount = Storage.Get(Storage.CurrentContext, oAmountKey).AsBigInteger();

            byte[] oSellerKey = getKey(OrderSellerPrefix, orderId);
            byte[] oSeller = Storage.Get(Storage.CurrentContext, oSellerKey);
            object[] param = { oBuyer, oSeller, oAmount };
            object[] retp = Ont_Proto("onordercancel", param);
            int errorCode = (int)retp[0];
            if (errorCode != Error_NO)
            {
                return retp;
            }

            byte[] oTimeKey = getKey(OrderTimePrefix, orderId);

            //删除买单
            Storage.Delete(Storage.CurrentContext, oBuyerKey);
            Storage.Delete(Storage.CurrentContext, oSellerKey);
            Storage.Delete(Storage.CurrentContext, oAmountKey);
            Storage.Delete(Storage.CurrentContext, oTimeKey);
            return ret;
        }

        public static object[] SellerTryCloseOrder(byte[] orderId, byte[] seller)
        {
            object[] ret = new object[3];
            ret[0] = Error_NO;
            if (!Runtime.CheckWitness(seller))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            byte[] oBuyerKey = getKey(OrderBuyerPrefix, orderId);
            byte[] oBuyer = Storage.Get(Storage.CurrentContext, oBuyerKey);
            if (oBuyer.Length == 0)
            {
                ret[0] = Error_InvalidateOrder;
                return ret;
            }
            byte[] oSellerKey = getKey(OrderSellerPrefix, orderId);
            byte[] oSeller = Storage.Get(Storage.CurrentContext, oSellerKey);
            if (oSeller != seller)
            {
                ret[0] = Error_InvalidateOperator;
                return ret;
            }

            byte[] oTimeKey = getKey(OrderTimePrefix, orderId);
            BigInteger oTime = Storage.Get(Storage.CurrentContext, oTimeKey).AsBigInteger();
            Header header = Blockchain.GetHeader(Blockchain.GetHeight());
            BigInteger timeNow = (BigInteger)header.Timestamp;
            BigInteger lockTime = (BigInteger)GetOrderLockTime()[1];
            ret[1] = timeNow;
            ret[2] = lockTime;
            if ((timeNow - oTime) < lockTime)
            {
                //还没到解锁时间
                ret[0] = Error_NoUnlockTime;
                return ret;
            }

            byte[] oAmountKey = getKey(OrderAmountPrefix, orderId);
            BigInteger oAmount = Storage.Get(Storage.CurrentContext, oAmountKey).AsBigInteger();

            object[] param = { oBuyer, oSeller, oAmount };
            object[] retp = Ont_Proto("onordercomplete", param);
            int errorCode = (int)retp[0];
            if (errorCode != Error_NO)
            {
                return retp;
            }

            //删除买单
            Storage.Delete(Storage.CurrentContext, oBuyerKey);
            Storage.Delete(Storage.CurrentContext, oSellerKey);
            Storage.Delete(Storage.CurrentContext, oAmountKey);
            Storage.Delete(Storage.CurrentContext, oTimeKey);
            return ret;
        }

        //修改管理员
        public static object[] ChangeAdmin(byte[] newAmdin)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            byte[] oldAdmin = Storage.Get(Storage.CurrentContext, P2PAdminKey);
            if (oldAdmin.Length == 0)
            {
                ret[0] = Error_UnInit;
                return ret;
            }
            //Only admin can change
            if (!Runtime.CheckWitness(oldAdmin))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, P2PAdminKey, newAmdin);
            return ret;
        }

        public static object[] GetAdmin()
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = Storage.Get(Storage.CurrentContext, P2PAdminKey);
            return ret;
        }

        public static object[] SetOrderLockTime(BigInteger lTime)
        {
            object[] ret = new object[1];
            ret[0] = Error_NO;
            if (lTime < 0)
            {
                ret[0] = Error_ParamInvalidate;
                return ret;
            }
            byte[] admin = Storage.Get(Storage.CurrentContext, P2PAdminKey);
            if (admin.Length == 0)
            {
                ret[0] = Error_UnInit;
                return ret;
            }
            if (!Runtime.CheckWitness(admin))
            {
                ret[0] = Error_WitnessInvalidate;
                return ret;
            }
            Storage.Put(Storage.CurrentContext, OrderLockTimeKey, lTime.AsByteArray());
            return ret;
        }

        public static object[] GetOrderLockTime()
        {
            object[] ret = new object[2];
            ret[0] = Error_NO;
            ret[1] = Storage.Get(Storage.CurrentContext, OrderLockTimeKey).AsBigInteger();
            return ret;
        }

        private static byte[] getKey(byte[] prefix, byte[] key)
        {
            return prefix.Concat(key);
        }
    }
}