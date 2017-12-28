package blockchain

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static uint Main()
    {
        uint height = Blockchain.GetHeight();
        return height;
    }
}
 */

func TestGetCurrentHeight(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6168184E656F2E426C6F636B636861696E2E4765744865696768746C766B00527AC46C766B00C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestGetCurrentHeight",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetCurrentHeight InvokeSmartContract error:%s", err)
		return false
	}
	resp, ok := res.(float64)
	if !ok {
		ctx.LogError("TestGetCurrentHeight result:%v assert failed.", res)
		return false
	}
	ctx.LogError("TestGetCurrentHeight current height:%v ", resp)
	return true
}