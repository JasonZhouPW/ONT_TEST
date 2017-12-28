package cond_loop

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestFor(ctx *testframework.TestFrameworkContext) bool {
	code := "53C56B6C766B00527AC4006C766B51527AC4006C766B52527AC46223006C766B51C36C766B52C3936C766B51527AC46C766B52C351936C766B52527AC46C766B52C36C766B00C39F63D5FF6C766B51C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestFor",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestFor DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestFor WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testFor(ctx, code, 23) {
		return false
	}

	if !testFor(ctx, code, -23) {
		return false
	}

	if !testFor(ctx, code, 0) {
		return false
	}

	return true
}

func testFor(ctx *testframework.TestFrameworkContext, code string, a int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("TestFor InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, forloop(a))
	if err != nil {
		ctx.LogError("TestFor test for %d failed %s", a, err)
		return false
	}
	return true
}

func forloop(a int) int {
	b := 0
	for i := 0; i < a; i++{
		b = b+i
	}
	return b
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static int Main(int a)
    {
        int b = 0;
        for(int i = 0;i < a;i++)
        {
            b = b+i;
        }
        return b;
    }
}
*/