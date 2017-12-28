package block

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"

)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static Transaction Main()
    {
        Block block = Blockchain.GetBlock(0);
        Transaction trans = block.GetTransaction(0);
        return trans;
    }
}
 */

func TestGetBlockTransaction(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c300617c68184e656f2e426c6f636b2e4765745472616e73616374696f6e6c766b51527ac46c766b51c36c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.InteropInterface),
		"TestGetBlockTransaction",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransaction DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBlockTransaction WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransaction InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestGetBlockTransaction :%+v ", res)
	return true
}
