package blockchain

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"bytes"
	"github.com/Ontology/core/transaction"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
        public static object Main()
        {
            uint height0 = Blockchain.GetHeight();
            Header head = Blockchain.GetHeader(height0);   //ok 3162


            Block block0 = Blockchain.GetBlock(head.Hash);
            int txCount = block0.GetTransactionCount();

            Transaction tx0 = block0.GetTransaction(0);
            byte[] txhash = tx0.Hash;
            Transaction getTx = Blockchain.GetTransaction(txhash);

            return getTx;
        }
}
 */

func TestGetTransaction(ctx *testframework.TestFrameworkContext) bool {
	code := "56C56B6168184E656F2E426C6F636B636861696E2E4765744865696768746C766B00527AC46C766B00C36168184E656F2E426C6F636B636861696E2E4765744865616465726C766B51527AC46C766B51C36168124E656F2E4865616465722E476574486173686168174E656F2E426C6F636B636861696E2E476574426C6F636B6C766B52527AC46C766B52C361681D4E656F2E426C6F636B2E4765745472616E73616374696F6E436F756E74756C766B52C300617C68184E656F2E426C6F636B2E4765745472616E73616374696F6E6C766B53527AC46C766B53C36168174E656F2E5472616E73616374696F6E2E476574486173686C766B54527AC46C766B54C361681D4E656F2E426C6F636B636861696E2E4765745472616E73616374696F6E6C766B55527AC46C766B55C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.InteropInterface),
		"TestGetTransaction",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetTransaction DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetTransaction WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetTransaction InvokeSmartContract error:%s", err)
		return false
	}
	hexstr, err := common.HexToBytes(res.(string))
	if err != nil {
		ctx.LogError("TestGetTransaction HexToBytes error:%s", err)
		return false
	}
	trans := new(transaction.Transaction)
	bf := bytes.NewBuffer(hexstr)
	if err := trans.Deserialize(bf); err != nil {
		ctx.LogError("TestGetTransaction HexToBytes error:%s", err)
		return false
	}
	ctx.LogError("TestGetTransaction :%+v ", trans)
	return true
}