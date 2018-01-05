package input

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
    public static int Main()
    {
        Block block = Blockchain.GetBlock(0);
        Transaction trans = block.GetTransaction(0);
        TransactionInput[] inputs =  trans.GetInputs();
        TransactionInput input = inputs[0];
        return input.PrevIndex;
    }
}
*/

func TestGetInputPrevIndex(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c300617c68184e656f2e426c6f636b2e4765745472616e73616374696f6e6c766b51527ac46c766b51c36168194e656f2e5472616e73616374696f6e2e476574496e707574736c766b52527ac46c766b52c300c36c766b53527ac46c766b53c36168124e656f2e496e7075742e476574496e6465786c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetInputPrevIndex",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetInputPrevIndex DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetInputPrevIndex WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetInputPrevIndex InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetInputPrevIndex :%+v ", res)
	return true
}
