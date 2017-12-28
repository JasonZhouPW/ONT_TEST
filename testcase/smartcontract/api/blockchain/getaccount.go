package blockchain

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/states"
	"bytes"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object Main()
    {
    	byte[] programHash = {222,22,168,155,127,237,137,116,234,99,88,103,178,63,254,214,234,83,239,81};
        Account account = Blockchain.GetAccount(programHash);
        return account;
    }
}
 */

func TestGetAccount(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b14de16a89b7fed8974ea635867b23ffed6ea53ef516c766b00527ac46c766b00c36168194e656f2e426c6f636b636861696e2e4765744163636f756e746c766b51527ac46c766b51c36c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.InteropInterface),
		"TestGetAccount",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetAccount DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetAccount WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetAccount InvokeSmartContract error:%s", err)
		return false
	}
	hexstr, err := common.HexToBytes(res.(string))
	if err != nil {
		ctx.LogError("TestGetAccount HexToBytes error:%s", err)
		return false
	}
	account := new(states.AccountState)
	bf := bytes.NewBuffer(hexstr)
	if err := account.Deserialize(bf); err != nil {
		ctx.LogError("TestGetAccount HexToBytes error:%s", err)
		return false
	}
	ctx.LogError("TestGetAccount :%+v ", account)
	return true
}