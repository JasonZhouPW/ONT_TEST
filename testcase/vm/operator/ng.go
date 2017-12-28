package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationNegative(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6C766B00527AC46C766B00C3009C616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Boolean},
		contract.ContractParameterType(contract.Boolean),
		"TestOperationNegative",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationNegative DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationNegative WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationNegative(ctx, code, true) {
		return false
	}

	if !testOperationNegative(ctx, code, false) {
		return false
	}

	return true
}

func testOperationNegative(ctx *testframework.TestFrameworkContext, code string, a bool) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("TestOperationNegative InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, !a)
	if err != nil {
		ctx.LogError("TestOperationNegative test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static bool Main(bool a)
    {
        return !a;
    }
}
*/
