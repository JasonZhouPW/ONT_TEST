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
    [DisplayName("transfer")]
    public static event Action<byte[], byte[], BigInteger> Transferred;
    public static void Main(byte[] from, byte[] to, BigInteger amount)
    {
        Transferred(from, to, amount);
    }
}
 */

func TestNotify(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6c766b00527ac46c766b51527ac46c766b52527ac4616c766b00c36c766b51c36c766b52c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.ByteArray,contract.Integer},
		contract.ContractParameterType(contract.Void),
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
		[]interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), ctx.OntClient.Account2.ProgramHash.ToArray(), 10},
	)
	if err != nil {
		ctx.LogError("TestNotify InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestNotify :%+v ", res)
	return true
}
