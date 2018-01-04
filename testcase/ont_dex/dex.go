package ont_dex

import (
	"github.com/ONT_TEST/testframework"
	. "github.com/Ontology/core/asset"
	//"github.com/Ontology/crypto"
	//"math/rand"
	"fmt"
	"github.com/Ontology/common"
	"math/rand"
	"github.com/Ontology/crypto"
)

var (
	assetName     = "TS01"
	isOntDexInit  = false
	isDexFundInit = false
)

func TestOntDex() {
	testframework.TFramework.RegTestCase("TestOntDexInter", TestOntDexInter)
	//testframework.TFramework.RegTestCase("TestDexFundDeposit", TestDexFundDeposit)
}

func TestDexFundDeposit(ctx *testframework.TestFrameworkContext) bool {
	err := initDexFund(ctx)
	if err != nil {
		ctx.LogError("TestDexFundDeposit initDexFund error:%s", err)
		ctx.FailNow()
		return false
	}
	amount := 1
	err = DexFund.Deposit(ctx, ctx.OntAsset.GetAssetId(assetName), ctx.OntClient.Account1, amount)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.Deposit error:%s", err)
		return false
	}

	balance, err := DexFund.AvailBalanceOf(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.AvailBalanceOf error:%s", err)
		return false
	}

	b := int(ctx.Ont.GetRawAssetAmount(common.Fixed64(balance)))
	if b != amount {
		ctx.LogError("AvailBalance error, balance:%d != %d ", b, amount)
		return false
	}
	return true
}

func initTestOntDex(ctx *testframework.TestFrameworkContext) error {
	if isOntDexInit {
		return nil
	}
	admin := ctx.OntClient.Admin
	assetPrecise := byte(8)
	assetType := Token
	recordType := UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	totalAmount := 1000000
	err := RegisterAsset(ctx, asset, totalAmount, admin, admin)
	if err != nil {
		return fmt.Errorf("RegisterAsset error:%s", err)
	}

	assetId := ctx.OntAsset.GetAssetId(assetName)
	amount := 10000
	err = IssueAsset(ctx, assetId, admin, ctx.OntClient.Account1, amount)
	if err != nil {
		return fmt.Errorf("IssueAsset to Account1 error:%s", err)
	}
	isOntDexInit = true
	return nil
}

func initDexFund(ctx *testframework.TestFrameworkContext) error {
	err := initTestOntDex(ctx)
	if err != nil {
		return fmt.Errorf("initDexFund initTestOntDex error:%s", err)
	}
	if isDexFundInit {
		return nil
	}
	err = DexFund.Deploy(ctx, ctx.OntClient.Admin)
	if err != nil {
		return fmt.Errorf("TestDexFund DexFund.Deploy error:%s", err)
	}
	assetId := ctx.OntAsset.GetAssetId(assetName)
	err = DexFund.Init(ctx, assetId.ToArray(), ctx.OntClient.Admin)
	if err != nil {
		return fmt.Errorf("TestDexFund DexFund.Init error:%s", err)
	}
	isDexFundInit = true
	return nil
}

func TestOntDexInter(ctx *testframework.TestFrameworkContext) bool {
	if !deployDexFund(ctx) {
		return false
	}
	admin := ctx.OntClient.Admin
	assetName := "TS01"
	assetPrecise := byte(8)
	assetType := Token
	recordType := UTXO
	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	totalAmount := 1000000

	err := RegisterAsset(ctx, asset, totalAmount, admin, admin)
	if err != nil{
		ctx.LogError("RegisterAsset error:%s", err)
		ctx.FailNow()
		return false
	}
	assetId := ctx.OntAsset.GetAssetId(assetName)

	if !dexFundInit(ctx, assetId.ToArray(), admin) {
		return false
	}

	buyer := ctx.OntClient.Account1
	amount := 1000
	err  = IssueAsset(ctx, assetId, admin, buyer, amount)
	if err != nil{
		ctx.LogError("IssueAsset error:%s", err)
		ctx.FailNow()
		return false
	}

	if !fundDeposit(ctx, assetId, buyer, amount){
		return false
	}

	//if !setFundCaller(ctx, admin, DExProtoCodeHash){
	//	return false
	//}
	if !deployDexProto(ctx) {
		return false
	}
	if !dexProtoInit(ctx, admin) {
		return false
	}
	if !addProtoCaller(ctx, admin, DEXP2PCodeHashReverse) {
		return false
	}
	if !deployDexP2P(ctx) {
		return false
	}
	if !dexP2PInit(ctx) {
		return false
	}
	seller := ctx.OntClient.Account2
	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestOntDexInter crypto.Sign error:%s", err)
		return false
	}
	amount = 10
	if !fundReceipt(ctx, buyer, amount) {
		return false
	}
	if !makeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
		return false
	}
	if !buyOrderComplete(ctx, orderId, buyer) {
		return false
	}
	orderId = []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err = crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestOntDexInter crypto.Sign error:%s", err)
		return false
	}
	amount = 11
	if !fundReceipt(ctx, buyer, amount) {
		return false
	}
	if !makeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
		return false
	}
	if !buyOrderCancel(ctx, orderId, buyer) {
		return false
	}
	return true
}
