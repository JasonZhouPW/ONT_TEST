package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	// "math/big"
)

func TestOperationDivide(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C396616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestOperationDivide",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationDivide DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationDivide WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationDivideFail(ctx, code, 10, 0) {
		return false
	}

	if !testOperationDivide(ctx, code, 23, 2) {
		return false
	}

	if !testOperationDivide(ctx, code, 544, 345) {
		return false
	}

	if !testOperationDivide(ctx, code, 3456345, 3545) {
		return false
	}

	if !testOperationDivide(ctx, code, -10, -234) {
		return false
	}

	if !testOperationDivide(ctx, code, -345, 34) {
		return false
	}

	if !testOperationDivide(ctx, code, 0, 100) {
		return false
	}

	return true
}

func testOperationDivide(ctx *testframework.TestFrameworkContext, code string, a, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationDivide InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, a/b)
	if err != nil {
		ctx.LogError("TestOperationDivide test %d / %d failed %s", a, b, err)
		return false
	}
	return true
}

func testOperationDivideFail(ctx *testframework.TestFrameworkContext, code string, a, b int) bool {
	_, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err == nil {
		ctx.LogError("testOperationDivideFail %v / %v should failed", a, b)
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
        return a / b;
    }
}
*/
