package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestDexFund(ctx *TestFrameworkContext) bool {
	code := DExFundCode
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
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
		ctx.LogError("TestDexFund DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDexFund WaitForGenerateBlock error:%s", err)
		return false
	}
	if !testDexFundInit(ctx, code) {
		return false
	}
	if !testReceipt(ctx, code) {
		return false
	}
	if !testLock(ctx, code) {
		return false
	}
	if !testUnLock(ctx, code) {
		return false
	}
	if !testPayment(ctx, code) {
		return false
	}
	return true
}

func testDexFundInit(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"init", []interface{}{[]byte{1}, ctx.OntClient.Account1.ProgramHash.ToArray()}},
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
	if errorCode != 0 {
		ctx.LogError("testDexFundInit failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testReceipt(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"receipt", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testReceipt error:%s", err)
		return false
	}
	ctx.LogInfo("testReceipt res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testReceipt getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testReceipt failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testPayment(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"payment", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testPayment error:%s", err)
		return false
	}
	ctx.LogInfo("testPayment res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testPayment getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testPayment failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testLock(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"lock", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testLock error:%s", err)
		return false
	}
	ctx.LogInfo("testLock res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testLock getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testLock failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testUnLock(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"unlock", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testUnLock error:%s", err)
		return false
	}
	ctx.LogInfo("testUnLock res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testUnLock getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testUnLock failed errorCode:%d", errorCode)
		return false
	}
	return true
}
