package output

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static byte[] Main()
    {
        Block block = Blockchain.GetBlock(0);
        Transaction trans = block.GetTransaction(0);
        TransactionOutput[] outputs =  trans.GetOutputs();
        TransactionOutput output = outputs[0];
        return output.AssetId;
    }
}
*/

func TestGetOutputAssetId(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c300617c68184e656f2e426c6f636b2e4765745472616e73616374696f6e6c766b51527ac46c766b51c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b52527ac46c766b52c300c36c766b53527ac46c766b53c36168154e656f2e4f75747075742e476574417373657449646c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetOutputAssetId",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetOutputAssetId DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetOutputAssetId WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetOutputAssetId InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetOutputAssetId :%+v ", res)
	return true
}
