package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestDexProto(ctx *TestFrameworkContext)bool{
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
		ctx.LogError("TestDexProto DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDexProto WaitForGenerateBlock error:%s", err)
		return false
	}
	//if !testDexProtoInit(ctx, code) {
	//	return false
	//}
	if !testOnMakeOrder(ctx, code) {
		return false
	}
	if !testOnOrderComplete(ctx, code) {
		return false
	}
	if !testOnOrderCancel(ctx, code) {
		return false
	}
	return true;
}

func testDexProtoInit(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"init", []interface{}{[]byte{1}, ctx.OntClient.Account1.ProgramHash.ToArray()}},
	)
	if err != nil {
		ctx.LogError("testDexProtoInit error:%s", err)
		return false
	}
	ctx.LogInfo("testDexProtoInit res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testDexProtoInit getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testDexProtoInit failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testOnMakeOrder(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"onmakeorder", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), ctx.OntClient.Account2.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testOnMakeOrder error:%s", err)
		return false
	}
	ctx.LogInfo("testOnMakeOrder res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testOnMakeOrder getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testOnMakeOrder failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testOnOrderComplete(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"onordercomplete", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), ctx.OntClient.Account2.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testOnOrderComplete error:%s", err)
		return false
	}
	ctx.LogInfo("testOnOrderComplete res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testOnOrderComplete getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testOnOrderComplete failed errorCode:%d", errorCode)
		return false
	}
	return true
}

func testOnOrderCancel(ctx *TestFrameworkContext, code string) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{"onordercancel", []interface{}{ctx.OntClient.Account1.ProgramHash.ToArray(), ctx.OntClient.Account2.ProgramHash.ToArray(), 10}},
	)
	if err != nil {
		ctx.LogError("testOnOrderCancel error:%s", err)
		return false
	}
	ctx.LogInfo("testOnOrderCancel res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		ctx.LogError("testOnOrderCancel getErrorCode error:%s", err)
		return false
	}
	if errorCode != 0 {
		ctx.LogError("testOnOrderCancel failed errorCode:%d", errorCode)
		return false
	}
	return true
}

