package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func deployDexFund(ctx *TestFrameworkContext) bool {
	code := DExFundCode
	c, _ := common.HexToBytes(code)
	codeHash, err := common.ToCodeHash(c)
	if err != nil {
		ctx.LogError("TestDexFund ToCodeHash error:%s", err)
		return false
	}
	ctx.LogInfo("TestDexFund CodeHash: %x , RCodeHash: %x", codeHash, codeHash.ToArrayReverse())
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"TestDexFund",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("deployDexFund DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("deployDexFund WaitForGenerateBlock error:%s", err)
		return false
	}
	//assetId := []byte("")
	//admin := ctx.OntClient.Admin
	//if !testDexFundInit(ctx, assetId, admin) {
	//	return false
	//}
	//buyer := ctx.OntClient.Account1
	//amount := 10
	//if !testReceipt(ctx, buyer, amount) {
	//	return false
	//}
	//if !testLock(ctx, buyer, amount) {
	//	return false
	//}
	//if !testUnLock(ctx, buyer, amount) {
	//	return false
	//}
	//if !testPayment(ctx, buyer, amount) {
	//	return false
	//}
	return true
}

func registerAsset(ctx *TestFrameworkContext, asset *asset.Asset, amount common.Fixed64, issuer, controler *account.Account) bool {
	regTx, err := ctx.Ont.NewAssetRegisterTransaction(asset, amount, issuer, controler)
	if err != nil {
		ctx.LogError("registerAsset NewAssetRegisterTransaction Asset:%+v Amount:%v Issuer:%+v Controler:%+v error:%s",
			asset,
			amount,
			issuer,
			controler,
			err)
		ctx.FailNow()
		return false
	}
	txHash, err := ctx.Ont.SendTransaction(controler, regTx)
	if err != nil {
		ctx.LogError("registerAsset SendTransaction AssetRegisterTransaction error:%s", err)
		ctx.FailNow()
		return false
	}
	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		ctx.LogError("registerAsset WaitForGenerateBlock error:%s", err)
		ctx.FailNow()
		return false
	}
	ctx.OntAsset.RegAsset(txHash, asset)
	return true
}

func dexFundInit(ctx *TestFrameworkContext, assetId []byte, admin *account.Account) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DExFundCode,
		[]interface{}{"init", []interface{}{assetId, admin.ProgramHash.ToArray(), []byte("")}},
	)
	if err != nil {
		ctx.LogError("testDexFundInit error:%s", err)
		return false
	}
	ctx.LogInfo("testDexFundInit res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testDexFundInit getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 && errorCode != 1010 {
		ctx.LogError("testDexFundInit failed errorCode:%d", errorCode)
		return false
	}
	return true
}

//
//func testDeposit(ctx *TestFrameworkContext, )bool{
//
//}
//
func fundReceipt(ctx *TestFrameworkContext, receiver *account.Account, amount int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		receiver,
		DExFundCode,
		[]interface{}{"receipt", []interface{}{ receiver.ProgramHash.ToArray(), amount}},
	)
	if err != nil {
		ctx.LogError("fundReceipt error:%s", err)
		return false
	}
	ctx.LogInfo("fundReceipt res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("fundReceipt getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("fundReceipt failed errorCode:%d", errorCode)
		return false
	}
	return true
}
//
//func testPayment(ctx *TestFrameworkContext,  buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"payment", []interface{}{ buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testPayment error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testPayment res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testPayment getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testPayment failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testLock(ctx *TestFrameworkContext, buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"lock", []interface{}{buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testLock error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testLock res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testLock getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testLock failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testUnLock(ctx *TestFrameworkContext, buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"unlock", []interface{}{buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testUnLock error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testUnLock res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testUnLock getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testUnLock failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}

func setFundCaller(ctx *TestFrameworkContext, admin *account.Account, caller []byte)bool{
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DExFundCode,
		[]interface{}{"setcaller", []interface{}{caller}},
	)
	if err != nil {
		ctx.LogError("setFundCaller error:%s", err)
		return false
	}
	ctx.LogInfo("setFundCaller res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("setFundCaller getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("setFundCaller failed errorCode:%d", errorCode)
		return false
	}
	return true
}