package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationAnd(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C3640C006C766B51C3616C756600616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Boolean, contract.Boolean},
		contract.ContractParameterType(contract.Boolean),
		"TestOperationAnd",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationAnd DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationAnd WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationAnd(ctx, code, true,true) {
		return false
	}

	if !testOperationAnd(ctx, code, true, false){
		return false
	}

	if !testOperationAnd(ctx, code, false, true){
		return false
	}

	if !testOperationAnd(ctx, code, false, false){
		return false
	}

	return true
}

func testOperationAnd(ctx *testframework.TestFrameworkContext, code string, a, b bool) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationAnd InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, a&&b)
	if err != nil {
		ctx.LogError("TestOperationAnd test failed %s", err)
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
        return a && b;
    }
}
*/
