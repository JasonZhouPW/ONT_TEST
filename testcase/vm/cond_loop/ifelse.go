package cond_loop

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestIfElse(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C3A163080051616C75666C766B00C36C766B51C3A26308004F616C756600616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestIfElse",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestIfElse DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestIfElse WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testIfElse(ctx, code, 23, 2) {
		return false
	}

	if !testIfElse(ctx, code, 2, 23) {
		return false
	}

	if !testIfElse(ctx, code, 0, 0) {
		return false
	}

	return true
}

func testIfElse(ctx *testframework.TestFrameworkContext, code string, a, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestIfElse InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, condIfElse(a,b))
	if err != nil {
		ctx.LogError("TestIfElse test %d ifelse %d failed %s", a, b, err)
		return false
	}
	return true
}

func condIfElse(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static int Main(int a, int b)
    {
        if(a > b)
        {
            return 1;
        }
        else if(a < b)
        {
            return -1;
        }
        else{
            return 0;
        }
    }
}
*/
