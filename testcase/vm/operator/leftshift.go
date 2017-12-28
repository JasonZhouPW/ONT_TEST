package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationLeftShift(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C3011F8498616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestOperationLeftShift",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationLeftShift DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 2)
	if err != nil {
		ctx.LogError("TestOperationLeftShift WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationLeftShift(ctx, code, 1, 2) {
		return false
	}

	if !testOperationLeftShift(ctx, code, 1326567565434, 2) {
		return false
	}

	if !testOperationLeftShift(ctx, code, 2, 3) {
		return false
	}

	 if !testOperationLeftShift(ctx, code, -1, 2) {
	 	return false
	 }

	 if !testOperationLeftShitFail(ctx, code, 1, -1) {
	 	return false
	 }

	return true
}

func testOperationLeftShift(ctx *testframework.TestFrameworkContext, code string, a int, b uint) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationLeftShift InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, a<<b)
	if err != nil {
		ctx.LogError("TestOperationLeftShift test %d << %d failed %s", a, b, err)
		return false
	}
	return true
}

func testOperationLeftShitFail(ctx *testframework.TestFrameworkContext, code string, a int, b int) bool {
	_, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err == nil {
		ctx.LogError("testOperationLeftShitFail %v << %v should failed", a, b)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static int Main(int a, int b)
    {
        return a << b;
    }
}
*/
