package helper

import (
	"fmt"
	"github.com/ONT_TEST/testcase/ont_dex"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/core/transaction/utxo"
)

var (
	AssetName = "dex"
	AssetId   common.Uint256
	isInit = false
)

func InitAsset(ctx *TestFrameworkContext, toAccount *account.Account) error {
	if isInit{
		return nil
	}
	admin := ctx.OntClient.Admin
	assetPrecise := byte(8)
	assetType := asset.Token
	recordType := asset.UTXO
	asset := ctx.Ont.CreateAsset(AssetName, assetPrecise, assetType, recordType)
	totalAmount := 1000000
	err := ont_dex.RegisterAsset(ctx, asset, totalAmount, admin, admin)
	if err != nil {
		return fmt.Errorf("RegisterAsset error:%s", err)
	}

	assetId := ctx.OntAsset.GetAssetId(AssetName)
	amount := 10000
	err = ont_dex.IssueAsset(ctx, assetId, admin, toAccount, amount)
	if err != nil {
		return fmt.Errorf("IssueAsset to Account%v error:%s", toAccount, err)
	}
	AssetId = assetId
	isInit = true
	return nil
}

func Transfer(ctx *TestFrameworkContext, assetId common.Uint256, from, to *account.Account, amount float64) (common.Uint256, error) {
	unspents, err := ctx.Ont.GetUnspendOutput(assetId, from.ProgramHash)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("GetUnspendOutput error:%s", err)
	}
	if unspents == nil {
		return common.Uint256{}, fmt.Errorf("GetUnspendOutput return nil")
	}

	assAmount := ctx.Ont.MakeAssetAmount(amount)
	txInputs := make([]*utxo.UTXOTxInput, 0, 1)
	txOutputs := make([]*utxo.TxOutput, 0, 2)

	for _, unspent := range unspents {
		if unspent.Value < assAmount {
			continue
		}
		input := &utxo.UTXOTxInput{
			ReferTxID:          unspent.Txid,
			ReferTxOutputIndex: uint16(unspent.Index),
		}
		txInputs = append(txInputs, input)
		output := &utxo.TxOutput{
			AssetID:     assetId,
			Value:       assAmount,
			ProgramHash: to.ProgramHash,
		}
		txOutputs = append(txOutputs, output)
		//dibs output
		dibs := unspent.Value - assAmount
		if dibs > 0 {
			output2 := &utxo.TxOutput{
				AssetID:     output.AssetID,
				Value:       dibs,
				ProgramHash: from.ProgramHash,
			}
			txOutputs = append(txOutputs, output2)
		}
		break
	}
	if len(txInputs) == 0 {
		return common.Uint256{}, fmt.Errorf("TxInput is nil")
	}

	transferTx, err := ctx.Ont.NewTransferAssetTransaction(txInputs, txOutputs)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("NewTransferAssetTransaction error:%s", err)
	}

	txHash, err := ctx.Ont.SendTransaction(ctx.OntClient.Account1, transferTx)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("SendTransaction error:%s", err)
	}

	if err != nil {
		return common.Uint256{}, err
	}

	return txHash, nil
}
