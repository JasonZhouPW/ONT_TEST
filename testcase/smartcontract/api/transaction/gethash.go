package transaction

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestGetTxHash(ctx *testframework.TestFrameworkContext) bool {
	txHash, err := getTransferTransaction(ctx, ctx.OntClient.Account1, ctx.OntClient.Account2)
	if err != nil {
		ctx.LogError("initTransaction error:%s", err)
		return false
	}

	code := "53c56b6c766b00527ac4616c766b00c361681d4e656f2e426c6f636b636861696e2e4765745472616e73616374696f6e6c766b51527ac46c766b51c36168174e656f2e5472616e73616374696f6e2e476574486173686c766b52527ac46203006c766b52c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetTxHash",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetTxHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetTxHash WaitForGenerateBlock error:%s", err)
		return false
	}

	//tx, err := ctx.Ont.GetTransaction(txHash)
	ctx.LogInfo("TestGetTxHash TxHash:%x", txHash.ToArray())

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{txHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetTxHash InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogInfo("TestGetTxHash res:%v", res)


	err = ctx.AssertToByteArray(res, txHash.ToArray())
	if err != nil {
		ctx.LogError("TestGetTxHash test failed %s", err)
		return false
	}
	return true
}

/*

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class A : SmartContract
{
    public static byte[] Main(byte[] txHash)
    {
        Transaction tx = Blockchain.GetTransaction(txHash);
        return tx.Hash;
    }
}
*/
