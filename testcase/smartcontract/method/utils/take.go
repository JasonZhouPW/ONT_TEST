package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestTake(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c3806c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.Integer},
		contract.ContractParameterType(contract.ByteArray),
		"TestTake",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestTake DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestTake WaitForGenerateBlock error:%s", err)
		return false
	}

	input := []byte("Hello World!")
	if !testTake(ctx, code, input, 0) {
		return false
	}

	if !testTake(ctx, code, input, len(input)-1) {
		return false
	}

	if !testTake(ctx, code, input, len(input)) {
		return false
	}
	return true
}

func testTake(ctx *testframework.TestFrameworkContext, code string, b []byte, count int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{b, count},
	)
	if err != nil {
		ctx.LogError("TestTake InvokeSmartContract error:%s", err)
		return false
	}
	r := count
	if count > len(b){
		r = len(b)
	}
	err = ctx.AssertToByteArray(res, b[0:r])
	if err != nil {
		ctx.LogError("TestTake test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static byte[] Main(byte[] arg, int count)
    {
        return arg.Take(count);
    }
}
*/
