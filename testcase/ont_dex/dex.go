package ont_dex

import (
	"fmt"
	"github.com/ONT_TEST/testframework"
	. "github.com/Ontology/core/asset"
	"github.com/Ontology/crypto"
	"math/rand"
	"time"
)

var (
	assetName      = "dex"
	isOntDexInit   = false
	isDexFundInit  = false
	isDexProtoInit = false
	isDexP2PInit   = false
)

func TestOntDex() {
	//testframework.TFramework.RegTestCase("TestOntDexInter", TestOntDexInter)
	testframework.TFramework.RegTestCase("TestDexFundDeposit", TestDexFundDeposit)
	//testframework.TFramework.RegTestCase("TestOrderComplete", TestOrderComplete)
	//testframework.TFramework.RegTestCase("TestOrderCancel", TestOrderCancel)
}

func TestDexFundDeposit(ctx *testframework.TestFrameworkContext) bool {
	err := initTestOntDex(ctx)
	if err != nil {
		ctx.LogError("initTestOntDex error:%s", err)
		ctx.FailNow()
		return false
	}
	availBefer, _, err := DexFund.BalanceOf(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.BalanceOf error:%s", err)
		return false
	}
	amount := 1.0
	err = DexFund.Deposit(ctx, ctx.OntAsset.GetAssetId(assetName), ctx.OntClient.Account1, amount)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.Deposit error:%s", err)
		return false
	}
	_, err = ctx.Ont.WaitForGenerateBlock(30 * time.Second)
	if err != nil {
		ctx.LogError("WaitForGenerateBlock error:%s", err)
	}
	availAfter, _, err := DexFund.BalanceOf(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.AvailBalanceOf error:%s", err)
		return false
	}
	delta := availAfter - availBefer
	if delta != amount {
		ctx.LogError("TestDexFundDeposit error, Delta :%d != %d ", delta, amount)
		return false
	}
	return true
}
//
//func TestMakeBuyOrder(ctx *testframework.TestFrameworkContext) bool{
//	err := initTestOntDex(ctx)
//	if err != nil {
//		ctx.LogError("initTestOntDex error:%s", err)
//		ctx.FailNow()
//		return false
//	}
//
//	buyer := ctx.OntClient.Account1
//	seller := ctx.OntClient.Account2
//	amount := 1.01
//
//	availBefer, totalBefer, err := DexFund.BalanceOf(ctx, buyer)
//	if err != nil {
//		ctx.LogError("TestMakeBuyOrder DexFund.BalanceOf error:%s", err)
//		return false
//	}
//
//	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
//	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
//	if err != nil {
//		ctx.LogError("TestOrderComplete crypto.Sign error:%s", err)
//		return false
//	}
//	err = DexP2P.MakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount)
//	if err != nil {
//		ctx.LogError("TestOrderComplete MakeBuyOrder error:%s", err)
//		return false
//	}
//
//
//
//	DexP2P.MakeBuyOrder(ctx, )
//
//	return true
//}

func TestOrderComplete(ctx *testframework.TestFrameworkContext) bool {
	err := initTestOntDex(ctx)
	if err != nil {
		ctx.LogError("initTestOntDex error:%s", err)
		ctx.FailNow()
		return false
	}
	buyer := ctx.OntClient.Account1
	seller := ctx.OntClient.Account2
	amount := 1.01
	assetId := ctx.OntAsset.GetAssetId(assetName)

	err = DexFund.Deposit(ctx, assetId, buyer, amount)
	if err != nil {
		ctx.LogError("TestOrderComplete DexFund.Deposit error:%s", err)
		return false
	}
	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestOrderComplete crypto.Sign error:%s", err)
		return false
	}
	err = DexP2P.MakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount)
	if err != nil {
		ctx.LogError("TestOrderComplete MakeBuyOrder error:%s", err)
		return false
	}
	err = DexP2P.BuyOrderComplete(ctx, orderId, buyer)
	if err != nil {
		ctx.LogError("TestOrderComplete BuyOrderComplete error:%s", err)
		return false
	}
	return true
}

func TestOrderCancel(ctx *testframework.TestFrameworkContext) bool {
	err := initTestOntDex(ctx)
	if err != nil {
		ctx.LogError("initTestOntDex error:%s", err)
		ctx.FailNow()
		return false
	}
	buyer := ctx.OntClient.Account1
	seller := ctx.OntClient.Account2
	amount := 2.0
	assetId := ctx.OntAsset.GetAssetId(assetName)

	err = DexFund.Deposit(ctx, assetId, buyer, amount)
	if err != nil {
		ctx.LogError("TestOrderCancel DexFund.Deposit error:%s", err)
		return false
	}
	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestOrderCancel crypto.Sign error:%s", err)
		return false
	}
	err = DexP2P.MakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount)
	if err != nil {
		ctx.LogError("TestOrderCancel MakeBuyOrder error:%s", err)
		return false
	}
	err = DexP2P.BuyOrderCancel(ctx, orderId, buyer)
	if err != nil {
		ctx.LogError("TestOrderComplete BuyOrderCancel error:%s", err)
		return false
	}
	return true
}

