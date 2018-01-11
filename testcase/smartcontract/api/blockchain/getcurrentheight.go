package blockchain

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
        return (int)Blockchain.GetHeight();
    }
}
*/

func TestGetCurrentHeight(ctx *testframework.TestFrameworkContext) bool {
	code := "51c56b616168184e656f2e426c6f636b636861696e2e4765744865696768746c766b00527ac46203006c766b00c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestGetCurrentHeight",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight WaitForGenerateBlock error:%s", err)
		return false
	}

	height, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("TestGetCurrentHeight GetBlockCount error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight InvokeSmartContract error:%s", err)
		return false
	}

	err = ctx.AssertToInt(res, int(height)-1)
	if err != nil{
		ctx.LogError("TestGetCurrentHeight res AssertToInt error:%s", err)
		return false
	}
	return true
}
