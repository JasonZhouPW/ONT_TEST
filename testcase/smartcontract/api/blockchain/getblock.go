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
            uint height = Blockchain.GetHeight();        //ok
            Header head = Blockchain.GetHeader(height);   //ok
            byte[] headHash = head.Hash;                  //
            Block block = Blockchain.GetBlock(headHash);  // ok
            return block;
    }
}
 */

func TestGetBlock(ctx *testframework.TestFrameworkContext) bool {
	code := "54C56B6168184E656F2E426C6F636B636861696E2E4765744865696768746C766B00527AC46C766B00C36168184E656F2E426C6F636B636861696E2E4765744865616465726C766B51527AC46C766B51C36168124E656F2E4865616465722E476574486173686C766B52527AC46C766B52C36168174E656F2E426C6F636B636861696E2E476574426C6F636B6C766B53527AC46C766B53C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.InteropInterface),
		"TestGetBlock",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBlock DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBlock WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetBlock InvokeSmartContract error:%s", err)
		return false
	}
	hexstr, err := common.HexToBytes(res.(string))
	if err != nil {
		ctx.LogError("TestGetBlock HexToBytes error:%s", err)
		return false
	}
	block := new(ledger.Block)
	bf := bytes.NewBuffer(hexstr)
	if err := block.Deserialize(bf); err != nil {
		ctx.LogError("TestGetBlock HexToBytes error:%s", err)
		return false
	}
	ctx.LogError("TestGetBlock :%+v ", block)
	return true
}
