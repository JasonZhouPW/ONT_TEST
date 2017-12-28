package operator

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestOperationRightShift(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C3011F8499616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestOperationRightShift",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestOperationRightShift DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 2)
	if err != nil {
		ctx.LogError("TestOperationRightShift WaitForGenerateBlock error:%s", err)
		return false
	}

	 if !testOperationRightShift(ctx, code, 1, 2) {
	 	return false
	 }

	 if !testOperationRightShift(ctx, code, 34252452, 3) {
	 	return false
	 }

	 if !testOperationRightShift(ctx, code, -1, 2) {
	 	return false
	 }

	if !testOperationRightShitFail(ctx, code, 1, -1) {
		return false
	}

	return true
}

func testOperationRightShift(ctx *testframework.TestFrameworkContext, code string, a int, b uint) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationRightShift InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, a>>b)
	if err != nil {
		ctx.LogError("TestOperationRightShift test %d >> %d failed %s", a, b, err)
		return false
	}
	return true
}

func testOperationRightShitFail(ctx *testframework.TestFrameworkContext, code string, a int, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err == nil {
		ctx.LogError("TestOperationRightShift %v >> %v should failed, but get %v", a, b, res)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static BigInteger Main(BigInteger a, int b)
    {
        return a >> b;
    }
}
*/
