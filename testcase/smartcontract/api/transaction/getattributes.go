package transaction

import (
	"encoding/json"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestGetAttributes(ctx *testframework.TestFrameworkContext) bool {
	txHash, err := getTransferTransaction(ctx, ctx.OntClient.Account1, ctx.OntClient.Account2)
	if err != nil {
		ctx.LogError("initTransaction error:%s", err)
		return false
	}

	code := "59c56b6c766b00527ac4616c766b00c361681d4e656f2e426c6f636b636861696e2e4765745472616e73616374696f6e6c766b51527ac46c766b51c361681d4e656f2e5472616e73616374696f6e2e476574417474726962757465736c766b52527ac46c766b52c3c0c56c766b53527ac4006c766b54527ac4628700616c766b52c36c766b54c3c36c766b55527ac452c56c766b56527ac46c766b56c3006c766b55c36168164e656f2e4174747269627574652e4765745573616765c46c766b56c3516c766b55c36168154e656f2e4174747269627574652e47657444617461c46c766b53c36c766b54c36c766b56c3c4616c766b54c351936c766b54527ac46c766b54c36c766b52c3c09f6c766b57527ac46c766b57c36364ff6c766b53c36c766b58527ac46203006c766b58c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Array),
		"TestGetAttributes",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetAttributes DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetAttributes WaitForGenerateBlock error:%s", err)
		return false
	}

	tx, err := ctx.Ont.GetTransaction(txHash)
	if err != nil {
		ctx.LogError("TestGetAttributes GetTransaction error:%s", err)
		return false
	}
	txAttrs := tx.Attributes
	d, _ := json.Marshal(txAttrs)
	ctx.LogInfo("TestGetAttributes Attributes:%s", d)

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{txHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetAttributes InvokeSmartContract error:%s", err)
		return false
	}

	ret, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestGetAttributes asset res []interface error")
		return false
	}

	for i, item := range ret {
		attr, ok := item.([]interface{})
		if !ok {
			ctx.LogError("TestGetAttributes asset item []interface error")
			return false
		}

		txAttr := txAttrs[i]
		err := ctx.AssertToInt(attr[0], int(txAttr.Usage))
		if err != nil {
			ctx.LogError("TestGetAttributes Usage AssertToInt error:%s", err)
			return false
		}

		err = ctx.AssertToByteArray(attr[1], txAttr.Data)
		if err != nil {
			ctx.LogError("TestGetAttributes Data AssertToByteArray error:%s", err)
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
        TransactionAttribute[] attrs = tx.GetAttributes();
        object[] ret = new object[attrs.Length];
        for (int i = 0; i < attrs.Length; i++)
        {
            TransactionAttribute attr = attrs[i];
            object[] item = new object[2];
            item[0] = attr.Usage;
            item[1] = attr.Data;
            ret[i] = item;
        }
        return ret;
    }
}
*/
