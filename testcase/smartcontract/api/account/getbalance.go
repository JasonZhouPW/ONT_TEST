package account

import (
	"github.com/ONT_TEST/testcase/smartcontract/api/helper"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System.Numerics;
class A : SmartContract
{
    public static BigInteger Main(byte[] programHash, byte[]assetId)
    {
        Account account = Blockchain.GetAccount(programHash);
        return account.GetBalance(assetId);
    }
}
*/

func TestGetBalance(ctx *testframework.TestFrameworkContext) bool {
	err := helper.InitAsset(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestGetBalance InitAsset error:%s", err)
		return false
	}
	code := "54c56b6c766b00527ac46c766b51527ac4616c766b00c36168194e656f2e426c6f636b636861696e2e4765744163636f756e746c766b52527ac46c766b52c36c766b51c3617c68164e656f2e4163636f756e742e47657442616c616e63656c766b53527ac46203006c766b53c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray, contract.ByteArray},
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

	account := ctx.OntClient.Account1
	assetId := helper.AssetId

	address, err := account.ProgramHash.ToAddress()
	if err != nil {
		ctx.LogError("TestGetBalance ProgramHash.ToAddress error:%s", err)
		return false
	}
	balance, err := ctx.Ont.GetBalance(address, assetId)
	if err != nil {
		ctx.LogError("TestGetBalance GetBalance error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{account.ProgramHash.ToArray(), assetId.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetBalance InvokeSmartContract error:%s", err)
		return false
	}

	err = ctx.AssertToInt(res, int(balance))
	if err != nil {
		ctx.LogError("TestGetBalance AssertToInt error:%s", err)
		return false
	}
	return true
}
