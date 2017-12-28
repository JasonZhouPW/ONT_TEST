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
    public static int Main()
    {
        Block block = Blockchain.GetBlock(0);
        return block.GetTransactionCount();
    }
}
 */

func TestGetBlockTransctionCount(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c361681d4e656f2e426c6f636b2e4765745472616e73616374696f6e436f756e746c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestGetBlockTransctionCount",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransctionCount DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBlockTransctionCount WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetBlockTransctionCount InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestGetBlockTransctionCount :%+v ", res)
	return true
}

