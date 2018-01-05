package asset

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
        byte[] assetid = { 38, 85, 252, 250, 174, 21, 228, 237, 101, 153, 195, 139, 84, 65, 175, 229, 65, 28, 69, 80, 135, 133, 182, 14, 16, 30, 91, 102, 5, 144, 76, 129 };
        Asset ast = Blockchain.GetAsset(assetid);
        return ast.Issuer;
    }
}
*/

func TestGetIssuer(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b202655fcfaae15e4ed6599c38b5441afe5411c45508785b60e101e5b6605904c816c766b00527ac46c766b00c36168174e656f2e426c6f636b636861696e2e47657441737365746c766b51527ac46c766b51c36168134e656f2e41737365742e4765744973737565726c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetIssuer",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetIssuer DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetIssuer WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetIssuer InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetIssuer :%+v ", res)
	return true
}