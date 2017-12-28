package cond_loop

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestSwitch(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B00C36C766B51527AC46C766B51C351907C907C9E63080051616C756600616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestSwitch",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestSwitch DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestSwitch WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testSwitch(ctx, code, 23) {
		return false
	}

	if !testSwitch(ctx, code, 1) {
		return false
	}

	if !testSwitch(ctx, code, 0) {
		return false
	}

	return true
}

func testSwitch(ctx *testframework.TestFrameworkContext, code string, a int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("TestSwitch InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, tswitch(a))
	if err != nil {
		ctx.LogError("TestSwitch test switch %d failed %s", a, err)
		return false
	}
	return true
}

func tswitch(a int) int {
	switch a {
	case 1:
		return 1
	default:
		return 0
	}
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static int Main(int a)
    {
        switch(a)
        {
            case 1:
                return 1;
            default:
                return 0;
        }
    }
}
*/