package transaction

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"encoding/json"
)

func  TestGetInputs(ctx *testframework.TestFrameworkContext)bool{
	txHash, err := getTransferTransaction(ctx, ctx.OntClient.Account1, ctx.OntClient.Account2)
	if err != nil {
		ctx.LogError("initTransaction error:%s", err)
		return false
	}

	code := "59c56b6c766b00527ac4616c766b00c361681d4e656f2e426c6f636b636861696e2e4765745472616e73616374696f6e6c766b51527ac46c766b51c36168194e656f2e5472616e73616374696f6e2e476574496e707574736c766b52527ac46c766b52c3c0c56c766b53527ac4006c766b54527ac4627f00616c766b52c36c766b54c3c36c766b55527ac452c56c766b56527ac46c766b56c3006c766b55c36168114e656f2e496e7075742e47657448617368c46c766b56c3516c766b55c36168124e656f2e496e7075742e476574496e646578c46c766b53c36c766b54c36c766b56c3c4616c766b54c351936c766b54527ac46c766b54c36c766b52c3c09f6c766b57527ac46c766b57c3636cff6c766b53c36c766b58527ac46203006c766b58c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Array),
		"TestGetInputs",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetInputs DeploySmartContract error:%s", err)
		return false
	}

	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetInputs WaitForGenerateBlock error:%s", err)
		return false
	}

	tx, err := ctx.Ont.GetTransaction(txHash)
	d, _ := json.Marshal(tx.UTXOInputs)
	ctx.LogInfo("TestGetInputs Inputs:%s", d)

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{txHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetInputs InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogInfo("TestGetInputs res:%v", res)
	return true
}
/*

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class A : SmartContract
{
    public static object[] Main(byte[] txHash)
    {
        Transaction tx = Blockchain.GetTransaction(txHash);
        TransactionInput[] inputs = tx.GetInputs();
        object[] ret = new object[inputs.Length];
        for (int i=0; i < inputs.Length; i++)
        {
            TransactionInput input = inputs[i];
            object[] item = new object[2];
            item[0] = input.PrevHash;
            item[1] = input.PrevIndex;
            ret[i] = item;
        }
        return ret;
    }
}
*/