package account

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
    public static object Main()
    {
        byte[] programHash = { 222, 22, 168, 155, 127, 237, 137, 116, 234, 99, 88, 103, 178, 63, 254, 214, 234, 83, 239, 81 };
        Account account = Blockchain.GetAccount(programHash);
	byte[] assetid = { 38, 85, 252, 250, 174, 21, 228, 237, 101, 153, 195, 139, 84, 65, 175, 229, 65, 28, 69, 80, 135, 133, 182, 14, 16, 30, 91, 102, 5, 144, 76, 129 };
        long balance = account.GetBalance(assetid);
        return account;
    }
}
 */

func TestGetBalance(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b14de16a89b7fed8974ea635867b23ffed6ea53ef516c766b00527ac46c766b00c36168194e656f2e426c6f636b636861696e2e4765744163636f756e746c766b51527a202655fcfaae15e4ed6599c38b5441afe5411c45508785b60e101e5b6605904c816c766b52527ac46c766b51c36c766b52c3617c68164e656f2e4163636f756e742e47657442616c616e63656c766b53527ac46c766b51c36c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestGetBalance",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBalance DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBalance WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetBalance InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetBalance :%+v ", res)
	return true
}