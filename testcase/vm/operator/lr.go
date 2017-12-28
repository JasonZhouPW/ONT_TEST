package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationLarger(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C3A0616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Boolean),
		"TestOperationLarger",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationLarger DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationLarger WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationLarger(ctx, code, 23, 2) {
		return false
	}

	if !testOperationLarger(ctx, code, -345, 34) {
		return false
	}

	if !testOperationLarger(ctx, code, -10, -234) {
		return false
	}

	if !testOperationLarger(ctx, code, 100, 100) {
		return false
	}

	return true
}

func testOperationLarger(ctx *testframework.TestFrameworkContext, code string, a, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationLarger InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, a > b)
	if err != nil {
		ctx.LogError("TestOperationLarger test %d > %d failed %s", a, b, err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static bool Main(int a, int b)
    {
        return a > b;
    }
}
*/
