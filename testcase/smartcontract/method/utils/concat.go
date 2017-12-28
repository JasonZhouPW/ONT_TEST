package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestConcat(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C37E616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestConcat",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestConcat DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestConcat WaitForGenerateBlock error:%s", err)
		return false
	}
	input1 := "Hello"
	input2 := "World"
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input1, input2},
	)
	if err != nil {
		ctx.LogError("TestConcat InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToByteArray(res, []byte(string(input1)+string(input2)))
	if err != nil {
		ctx.LogError("TestConcat test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static byte[] Main(byte[] arg1, byte[] arg2)
    {
        return arg1.Concat(arg2);
    }
}
*/
