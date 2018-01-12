package runtime

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
)


/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System;
using System.ComponentModel;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static void Main(string msg)
    {
        Runtime.Log(msg);
    }
}
 */

func TestLog(ctx *testframework.TestFrameworkContext) bool {
	code := "51c56b6c766b00527ac4616c766b00c361680f4e656f2e52756e74696d652e4c6f6761616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestLog",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestLog DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestLog WaitForGenerateBlock error:%s", err)
		return false
	}

	input := "Hello World!"
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestLog InvokeSmartContract error:%s", err)
		return false
	}

	notify , ok := res.(map[string]interface{})
	if !ok {
		ctx.LogError("TestLog res asset to map[string]interface{} error:%s", err)
		return false
	}

	log := notify["Message"]
	if input !=log {
		ctx.LogError("TestLog log error %s != %s", input, log)
		return false
	}
	return true
}
