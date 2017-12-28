package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationSelfAdd(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6C766B00527AC46C766B00C35193766A00527AC4616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestOperationSelfAdd",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationSelfAdd DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationSelfAdd WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationSelfAdd(ctx, code, 1) {
		return false
	}

	if !testOperationSelfAdd(ctx, code, -1){
		return false
	}

	if !testOperationSelfAdd(ctx, code, 0){
		return false
	}

	return true
}

func testOperationSelfAdd(ctx *testframework.TestFrameworkContext, code string, a int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("TestOperationSelfAdd InvokeSmartContract error:%s", err)
		return false
	}
	a ++
	err = ctx.AssertToInt(res, a)
	if err != nil {
		ctx.LogError("TestOperationSelfAdd test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static int Main(int a)
    {
        return ++a;
    }
}
*/
