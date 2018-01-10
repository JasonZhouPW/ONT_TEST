package asset

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/**
using Neo.SmartContract.Framework;
using System.Numerics;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static object[] Main(byte[] assetId)
    {
        Asset asset = Blockchain.GetAsset(assetId);
        object[] ret = new object[8];
        ret[0] = asset.AssetId;
        ret[1] = asset.AssetType;
        ret[2] = asset.Amount;
        ret[3] = asset.Precision;
        ret[4] = asset.Issuer;
        ret[5] = asset.Owner;
        ret[6] = asset.Admin;
        ret[7] = asset.Available;
        return ret;
    }
}
*/

func TestAsset(ctx *testframework.TestFrameworkContext) bool {
	assetType := asset.Token
	assetName := "tst"
	assetAmount := 1000000.0
	assetPrecision := byte(8)
	assetOwner := ctx.OntClient.Admin
	assetAdmin := ctx.OntClient.Admin
	assetIssuer := ctx.OntClient.Account1

	assetId, err := CreateAsset(
		ctx,
		assetType,
		assetName,
		assetAmount,
		assetPrecision,
		assetOwner,
		assetAdmin,
		assetIssuer)
	if err != nil {
		ctx.LogError("CreateAsset error:%s", err)
		return false
	}

	code := "54c56b6c766b00527ac4616c766b00c36168174e656f2e426c6f636b636861696e2e47657441737365746c766b51527ac458c56c766b52527ac46c766b52c3006c766b51c36168144e656f2e41737365742e47657441737365744964c46c766b52c3516c766b51c36168164e656f2e41737365742e476574417373657454797065c46c766b52c3526c766b51c36168134e656f2e41737365742e476574416d6f756e74c46c766b52c3536c766b51c36168164e656f2e41737365742e476574507265636973696f6ec46c766b52c3556c766b51c36168124e656f2e41737365742e4765744f776e6572c46c766b52c3566c766b51c36168124e656f2e41737365742e47657441646d696ec46c766b52c3576c766b51c36168164e656f2e41737365742e476574417661696c61626c65c46c766b52c36c766b53527ac46203006c766b53c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Array),
		"TestAsset",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestAsset DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestAsset WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{assetId.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestAsset InvokeSmartContract error:%s", err)
		return false
	}

	ret, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestAsset asset res to []interface{} error:%s", err)
		return false
	}

	assId := ret[0]
	err = ctx.AssertToByteArray(assId, assetId.ToArray())
	if err != nil {
		ctx.LogError("TestAsset AssetId AssertToByteArray error:%s", err)
		return false
	}

	assType := ret[1]
	err = ctx.AssertToInt(assType, int(assetType))
	if err != nil {
		ctx.LogError("TestAsset AssetType AssertToInt error:%s", err)
		return false
	}

	amount := ret[2]
	err = ctx.AssertToInt(amount, int(ctx.Ont.MakeAssetAmount(assetAmount)))
	if err != nil {
		ctx.LogError("TestGetAmount AssetAmount AssertToInt error:%s", err)
		return false
	}

	precision := ret[3]
	err = ctx.AssertToInt(precision, int(assetPrecision))
	if err != nil {
		ctx.LogError("TestGetAmount AssetPrecision AssertToInt error:%s", err)
		return false
	}

	//issuer := ret[4]

	owner := ret[5]
	ownerPk, _ := assetOwner.PublicKey.EncodePoint(true)
	err = ctx.AssertToByteArray(owner, ownerPk)
	if err != nil {
		ctx.LogError("TestGetAmount AssetOwner AssertToByteArray error:%s", err)
		return false
	}

	admin := ret[6]
	err = ctx.AssertToByteArray(admin, assetAdmin.ProgramHash.ToArray())
	if err != nil {
		ctx.LogError("TestGetAmount AssetOwner AssertToByteArray error:%s", err)
		return false
	}

	avail := ret[7]
	err = ctx.AssertToInt(avail, 0)
	if err != nil {
		ctx.LogError("TestGetAmount AssetOwner Available error:%s", err)
		return false
	}
	return true
}
