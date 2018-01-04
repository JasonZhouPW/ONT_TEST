package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func deployDexP2P(ctx *TestFrameworkContext) bool {
	code := DEXP2PCode
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"TestDExP2P",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("deployDexP2P DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("deployDexP2P WaitForGenerateBlock error:%s", err)
		return false
	}

	//if !testDexP2PInit(ctx){
	//	return false
	//}
	//buyer := ctx.OntClient.Account1
	//seller := ctx.OntClient.Account2
	//orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	//orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	//if err != nil {
	//	ctx.LogError("TestDexP2P crypto.Sign error:%s", err)
	//	return false
	//}
	//amount := 10
	//if !testReceipt(ctx, buyer, amount) {
	//	return false
	//}
	//if !testMakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
	//	return false
	//}
	//
	//if !testBuyOrderComplete(ctx, orderId, buyer){
	//	return false
	//}
	//
	//orderId = []byte(fmt.Sprint("%d", rand.Int31()))
	//orderSig, err = crypto.Sign(buyer.PrivateKey, orderId)
	//if err != nil {
	//	ctx.LogError("TestDexP2P crypto.Sign error:%s", err)
	//	return false
	//}
	//amount = 11
	//if !testReceipt(ctx, buyer, amount) {
	//	return false
	//}
	//if !testMakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
	//	return false
	//}
	//
	//if !testBuyOrderCancel(ctx, orderId, buyer){
	//	return false
	//}
	return true
}

func dexP2PInit(ctx *TestFrameworkContext)bool{
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		DEXP2PCode,
		[]interface{}{"init", []interface{}{}},
	)
	if err != nil {
		ctx.LogError("dexP2PInit error:%s", err)
		return false
	}
	if err != nil {
		ctx.LogError("dexP2PInit error:%s", err)
		return false
	}
	ctx.LogInfo("dexP2PInit res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("dexP2PInit getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 && errorCode != 3009 {
		ctx.LogError("dexP2PInit failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func makeBuyOrder(ctx *TestFrameworkContext, orderSig, orderId []byte, buyer, seller *account.Account, amount int) bool {
	buyerPk, err := buyer.PublicKey.EncodePoint(true)
	if err != nil {
		ctx.LogError("makeBuyOrder PublicKey.EncodePoint error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		seller,
		DEXP2PCode,
		[]interface{}{"makebuyorder", []interface{}{orderSig, orderId, buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), buyerPk, amount}},
	)
	if err != nil {
		ctx.LogError("makeBuyOrder error:%s", err)
		return false
	}
	ctx.LogInfo("makeBuyOrder res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("makeBuyOrder getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("makeBuyOrder failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func buyOrderComplete(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account)bool{
	res, err := ctx.Ont.InvokeSmartContract(
		buyer,
		DEXP2PCode,
		[]interface{}{"buyordercomplete", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
	)
	if err != nil {
		ctx.LogError("buyOrderComplete error:%s", err)
		return false
	}
	ctx.LogInfo("buyOrderComplete res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("buyOrderComplete getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("buyOrderComplete failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func buyOrderCancel(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account)bool{
	res, err := ctx.Ont.InvokeSmartContract(
		buyer,
		DEXP2PCode,
		[]interface{}{"buyordercancel", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
	)
	if err != nil {
		ctx.LogError("buyOrderComplete error:%s", err)
		return false
	}
	ctx.LogInfo("buyOrderComplete res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("buyOrderComplete getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("buyOrderComplete failed errorCode:%d", errorCode)
		return false
	}
	return true
}

