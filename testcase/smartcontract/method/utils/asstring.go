package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestAsString(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6C766B00527AC46C766B00C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.String),
		"TestAsString",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestAsString DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestAsString WaitForGenerateBlock error:%s", err)
		return false
	}
	input := []byte("Hello World")
	if !testAsString(ctx,code, input){
		return false
	}
	input = []byte("")
	if !testAsString(ctx,code, input){
		return false
	}
	return true
}

func testAsString(ctx *testframework.TestFrameworkContext , code string, input []byte)bool{
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestAsString InvokeSmartContract error:%s", err)
		return false
	}

	err = ctx.AssertToString(res, string(input))
	if err != nil {
		ctx.LogError("TestAsString test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static string Main(byte[] input)
    {
        return input.AsString();
    }
}
*/