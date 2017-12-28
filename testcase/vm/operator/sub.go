package operator

import (
"github.com/ONT_TEST/testframework"
"github.com/Ontology/core/contract"
"github.com/Ontology/smartcontract/types"
"time"
)

func TestOperationSub(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C394616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestOperationSub",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationSub DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationSub WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationSub(ctx, code, 1,2) {
		return false
	}

	if !testOperationSub(ctx, code, -1, 1){
		return false
	}

	if !testOperationSub(ctx, code, -1, -2){
		return false
	}

	if !testOperationSub(ctx, code, 0, 0){
		return false
	}

	return true
}

func testOperationSub(ctx *testframework.TestFrameworkContext, code string, a, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationSub InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, a-b)
	if err != nil {
		ctx.LogError("TestOperationSub test failed %s", err)
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
        return a - b;
    }
}
*/
