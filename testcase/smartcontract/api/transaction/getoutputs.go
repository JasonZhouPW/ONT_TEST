package transaction

import (
	"encoding/json"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestGetOutputs(ctx *testframework.TestFrameworkContext) bool {
	txHash, err := getTransferTransaction(ctx, ctx.OntClient.Account1, ctx.OntClient.Account2)
	if err != nil {
		ctx.LogError("initTransaction error:%s", err)
		return false
	}

	code := "59c56b6c766b00527ac4616c766b00c361681d4e656f2e426c6f636b636861696e2e4765745472616e73616374696f6e6c766b51527ac46c766b51c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b52527ac46c766b52c3c0c56c766b53527ac4006c766b54527ac462ab00616c766b52c36c766b54c3c36c766b55527ac453c56c766b56527ac46c766b56c3006c766b55c36168154e656f2e4f75747075742e47657441737365744964c46c766b56c3516c766b55c36168184e656f2e4f75747075742e47657453637269707448617368c46c766b56c3526c766b55c36168134e656f2e4f75747075742e47657456616c7565c46c766b53c36c766b54c36c766b56c3c4616c766b54c351936c766b54527ac46c766b54c36c766b52c3c09f6c766b57527ac46c766b57c36340ff6c766b53c36c766b58527ac46203006c766b58c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Array),
		"TestGetOutputs",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetOutputs DeploySmartContract error:%s", err)
		return false
	}

	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetOutputs WaitForGenerateBlock error:%s", err)
		return false
	}

	tx, err := ctx.Ont.GetTransaction(txHash)
	txOutputs := tx.Outputs
	d, _ := json.Marshal(txOutputs)
	ctx.LogInfo("TestGetOutputs Outputs:%s", d)

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{txHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetOutputs InvokeSmartContract error:%s", err)
		return false
	}

	ret, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestGetOutputs asset res to []interface error")
		return false
	}

	for i, item := range ret {
		output, ok := item.([]interface{})
		if !ok {
			ctx.LogError("TestGetOutputs asset item to []interface error")
			return false
		}
		txOutput := txOutputs[i]
		err := ctx.AssertToByteArray(output[0], txOutput.AssetID.ToArray())
		if err != nil {
			ctx.LogError("TestGetOutputs AssetID AssertToByteArray error:%s", err)
			return false
		}
		err = ctx.AssertToByteArray(output[1], txOutput.ProgramHash.ToArray())
		if err != nil {
			ctx.LogError("TestGetOutputs ProgramHash AssertToByteArray error:%s", err)
			return false
		}
		err = ctx.AssertToInt(output[2], int(txOutput.Value))
		if err != nil {
			ctx.LogError("TestGetOutputs Value AssertToInt error:%s", err)
			return false
		}
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
    public static object[] Main(byte[] txHash)
    {
        Transaction tx = Blockchain.GetTransaction(txHash);
        TransactionOutput[] outputs = tx.GetOutputs();
        object[] ret = new object[outputs.Length];
        for (int i=0; i < outputs.Length; i++)
        {
            TransactionOutput output = outputs[i];
            object[] item = new object[3];
            item[0] = output.AssetId;
            item[1] = output.ScriptHash;
            item[2] = output.Value;
            ret[i] = item;
        }
        return ret;
    }
}
*/
