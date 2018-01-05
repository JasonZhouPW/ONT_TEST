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

public class HelloWorld : SmartContract
{
    [DisplayName("transfer")]
    public static event Action<byte[], byte[], BigInteger> Transferred;
    public static void Main(byte[] pk)
    {
        byte[] p1 = { 1 };
        byte[] p2 = { 2 };
        Transferred(p1, p2, (BigInteger)12);
    }
}
 */

func TestNotify(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6101016c766b00527ac401026c766b51527ac453c576006c766b00c3c476516c766b51c3c476525cc46c766b52527ac46c766b52c36168124e656f2e52756e74696d652e4e6f7469667961616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestNotify",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestNotify DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestNotify WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestNotify InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestNotify :%+v ", res)
	return true
}
