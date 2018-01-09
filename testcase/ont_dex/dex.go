package ont_dex

import (
	"fmt"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/crypto"
	"math/rand"
	"time"
)

var (
	isOntDexInit   = false
)

func TestOntDex() {
	//testframework.TFramework.RegTestCase("TestDexFundDeposit", TestDexFundDeposit)
	//testframework.TFramework.RegTestCase("TestMakeBuyOrder", TestMakeBuyOrder)
	//testframework.TFramework.RegTestCase("TestOrderComplete", TestOrderComplete)
	//testframework.TFramework.RegTestCase("TestOrderCancel", TestOrderCancel)
	testframework.TFramework.RegTestCase("TestSellerTryCloseOrder", TestSellerTryCloseOrder)
}

func TestDexFundDeposit(ctx *testframework.TestFrameworkContext) bool {
	err := initTestOntDex(ctx)
	if err != nil {
		ctx.LogError("initTestOntDex error:%s", err)
		ctx.FailNow()
		return false
	}
	availBefer, totalBefer, err := DexFund.BalanceOf(ctx, ctx.OntClient.Account1)
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

	availAfter, totalAfter, err := DexFund.BalanceOf(ctx, ctx.OntClient.Account1)
	if err != nil {
		ctx.LogError("TestDexFund DexFund.AvailBalanceOf error:%s", err)
		return false
	}
	delta := availAfter - availBefer
	if delta != amount {
		ctx.LogError("TestDexFundDeposit error, Avail Delta :%v != %v ", delta, amount)
		return false
	}

	delta = totalAfter - totalBefer
	if delta != amount {
		ctx.LogError("TestDexFundDeposit error, Total Delta :%v != %v ", delta, amount)
		return false
	}
	return true
}

func TestMakeBuyOrder(ctx *testframework.TestFrameworkContext) bool {
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
		ctx.LogError("TestOrderComplete DexFund.Deposit error:%s", err)
		return false
	}
	availBefor, totalBefor, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestMakeBuyOrder DexFund.BalanceOf error:%s", err)
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

	availAfter, totalAfter, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestMakeBuyOrder DexFund.BalanceOf error:%s", err)
		return false
	}

	if availAfter != (availBefor - amount) {
		ctx.LogError("MakeBuyOrder availfund: %v != %v", availAfter, (availBefor - amount))
		return false
	}

	if totalAfter != (totalBefor - amount){
		ctx.LogError("MakeBuyOrder totalfund: %v != %v", totalAfter, totalBefor)
		return false
	}
	return true
}

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

	buyerAvailBeforComplete, buyerTotalBeforComplete, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestOrderComplete DexFund.BalanceOf error:%s", err)
		return false
	}

	sellerAvailBeforComplete, sellerTotalBeforComplete, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestOrderComplete DexFund.BalanceOf error:%s", err)
		return false
	}

	err = DexP2P.BuyOrderComplete(ctx, orderId, buyer)
	if err != nil {
		ctx.LogError("TestOrderComplete BuyOrderComplete error:%s", err)
		return false
	}

	sellerAvailAfterComplete, sellerTotalAfterComplete, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestOrderComplete DexFund.BalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("sellerAvailAfterComplete:%v, sellerTotalAfterComplete:%v", sellerAvailAfterComplete, sellerTotalAfterComplete)

	if sellerAvailAfterComplete != (sellerAvailBeforComplete + amount) {
		ctx.LogError("TestOrderComplete sellerAvailFundAfterComplete %v != %v", sellerAvailAfterComplete, sellerAvailBeforComplete+amount)
		return false
	}
	if sellerTotalAfterComplete != (sellerTotalBeforComplete + amount) {
		ctx.LogError("TestOrderComplete sellerTotalFundAfterComplete %v != %v", sellerTotalAfterComplete, sellerTotalBeforComplete+amount)
		return false
	}

	buyerAvailAfterComplete, buyerTotalAfterComplete, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestOrderComplete DexFund.BalanceOf error:%s", err)
		return false
	}

	ctx.LogInfo("buyerAvailAfterComplete:%v, buyerTotalAfterComplete:%v", buyerAvailAfterComplete, buyerTotalAfterComplete)

	if buyerAvailAfterComplete != buyerAvailBeforComplete {
		ctx.LogError("TestOrderComplete buyerAvailAfterComplete %v != %v", buyerAvailAfterComplete, buyerAvailBeforComplete)
		return false
	}

	if buyerTotalAfterComplete != (buyerTotalBeforComplete - amount) {
		ctx.LogError("TestOrderComplete buyerTotalAfterComplete %v != %v", buyerTotalAfterComplete, buyerTotalBeforComplete-amount)
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

	buyerAvailBeforCancel, buyerTotalBeforCancel, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestOrderCancel DexFund.BalanceOf error:%s", err)
		return false
	}

	sellerAvailBeforCancel, sellerTotalBeforCancel, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestOrderCancel DexFund.BalanceOf error:%s", err)
		return false
	}

	err = DexP2P.BuyOrderCancel(ctx, orderId, buyer)
	if err != nil {
		ctx.LogError("TestOrderCancel BuyOrderCancel error:%s", err)
		return false
	}

	sellerAvailAfterCancel, sellerTotalAfterCancel, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestOrderCancel DexFund.BalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("sellerAvailAfterCancel:%v, sellerTotalAfterCancel:%v", sellerAvailAfterCancel, sellerTotalAfterCancel)
	if sellerAvailAfterCancel != sellerAvailBeforCancel {
		ctx.LogError("TestOrderCancel sellerAvailFundAfterCancel %v != %v", sellerAvailAfterCancel, sellerAvailBeforCancel)
		return false
	}
	if sellerTotalAfterCancel != sellerTotalBeforCancel {
		ctx.LogError("TestOrderCancel sellerTotalFundAfterCancel %v != %v", sellerTotalAfterCancel, sellerTotalBeforCancel)
		return false
	}

	buyerAvailAfterCancel, buyerTotalAfterCancel, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestOrderCancel DexFund.BalanceOf error:%s", err)
		return false
	}
	ctx.LogInfo("buyerAvailAfterCancel:%v, buyerTotalAfterCancel:%v", buyerAvailAfterCancel, buyerTotalAfterCancel)
	if buyerAvailAfterCancel != buyerAvailBeforCancel+amount {
		ctx.LogError("TestOrderCancel buyerAvailAfterCancel %v != %v", buyerAvailAfterCancel, buyerAvailBeforCancel+amount)
		return false
	}

	if buyerTotalAfterCancel != buyerTotalBeforCancel {
		ctx.LogError("TestOrderCancel buyerTotalAfterCancel %v != %v", buyerTotalAfterCancel, buyerTotalBeforCancel)
		return false
	}

	return true
}

