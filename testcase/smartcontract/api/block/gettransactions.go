package block

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"

	"github.com/ONT_TEST/testcase/smartcontract/api/helper"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object[] Main(int height)
    {
        Block block = Blockchain.GetBlock((uint)height);
        Transaction[] txs = block.GetTransactions();
        int count = block.GetTransactionCount();
        object[] ret = new object[count];
        for (int i = 0; i < count; i++)
        {
            ret[i] = txs[i].Hash;
        }
        return ret;
    }
}
*/

func TestGetBlockTransactions(ctx *testframework.TestFrameworkContext) bool {
	err := helper.InitAsset(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions InitAsset error:%s", err)
		return false
	}
	code := "58c56b6c766b00527ac4616c766b00c36168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b51527ac46c766b51c36168194e656f2e426c6f636b2e4765745472616e73616374696f6e736c766b52527ac46c766b51c361681d4e656f2e426c6f636b2e4765745472616e73616374696f6e436f756e746c766b53527ac46c766b53c3c56c766b54527ac4006c766b55527ac4624300616c766b54c36c766b55c36c766b52c36c766b55c3c36168174e656f2e5472616e73616374696f6e2e47657448617368c4616c766b55c351936c766b55527ac46c766b55c36c766b53c39f6c766b56527ac46c766b56c363a9ff6c766b54c36c766b57527ac46203006c766b57c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Array),
		"TestGetBlockTransactions",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions DeploySmartContract error:%s", err)
		return false
	}

	tx, err := helper.Transfer(ctx, helper.AssetId, ctx.OntClient.Account1, ctx.OntClient.Account2, 1.0)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions Transfer error:%s", err)
		return false
	}

	ctx.LogInfo("TestGetBlockTransactions Transfer TxHash:%x", tx.ToArray())

	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions WaitForGenerateBlock error:%s", err)
		return false
	}

	height, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("TestGetBlockTransactions GetBlockCount error:%s", err)
		return false
	}

	height -= 1
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{height},
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions InvokeSmartContract error:%s", err)
		return false
	}

	block, err := ctx.Ont.GetBlockByHeight(height)
	if err != nil {
		ctx.LogError("TestGetBlockTransactions GetBlockByHeight error:%s", err)
		return false
	}

	txList := block.Transactions

	ret, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestGetBlockTransactions res asset to []interface{} failed")
		return false
	}

	if len(txList) != len(ret) {
		ctx.LogError("TestGetBlockTransactions TxCount %v != %v", len(txList), len(ret))
		return false
	}

	for i, txHash := range ret {
		txh := txList[i].Hash()
		err := ctx.AssertToByteArray(txHash, txh.ToArray())
		if err != nil {
			ctx.LogError("TestGetBlockTransactions TxHash AssertToByteArray error:%s", err)
			return false
		}
	}

	return true
}
