package datatype

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestByteArray(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c39c6c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.ByteArray},
		contract.ContractParameterType(contract.Boolean),
		"TestByteArray",
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
		ctx.LogError("TestArray WaitForGenerateBlock error:%s", err)
		return false
	}

	arg1 := []byte("Hello")
	arg2 := []byte("World")

	if !testByteArray(ctx, code, arg1, arg1, true) {
		return false
	}
	if !testByteArray(ctx, code, arg1, arg2, false) {
		return false
	}
	return true
}

func testByteArray(ctx *testframework.TestFrameworkContext, code string, arg1, arg2 []byte, expect bool) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{arg1, arg2},
	)
	if err != nil {
		ctx.LogError("testByteArray InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, expect)
	if err != nil {
		ctx.LogError("testByteArray test failed %s", err)
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

namespace Hello
{
    public class A : SmartContract
    {
        public static bool Main(byte[] arg1, byte[] arg2)
        {
            return arg1 == arg2;
        }
    }
}
*/
