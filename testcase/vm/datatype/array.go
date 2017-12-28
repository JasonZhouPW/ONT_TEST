package datatype

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestArray(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b6c766b00527ac4616c766b00c3c06c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Array},
		contract.ContractParameterType(contract.Integer),
		"TestArray",
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
	params := []interface{}{[]byte("Hello"), []byte("world")}
	if !testArray(ctx, code, params) {
		return false
	}
	params = []interface{}{[]byte("Hello"), []byte("world"), "123456", 8}
	if !testArray(ctx, code, params) {
		return false
	}
	return true
}

func testArray(ctx *testframework.TestFrameworkContext, code string, params []interface{}) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{params},
	)
	if err != nil {
		ctx.LogError("TestArray InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, len(params))
	if err != nil {
		ctx.LogError("TestArray test failed %s", err)
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
    public class Hello : SmartContract
    {
        public static int Main(params object[] args)
        {
            return args.Length;
        }
    }
}
*/
