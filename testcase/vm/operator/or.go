package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationOr(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C3630C006C766B51C3616C756651616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Boolean, contract.Boolean},
		contract.ContractParameterType(contract.Boolean),
		"TestOperationOr",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationOr DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationOr WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationOr(ctx, code, true,true) {
		return false
	}

	if !testOperationOr(ctx, code, true, false){
		return false
	}

	if !testOperationOr(ctx, code, false, true){
		return false
	}

	if !testOperationOr(ctx, code, false, false){
		return false
	}

	return true
}

func testOperationOr(ctx *testframework.TestFrameworkContext, code string, a, b bool) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationOr InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, a||b)
	if err != nil {
		ctx.LogError("TestOperationOr test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static bool Main(bool a, bool b)
    {
        return a || b;
    }
}
*/
