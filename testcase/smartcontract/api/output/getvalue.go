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
    public static long Main()
    {
        Block block = Blockchain.GetBlock(0);
        Transaction trans = block.GetTransaction(0);
        TransactionOutput[] outputs =  trans.GetOutputs();
        TransactionOutput output = outputs[0];
        return output.Value;
    }
}
*/

func TestGetOutputValue(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c300617c68184e656f2e426c6f636b2e4765745472616e73616374696f6e6c766b51527ac46c766b51c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b52527ac46c766b52c300c36c766b53527ac46c766b53c36168134e656f2e4f75747075742e47657456616c75656c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetOutputValue",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetOutputValue DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetOutputValue WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetOutputValue InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetOutputValue :%+v ", res)
	return true
}
