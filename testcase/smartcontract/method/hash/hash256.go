package hash

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"github.com/Ontology/vm/neovm"
)

func TestHash256(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6C766B00527AC46C766B00C3AA616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Hash256),
		"TestHash256",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestHash256 DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestHash256 WaitForGenerateBlock error:%s", err)
		return false
	}
	input := []byte("Hello World")
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestHash256 InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToByteArray(res, hash256(input))
	if err != nil {
		ctx.LogError("TestHash256 test failed %s", err)
		return false
	}
	return true
}

func hash256(input []byte) []byte{
	return new(neovm.ECDsaCrypto).Hash256(input)
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static byte[] Main(byte[] input)
    {
        return Hash256(input);
    }
}
*/
