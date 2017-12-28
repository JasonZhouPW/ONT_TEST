package asset

import (
	"github.com/Ontology/common"
	. "github.com/Ontology/core/asset"
	"github.com/Ontology/core/transaction/payload"
	. "github.com/ONT_TEST/ont"
	. "github.com/ONT_TEST/testframework"
	"time"
)

func TestRegisterAssetTransaction(ctx *TestFrameworkContext) bool {
	assetName := "TS01"
	assetPrecise := byte(4)
	assetType := Token
	recordType := UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	assetAmount := ctx.Ont.MakeAssetAmount(20000)
	if !testRegisterAssetTransaction(asset, assetAmount, ctx) {
		ctx.LogError("TestRegisterAssetTransaction Asset:%+v Amount:%v test failed.",
			asset, assetAmount)
		ctx.FailNow()
		return false
	}

	assetName = "TS02"
	assetPrecise = byte(8)
	assetType = Share
	recordType = UTXO
	asset = ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	assetAmount = ctx.Ont.MakeAssetAmount(100000)
	if !testRegisterAssetTransaction(asset, assetAmount, ctx) {
		ctx.LogError("TestRegisterAssetTransaction Asset:%+v Amount:%v test failed.",
		asset, assetAmount)
		ctx.FailNow()
		return false
	}

	return true
}

func TestRegisterAssetPreciseTransaction(ctx *TestFrameworkContext) bool {
	assetName := "TS01"
	assetPrecise := byte(4)
	assetType := Token
	recordType := UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	assetAmount := ctx.Ont.MakeAssetAmount(100.00001)
	regTx, err := ctx.Ont.NewAssetRegisterTransaction(asset, assetAmount, ctx.OntClient.Admin, ctx.OntClient.Admin)
	if err != nil {
		ctx.LogError("NewAssetRegisterTransaction Asset:%+v Amount:%v Admin:%+v Account:%+v error:%s",
			asset,
			assetAmount,
			ctx.OntClient.Admin,
			ctx.OntClient.Admin,
			err)

		ctx.FailNow()
		return false
	}

	//Should failed
	_, err = ctx.Ont.SendTransaction(ctx.OntClient.Admin, regTx)
	if err == nil {
		ctx.LogError("SendTransaction AssetRegisterTransaction should failed error:%s", err)
		return false
	}
	return true
}

func TestRegisterAssetMaxPreciseTransaction(ctx *TestFrameworkContext) bool {
	assetName := "TS01"
	assetPrecise := byte(9)
	assetType := Token
	recordType := UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	assetAmount := ctx.Ont.MakeAssetAmount(10000)

	regTx, err := ctx.Ont.NewAssetRegisterTransaction(asset, assetAmount, ctx.OntClient.Admin, ctx.OntClient.Admin)
	if err != nil {
		ctx.LogError("NewAssetRegisterTransaction Asset:%+v Amount:%v Admin:%+v Account:%+v error:%s",
			asset,
			assetAmount,
			ctx.OntClient.Admin,
			ctx.OntClient.Admin,
			err)
		return false
	}
	//Should failed
	_, err = ctx.Ont.SendTransaction(ctx.OntClient.Admin, regTx)
	if err == nil {
		ctx.LogError("SendTransaction AssetRegisterTransaction should failed err:s", err)
		return false
	}
	return true
}

func testRegisterAssetTransaction(asset *Asset, assetAmount common.Fixed64, ctx *TestFrameworkContext) bool {
	regTx, err := ctx.Ont.NewAssetRegisterTransaction(asset, assetAmount, ctx.OntClient.Admin, ctx.OntClient.Admin)
	if err != nil {
		ctx.LogError("NewAssetRegisterTransaction Asset:%+v Amount:%v Admin:%+v Account:%+v error:%s",
			asset,
			assetAmount,
			ctx.OntClient.Admin,
			ctx.OntClient.Admin,
			err)

		ctx.FailNow()
		return false
	}

	txHash, err := ctx.Ont.SendTransaction(ctx.OntClient.Admin, regTx)
	if err != nil {
		ctx.LogError("SendTransaction AssetRegisterTransaction error:%s", err)
		ctx.FailNow()
		return false
	}

	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
		ctx.FailNow()
		return false
	}

	regTx2, err := ctx.Ont.GetTransaction(txHash)
	if err != nil {
		ctx.LogError("GetTransaction Hash:%x error:%s", txHash, err)
		return false
	}

	regAssetPayload := regTx2.Payload.(*payload.RegisterAsset)
	asset2 := regAssetPayload.Asset
	if !AssetEqualTo(asset, asset2) || regAssetPayload.Amount != assetAmount {
		ctx.LogError("Asset get from transaction not equal.")
		return false
	}

	ctx.OntAsset.RegAsset(txHash, asset)
	return true
}
