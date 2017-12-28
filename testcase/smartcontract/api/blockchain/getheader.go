package blockchain

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/ledger"
	"github.com/Ontology/common"
	"bytes"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object Main()
    {
        uint height = Blockchain.GetHeight();
        Header head = Blockchain.GetHeader(height);
        return head;
    }
}
 */

func TestGetHeader(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6168184E656F2E426C6F636B636861696E2E4765744865696768746C766B00527AC46C766B00C36168184E656F2E426C6F636B636861696E2E4765744865616465726C766B51527AC46C766B51C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.InteropInterface),
		"TestGetHeader",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetHeader DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetHeader WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetHeader InvokeSmartContract error:%s", err)
		return false
	}
	hexstr, err := common.HexToBytes(res.(string))
	if err != nil {
		ctx.LogError("TestGetHeader HexToBytes error:%s", err)
		return false
	}
	header := new(ledger.Header)
	bf := bytes.NewBuffer(hexstr)
	if err := header.Deserialize(bf); err != nil {
		ctx.LogError("TestGetHeader HexToBytes error:%s", err)
		return false
	}
	ctx.LogError("TestGetHeader :%+v ", header.Blockdata)
	return true
}