func TestSellerTryCloseOrder(ctx *testframework.TestFrameworkContext) bool {
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
		ctx.LogError("TestSellerTryCloseOrder DexFund.Deposit error:%s", err)
		return false
	}
	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder crypto.Sign error:%s", err)
		return false
	}

	buyerAvailBeforTryClose, buyerTotalBeforTryClose, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder DexFund.BalanceOf error:%s", err)
		return false
	}

	sellerAvailBeforTryClose, sellerTotalBeforTryClose, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder DexFund.BalanceOf error:%s", err)
		return false
	}

	err = DexP2P.MakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder MakeBuyOrder error:%s", err)
		return false
	}

	//should failed
	err = DexP2P.SellerTryCloseOrder(ctx, orderId, seller)
	if err == nil {
		ctx.LogError("TestSellerTryCloseOrder SellerTryCloseOrder should failed")
		return false
	}

	lockTime, err := DexP2P.GetOrderLockTime(ctx)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder GetOrderLockTime error:%s", err)
		return false
	}

	time.Sleep(time.Second * time.Duration(lockTime))

	err = DexP2P.SellerTryCloseOrder(ctx, orderId, seller)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder SellerTryCloseOrder error:%s", err)
		return false
	}

	sellerAvailAfterTryClose, sellerTotalAfterTryClose, err := DexFund.BalanceOf(ctx, seller)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder DexFund.BalanceOf error:%s", err)
		return false
	}

	ctx.LogInfo("TestSellerTryCloseOrder sellerAvailAfterTryClose:%v, sellerTotalAfterTryClose:%v", sellerAvailAfterTryClose, sellerTotalAfterTryClose)

	if sellerAvailAfterTryClose != (sellerAvailBeforTryClose + amount) {
		ctx.LogError("TestSellerTryCloseOrder sellerAvailFundAfterTryClose %v != %v", sellerAvailAfterTryClose, sellerAvailBeforTryClose+amount)
		return false
	}
	if sellerTotalAfterTryClose != (sellerTotalBeforTryClose + amount) {
		ctx.LogError("TestSellerTryCloseOrder sellerTotalFundAfterTryClose %v != %v", sellerTotalAfterTryClose, sellerTotalBeforTryClose+amount)
		return false
	}

	buyerAvailAfterTryClose, buyerTotalAfterTryClose, err := DexFund.BalanceOf(ctx, buyer)
	if err != nil {
		ctx.LogError("TestSellerTryCloseOrder DexFund.BalanceOf error:%s", err)
		return false
	}

	ctx.LogInfo("TestSellerTryCloseOrder buyerAvailAfterTryClose:%v, buyerTotalAfterTryClose:%v", buyerAvailAfterTryClose, buyerTotalAfterTryClose)

	if buyerAvailAfterTryClose != buyerAvailBeforTryClose {
		ctx.LogError("TestSellerTryCloseOrder buyerAvailAfterTryClose %v != %v", buyerAvailAfterTryClose, buyerAvailBeforTryClose)
		return false
	}

	if buyerTotalAfterTryClose != (buyerTotalBeforTryClose - amount) {
		ctx.LogError("TestSellerTryCloseOrder buyerTotalAfterTryClose %v != %v", buyerTotalAfterTryClose, (buyerTotalBeforTryClose-amount))
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
	if isOntDexInit {
		return nil
	}
	err := InitAsset(ctx, ctx.OntClient.Account1)
	if err != nil {
		return err
	}

	err = DexFund.Deploy(ctx, ctx.OntClient.Admin)
	if err != nil {
		return fmt.Errorf("TestDexFund DexFund.Deploy error:%s", err)
	}
	err = DexProto.Deploy(ctx)
	if err != nil {
		return fmt.Errorf("DexProto.Deploy error:%s", err)
	}
	err = DexP2P.Deploy(ctx)
	if err != nil {
		return fmt.Errorf("DexP2P.Deploy error:%s", err)
	}

	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}

	assetId := ctx.OntAsset.GetAssetId(assetName)
	err = DexFund.Init(ctx, assetId.ToArray(), ctx.OntClient.Admin, DexProto.CodeHash().ToArray())
	if err != nil {
		return fmt.Errorf("DexFund.Init error:%s", err)
	}
	err = DexProto.Init(ctx, ctx.OntClient.Admin, DexP2P.CodeHash().ToArray())
	if err != nil {
		return fmt.Errorf("DexProto.Init error:%s", err)
	}
	err = DexP2P.Init(ctx, ctx.OntClient.Admin, 10)
	if err != nil {
		return fmt.Errorf("DexP2P.Init error:%s", err)
	}
	isOntDexInit = true
	return nil
}