package header

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
    public static byte[] Main()
    {
        uint height = Blockchain.GetHeight();
        Header head = Blockchain.GetHeader(height);
        return head.PrevHash;
    }
}
 */

func TestGetPrevHash(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b616168184e656f2e426c6f636b636861696e2e4765744865696768746c766b00527ac46c766b00c36168184e656f2e426c6f636b636861696e2e4765744865616465726c766b51527ac46c766b51c36168164e656f2e4865616465722e47657450726576486173686c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetPrevHash",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetPrevHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetPrevHash WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetPrevHash InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestGetPrevHash :%+v ", res)
	return true
}
