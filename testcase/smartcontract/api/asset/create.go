package asset

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static byte[] Main(byte[] pk)
    {
        Asset asset = Asset.Create(0, "test", 100000000, 8, pk, pk, pk);
        return asset.AssetId;
    }
}
 */

func TestAssetCreate(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6c766b00527ac4610004746573740400e1f505586c766b00c36c766b00c36c766b00c36156795179587275517275557952795772755272755479537956727553727568104e656f2e41737365742e4372656174656c766b51527ac46c766b51c36168144e656f2e41737365742e476574417373657449646c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestAssetCreate",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	pk := ctx.OntClient.Account1.PublicKey
	encodedPubKey, _ := pk.EncodePoint(true)
	if err != nil {
		ctx.LogError("TestAssetCreate DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestAssetCreate WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{encodedPubKey},
	)
	if err != nil {
		ctx.LogError("TestAssetCreate InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestAssetCreate :%+v ", res)
	return true
}