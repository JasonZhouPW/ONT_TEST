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
        return head.MerkleRoot;
    }
}
 */

func TestGetMerkleRoot(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b616168184e656f2e426c6f636b636861696e2e4765744865696768746c766b00527ac46c766b00c36168184e656f2e426c6f636b636861696e2e4765744865616465726c766b51527ac46c766b51c36168184e656f2e4865616465722e4765744d65726b6c65526f6f746c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetMerkleRoot",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetMerkleRoot DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetMerkleRoot WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetMerkleRoot InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestGetMerkleRoot :%+v ", res)
	return true
}
