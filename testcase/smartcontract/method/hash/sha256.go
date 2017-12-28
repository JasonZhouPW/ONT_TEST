package hash

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"crypto/sha256"
)

func TestSha256(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b6c766b00527ac4616c766b00c3a86c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Hash256),
		"TestSha256",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestSha256 DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestSha256 WaitForGenerateBlock error:%s", err)
		return false
	}
	input := []byte("Hello World")
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestSha256 InvokeSmartContract error:%s", err)
		return false
	}
	data := csha256(input)
	temp := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		temp[i] = data[i]
	}
	err = ctx.AssertToByteArray(res, temp)
	if err != nil {
		ctx.LogError("TestSha256 test failed %s", err)
		return false
	}
	return true
}

func csha256(input []byte) [32]byte{
	return sha256.Sum256(input)
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static byte[] Main(byte[] input)
    {
        return Sha256(input);
    }
}
*/
