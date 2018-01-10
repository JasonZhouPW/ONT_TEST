package account

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testcase/smartcontract/api/helper"
	"time"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class A : SmartContract
{
    public static byte[] Main(byte[] programHash)
    {
        Account account = Blockchain.GetAccount(programHash);
        return account.ScriptHash;
    }
}
*/

func TestGetScriptHash(ctx *testframework.TestFrameworkContext) bool {
	account := ctx.OntClient.Account1
	err := helper.InitAsset(ctx,account )
	if err != nil {
		ctx.LogError("TestGetScriptHash InitAsset error:%s", err)
		return false
	}
	code := "53c56b6c766b00527ac4616c766b00c36168194e656f2e426c6f636b636861696e2e4765744163636f756e746c766b51527ac46c766b51c36168194e656f2e4163636f756e742e476574536372697074486173686c766b52527ac46203006c766b52c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetScriptHash",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{account.ProgramHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetScriptHash InvokeSmartContract error:%s", err)
		return false
	}

	err = ctx.AssertToByteArray(res, account.ProgramHash.ToArray())
	if err != nil {
		ctx.LogError("TestGetScriptHash AssertToByteArray error:%s", err)
		return false
	}
	return true
}
