package asset

import (
	"github.com/ONT_TEST/testframework"
	"fmt"
	"github.com/Ontology/common"
	"github.com/Ontology/core/transaction/utxo"
	"time"
)

func TestIssueAssetTransaction(ctx *testframework.TestFrameworkContext) bool {
	assetName := "TS01"
	assetId := ctx.OntAsset.GetAssetId(assetName)
	empty := common.Uint256{}
	if assetId == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetName)
		ctx.FailNow()
		return false
	}
	//asset := ctx.OntAsset.GetAssetByName(assetName)
	programHash, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(100),
		AssetID:     assetId,
		ProgramHash: programHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		ctx.LogError("NewIssueAssetTransaction error:%s", err)
		return false
	}
	txHash, err := ctx.Ont.SendTransaction(ctx.OntClient.Admin, issueTx)
	if err != nil {
		ctx.LogError("SendTransaction error:%s", err)
		return false
	}

	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
		ctx.FailNow()
		return false
	}
	issueTx2, err := ctx.Ont.GetTransaction(txHash)
	if err != nil {
		ctx.LogError("GetTransaction TxHash:%x error:%s", txHash, err)
		return false
	}
	if len(issueTx2.Outputs) == 0 {
		ctx.LogError("GetTransaction Outputs error")
		return false
	}

	txOutputsRes := issueTx2.Outputs
	ok, err := checkIssueAssetTxResult(txOutputs, txOutputsRes)
	if err != nil {
		ctx.LogError("checkIssueAssetTxResult error:%s", err)
		return false
	}
	return ok
}

func TestIssueAssetMutiTransaction(ctx *testframework.TestFrameworkContext) bool {
	empty := common.Uint256{}
	assetName1 := "TS01"
	assetId1 := ctx.OntAsset.GetAssetId(assetName1)
	if assetId1 == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetId1)
		ctx.FailNow()
		return false
	}
	//asset1 := ctx.OntAsset.GetAssetByName(assetName1)
	assetName2 := "TS02"
	assetId2 := ctx.OntAsset.GetAssetId(assetName2)
	if assetId1 == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetId2)
		ctx.FailNow()
		return false
	}
	//asset2 := ctx.OntAsset.GetAssetByName(assetName2)
	programHash1, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	programHash2, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account2)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	txOutput1 := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(100),
		AssetID:     assetId1,
		ProgramHash: programHash1,
	}
	txOutput2 := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(100),
		AssetID:     assetId2,
		ProgramHash: programHash2,
	}
	txOutput3 := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(100),
		AssetID:     assetId2,
		ProgramHash: programHash1,
	}
	txOutputs := []*utxo.TxOutput{txOutput1, txOutput2, txOutput3}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		ctx.LogError("NewIssueAssetTransaction error:%s", err)
		return false
	}
	txHash, err := ctx.Ont.SendTransaction(ctx.OntClient.Admin, issueTx)
	if err != nil {
		ctx.LogError("SendTransaction error:%s", err)
		return false
	}

	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
		ctx.FailNow()
		return false
	}
	issueTx2, err := ctx.Ont.GetTransaction(txHash)
	if err != nil {
		ctx.LogError("GetTransaction TxHash:%x error:%s", txHash, err)
		return false
	}
	if len(issueTx2.Outputs) == 0 {
		ctx.LogError("GetTransaction Outputs error")
		return false
	}

	txOutputsRes := issueTx2.Outputs
	ok, err := checkIssueAssetTxResult(txOutputs, txOutputsRes)
	if err != nil {
		ctx.LogError("checkIssueAssetTxResult error:%s", err)
		return false
	}
	return ok
}

func TestIssueAssetOverAmountTransaction(ctx *testframework.TestFrameworkContext) bool {
	assetName := "TS01"
	assetId := ctx.OntAsset.GetAssetId(assetName)
	empty := common.Uint256{}
	if assetId == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetName)
		ctx.FailNow()
		return false
	}
	//asset := ctx.OntAsset.GetAssetByName(assetName)
	programHash, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(10000000),
		AssetID:     assetId,
		ProgramHash: programHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		ctx.LogError("NewIssueAssetTransaction error:%s", err)
		return false
	}

	//Should failed
	_, err = ctx.Ont.SendTransaction(ctx.OntClient.Admin, issueTx)
	if err == nil {
		ctx.LogError("SendTransaction should failed. err:%s", err)
		return false
	}
	return true
}

func TestIssueAssetNegAmountTransaction(ctx *testframework.TestFrameworkContext) bool {
	assetName := "TS01"
	assetId := ctx.OntAsset.GetAssetId(assetName)
	empty := common.Uint256{}
	if assetId == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetName)
		ctx.FailNow()
		return false
	}
	//asset := ctx.OntAsset.GetAssetByName(assetName)
	programHash, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(-1),
		AssetID:     assetId,
		ProgramHash: programHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		ctx.LogError("NewIssueAssetTransaction error:%s", err)
		return false
	}

	//Should failed
	_, err = ctx.Ont.SendTransaction(ctx.OntClient.Admin, issueTx)
	if err == nil {
		ctx.LogError("SendTransaction error should failed. err:%s", err)
		return false
	}
	return true
}

func TestIssueAssetPreciseTransaction(ctx *testframework.TestFrameworkContext) bool {
	assetName := "TS01"
	assetId := ctx.OntAsset.GetAssetId(assetName)
	empty := common.Uint256{}
	if assetId == empty {
		ctx.LogError("AssetName:%s doesnot exist", assetName)
		ctx.FailNow()
		return false
	}
	//asset := ctx.OntAsset.GetAssetByName(assetName)
	programHash, err := ctx.Ont.GetAccountProgramHash(ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("GetProgramHash error:%s", err)
		return false
	}
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(100.00001),
		AssetID:     assetId,
		ProgramHash: programHash,
	}
	ctx.LogInfo("Amount:%v", output.Value)
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		ctx.LogError("NewIssueAssetTransaction error:%s", err)
		return false
	}

	//Should failed
	_, err = ctx.Ont.SendTransaction(ctx.OntClient.Admin, issueTx)
	if err == nil {
		ctx.LogError("SendTransaction error.Transaction shuld be rejected err:%s", err)
		return false
	}
	return true
}

func checkIssueAssetTxResult(txOutputs, txOutputsRes []*utxo.TxOutput) (bool, error) {
	if len(txOutputs) != len(txOutputsRes) {
		return false, fmt.Errorf("len(txOutputs):%v != len(txOutputsRes):%v", len(txOutputs), len(txOutputsRes))
	}
	for i, txOutputRes := range txOutputsRes {
		txOutput := txOutputs[i]
		if txOutput.ProgramHash != txOutputRes.ProgramHash &&
			txOutput.AssetID != txOutputRes.AssetID &&
			txOutput.Value != txOutputRes.Value {
			return false, fmt.Errorf("IssueAssetTransaction ProgramHash:%x != %x AssetID:%x != %x Value:%v != %v",
				txOutputRes.ProgramHash,
				txOutput.ProgramHash,
				txOutputRes.AssetID,
				txOutput.AssetID,
				txOutputRes.Value,
				txOutput.Value,
			)
		}
	}
	return true, nil
}
