package event

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

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System;
using System.ComponentModel;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static void Main()
    {
        Runtime.Log("Hello World!");

    }
}
 */

func TestLog(ctx *testframework.TestFrameworkContext) bool {
	code := "00c56b610c48656c6c6f20576f726c642161680f4e656f2e52756e74696d652e4c6f6761616c7566"
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
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestLog InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestLog :%+v ", res)
	return true
}