//func TestFundAdmin(ctx *testframework.TestFrameworkContext) bool {
//	oldAdmin, err := DexFund.GetAdmin(ctx)
//	if err != nil {
//		ctx.LogError("TestFundAdmin GetAdmin error:%s", err)
//		return false
//	}
//
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second)
//	if err != nil{
//		ctx.LogError("TestFundAdmin WaitForGenerateBlock error:%s", err)
//		return false
//	}
//	newAdmin := ctx.OntClient.Account1
//	err = DexFund.ChangeAdmin(ctx, oldAdmin, newAdmin)
//	if err != nil {
//		ctx.LogError("TestFundAdmin ChangeAdmin error:%s", err)
//		return false
//	}
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second)
//	if err != nil{
//		ctx.LogError("TestFundAdmin WaitForGenerateBlock error:%s", err)
//		return false
//	}
//}

func initTestOntDex(ctx *testframework.TestFrameworkContext) error {
	err := initAsset(ctx)
	if err != nil {
		return err
	}
	err = initDexFund(ctx)
	if err != nil {
		return err
	}
	//err = initDexProto(ctx)
	//if err != nil {
	//	return err
	//}
	//err = initDexP2P(ctx)
	//if err != nil {
	//	return err
	//}
	return nil
}

func initAsset(ctx *testframework.TestFrameworkContext) error {
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
	if isDexFundInit {
		return nil
	}
	err := DexFund.Deploy(ctx, ctx.OntClient.Admin)
	if err != nil {
		return fmt.Errorf("TestDexFund DexFund.Deploy error:%s", err)
	}
	assetId := ctx.OntAsset.GetAssetId(assetName)
	err = DexFund.Init(ctx, assetId.ToArray(), ctx.OntClient.Admin, DexProto.CodeHash().ToArray())
	if err != nil {
		return fmt.Errorf("TestDexFund DexFund.Init error:%s", err)
	}
	isDexFundInit = true
	return nil
}

func initDexProto(ctx *testframework.TestFrameworkContext) error {
	if isDexProtoInit {
		return nil
	}
	err := DexProto.Deploy(ctx)
	if err != nil {
		return fmt.Errorf("DexProto.Deploy error:%s", err)
	}
	err = DexProto.Init(ctx, ctx.OntClient.Admin, DexP2P.CodeHash().ToArray())
	if err != nil {
		return fmt.Errorf("DexProto.Init error:%s", err)
	}
	isDexProtoInit = true
	return nil
}

func initDexP2P(ctx *testframework.TestFrameworkContext) error {
	if isDexP2PInit {
		return nil
	}
	err := DexP2P.Deploy(ctx)
	if err != nil {
		return fmt.Errorf("DexP2P.Deploy error:%s", err)
	}
	err = DexP2P.Init(ctx, ctx.OntClient.Admin, 5)
	if err != nil {
		return fmt.Errorf("DexP2P.Init error:%s", err)
	}
	isDexP2PInit = true
	return nil
}

//
//func TestOntDexInter(ctx *testframework.TestFrameworkContext) bool {
//	if !deployDexFund(ctx) {
//		return false
//	}
//	admin := ctx.OntClient.Admin
//	assetName := "TS01"
//	assetPrecise := byte(8)
//	assetType := Token
//	recordType := UTXO
//	asset := ctx.Ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
//	totalAmount := 1000000
//
//	err := RegisterAsset(ctx, asset, totalAmount, admin, admin)
//	if err != nil {
//		ctx.LogError("RegisterAsset error:%s", err)
//		ctx.FailNow()
//		return false
//	}
//	assetId := ctx.OntAsset.GetAssetId(assetName)
//
//	if !dexFundInit(ctx, assetId.ToArray(), admin) {
//		return false
//	}
//
//	buyer := ctx.OntClient.Account1
//	amount := 1000
//	err = IssueAsset(ctx, assetId, admin, buyer, amount)
//	if err != nil {
//		ctx.LogError("IssueAsset error:%s", err)
//		ctx.FailNow()
//		return false
//	}
//
//	if !fundDeposit(ctx, assetId, buyer, amount) {
//		return false
//	}
//
//	//if !setFundCaller(ctx, admin, DExProtoCodeHash){
//	//	return false
//	//}
//	if !deployDexProto(ctx) {
//		return false
//	}
//	if !dexProtoInit(ctx, admin) {
//		return false
//	}
//	if !addProtoCaller(ctx, admin, DEXP2PCodeHashReverse) {
//		return false
//	}
//	if !deployDexP2P(ctx) {
//		return false
//	}
//	if !dexP2PInit(ctx) {
//		return false
//	}
//	seller := ctx.OntClient.Account2
//	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
//	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
//	if err != nil {
//		ctx.LogError("TestOntDexInter crypto.Sign error:%s", err)
//		return false
//	}
//	amount = 10
//	if !fundReceipt(ctx, buyer, amount) {
//		return false
//	}
//	if !makeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
//		return false
//	}
//	if !buyOrderComplete(ctx, orderId, buyer) {
//		return false
//	}
//	orderId = []byte(fmt.Sprint("%d", rand.Int31()))
//	orderSig, err = crypto.Sign(buyer.PrivateKey, orderId)
//	if err != nil {
//		ctx.LogError("TestOntDexInter crypto.Sign error:%s", err)
//		return false
//	}
//	amount = 11
//	if !fundReceipt(ctx, buyer, amount) {
//		return false
//	}
//	if !makeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
//		return false
//	}
//	if !buyOrderCancel(ctx, orderId, buyer) {
//		return false
//	}
//	return true
//}
