package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func deployDexProto(ctx *TestFrameworkContext) bool {
	code := DExProtoCode
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"TestDexProto",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("deployDexProto DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("deployDexProto WaitForGenerateBlock error:%s", err)
		return false
	}
	//admin := ctx.OntClient.Admin
	//if !testDexProtoInit(ctx, admin) {
	//	return false
	//}
	//buyer := ctx.OntClient.Account1
	//seller := ctx.OntClient.Account2
	//amount := 11
	//if !testReceipt(ctx, buyer, amount) {
	//	return false
	//}
	//if !testOnMakeOrder(ctx, buyer, seller, amount) {
	//	return false
	//}
	//if !testOnOrderComplete(ctx, buyer, seller, amount) {
	//	return false
	//}
	//buyer, seller = seller, buyer
	//amount = 12
	//if !testReceipt(ctx, buyer, amount) {
	//	return false
	//}
	//if !testOnMakeOrder(ctx, buyer, seller, amount) {
	//	return false
	//}
	//if !testOnOrderCancel(ctx, buyer, seller, amount) {
	//	return false
	//}
	return true
}

func dexProtoInit(ctx *TestFrameworkContext, admin *account.Account) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DExProtoCode,
		[]interface{}{"init", []interface{}{admin.ProgramHash.ToArray(), []byte("")}},
	)
	if err != nil {
		ctx.LogError("dexProtoInit error:%s", err)
		return false
	}
	ctx.LogInfo("dexProtoInit res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("dexProtoInit getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 && errorCode != 2009 {
		ctx.LogError("dexProtoInit failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func addProtoCaller(ctx *TestFrameworkContext,admin *account.Account, caller []byte) bool{
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DExProtoCode,
		[]interface{}{"addcaller", []interface{}{caller}},
	)
	if err != nil {
		ctx.LogError("addProtoCaller error:%s", err)
		return false
	}
	ctx.LogInfo("addProtoCaller res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("addProtoCaller getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("addProtoCaller failed errorCode:%d", errorCode)
		return false
	}
	return true
}
//
//func testOnMakeOrder(ctx *TestFrameworkContext,buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onmakeorder", []interface{}{ buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnMakeOrder error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnMakeOrder res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnMakeOrder getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnMakeOrder failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testOnOrderComplete(ctx *TestFrameworkContext, buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onordercomplete", []interface{}{ buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnOrderComplete error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnOrderComplete res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnOrderComplete getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnOrderComplete failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testOnOrderCancel(ctx *TestFrameworkContext, buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onordercancel", []interface{}{ buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnOrderCancel error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnOrderCancel res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnOrderCancel getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnOrderCancel failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
