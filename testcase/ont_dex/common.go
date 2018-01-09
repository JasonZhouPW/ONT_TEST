package ont_dex

import (
	"encoding/hex"
	"fmt"
	vmtype "github.com/Ontology/vm/neovm/types"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/core/transaction/utxo"
	"reflect"
	"time"
)

var (
	assetName      = "dex"
)

func GetErrorCode(res interface{}) (int, error) {
	v, err := GetRetValue(res, 0, reflect.Int)
	if err != nil {
		return 0, err
	}
	return v.(int), nil
}

func GetRetValue(res interface{}, index int, vType reflect.Kind) (interface{}, error) {
	rt, ok := res.([]interface{})
	if !ok {
		return 0, fmt.Errorf("%s assert to array failed.", res)
	}
	vs, ok := rt[index].(string)
	if !ok {
		return 0, fmt.Errorf("%s assert string")
	}
	v, err := hex.DecodeString(vs)
	if err != nil {
		return 0, fmt.Errorf("hex.DecodeString:%s error:%s", err)
	}
	switch vType {
	case reflect.Int:
		return int(vmtype.ConvertBytesToBigInteger(v).Int64()), nil
	case reflect.String:
		return vs, nil
	}
	return nil, fmt.Errorf("unsupport type:%v", vType)
}

func RegisterAsset(ctx *TestFrameworkContext, asset *asset.Asset, amount int, issuer, controller *account.Account) error {
	regTx, err := ctx.Ont.NewAssetRegisterTransaction(asset, ctx.Ont.MakeAssetAmount(float64(amount)), issuer, controller)
	if err != nil {
		return fmt.Errorf("NewAssetRegisterTransaction error:%s", err)
	}
	txHash, err := ctx.Ont.SendTransaction(controller, regTx)
	if err != nil {
		return fmt.Errorf("SendTransaction error:%s", err)
	}
	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	ctx.OntAsset.RegAsset(txHash, asset)
	return nil
}

func IssueAsset(ctx *TestFrameworkContext, assetId common.Uint256, controller, toAccount *account.Account, amount int) error {
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(float64(amount)),
		AssetID:     assetId,
		ProgramHash: toAccount.ProgramHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		return fmt.Errorf("NewIssueAssetTransaction error:%s", err)
	}
	_, err = ctx.Ont.SendTransaction(controller, issueTx)
	if err != nil {
		return fmt.Errorf("SendTransaction error:%s", err)
	}
	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return nil
}

func InitAsset(ctx *TestFrameworkContext, toAccount *account.Account) error {
	admin := ctx.OntClient.Admin
	assetPrecise := byte(8)
	assetType := asset.Token
	recordType := asset.UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	totalAmount := 1000000
	err := RegisterAsset(ctx, asset, totalAmount, admin, admin)
	if err != nil {
		return fmt.Errorf("RegisterAsset error:%s", err)
	}

	assetId := ctx.OntAsset.GetAssetId(assetName)
	amount := 10000
	err = IssueAsset(ctx, assetId, admin, toAccount, amount)
	if err != nil {
		return fmt.Errorf("IssueAsset to Account%v error:%s",toAccount, err)
	}
	return nil
}

func GetAssetName()string{
	return assetName
}