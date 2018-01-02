package datatype

import (
	"encoding/hex"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	vmtype "github.com/Ontology/vm/neovm/types"
	"time"
)

func TestReturnType(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b6c766b00527ac46c766b51527ac46c766b52527ac46153c56c766b53527ac46c766b53c3006c766b00c3c46c766b53c3516c766b51c3c46c766b53c3526c766b52c3c46c766b53c36c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Array),
		"TestReturnType",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestArray DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestReturnType WaitForGenerateBlock error:%s", err)
		return false
	}
	if !testReturnType(ctx, code, []int{100343, 2433554}, []byte("Hello world")) {
		return false
	}
	return true
}

func testReturnType(ctx *testframework.TestFrameworkContext, code string, args []int, arg3 []byte) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{args[0], args[1], arg3},
	)
	if err != nil {
		ctx.LogError("TestReturnType InvokeSmartContract error:%s", err)
		return false
	}

	rt, ok := res.([]interface{})
	if !ok {
		ctx.LogError("%s assert to array failed.", res)
		return false
	}

	vs, ok := rt[0].(string)
	if !ok {
		ctx.LogError("%s assert ")
	}
	v, err := hex.DecodeString(vs)
	if err != nil {
		ctx.LogError("hex.DecodeString:%s error:%s", err)
		return false
	}

	b := vmtype.ConvertBytesToBigInteger(v)
	if int(b.Int64()) != args[0]{
		ctx.LogError("%d != %d ", b.Int64(), args[0])
		return false
	}

	err = ctx.AssertToByteArray(rt[2], arg3)
	if err != nil {
		ctx.LogError("AssertToByteArray error:%s", err)
		return false
	}

	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System;
using System.Numerics;
namespace ONT_DEx
{
    public class ONT_P2P : SmartContract
    {
        public static object[] Main(int arg1, int arg2, byte[] arg3)
        {
            object[] ret = new object[3];
            ret[0] = arg1;
            ret[1] = arg2;
            ret[2] = arg3;
            return ret;
        }
    }
}
*/
