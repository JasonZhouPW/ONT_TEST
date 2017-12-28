package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestRange(ctx *testframework.TestFrameworkContext) bool {
	code := "53C56B6C766B00527AC46C766B51527AC46C766B52527AC46C766B00C36C766B51C36C766B52C37F616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.ByteArray),
		"TestRange",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestRange DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestRange WaitForGenerateBlock error:%s", err)
		return false
	}

	input := []byte("Hello World!")
	if !testRange(ctx, code, input, 0, len(input)) {
		return false
	}

	if !testRange(ctx, code, input, 1, len(input)-2) {
		return false
	}

	if !testRange(ctx, code, input, 2, len(input) - 3) {
		return false
	}
	return true
}

func testRange(ctx *testframework.TestFrameworkContext, code string, b []byte, start, count int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{b, start, count},
	)
	if err != nil {
		ctx.LogError("TestRange InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToByteArray(res, b[start:start+count])
	if err != nil {
		ctx.LogError("TestRange test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static byte[] Main(byte[] arg, int start, int count)
    {
        return arg.Range(start, count);
    }
}
*/
