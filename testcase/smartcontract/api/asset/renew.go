package asset

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestRenew(ctx *testframework.TestFrameworkContext) bool {
	err := InitAsset(ctx)
	if err != nil {
		ctx.LogError("TestRenew CreateAsset error:%s", err)
		return false
	}

	code := "54c56b6c766b00527ac46c766b51527ac4616c766b00c36168174e656f2e426c6f636b636861696e2e47657441737365746c766b52527ac46c766b52c36c766b51c3617c680f4e656f2e41737365742e52656e65776c766b53527ac46203006c766b53c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray,contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"TestRenew",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestRenew DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestRenew WaitForGenerateBlock error:%s", err)
		return false
	}

	height, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("TestRenew GetBlockCount error:%s", err)
		return false
	}
	years := 1
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{GetAssetId().ToArray(), years},
	)
	if err != nil {
		ctx.LogError("TestRenew InvokeSmartContract error:%s", err)
		return false
	}

	blockNumOneYear := 2000000
	err = ctx.AssertToInt(res, (years * blockNumOneYear) + int(height-1))
	if err != nil {
		ctx.LogError("TestRenew test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using System.Numerics;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static BigInteger Main(byte[] assetId, BigInteger years)
    {
        Asset asset = Blockchain.GetAsset(assetId);
        return asset.Renew((byte)years);
    }
}
*/
