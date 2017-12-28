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
        byte[] contractid = { 174, 75, 194, 130, 188, 134, 82, 102, 32, 155, 40, 243, 113, 252, 30, 177, 247, 188, 14, 116 };
        Contract ast = Blockchain.GetContract(contractid);
        return ast;
    }
}
 */

func TestGetContract(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b14ae4bc282bc865266209b28f371fc1eb1f7bc0e746c766b00527ac46c766b00c361681a4e656f2e426c6f636b636861696e2e476574436f6e74726163746c766b51527ac46c766b51c36c766b52527ac46203006c766b52c3616c7566"
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
		ctx.LogError("TestGetContract DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetContract WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetContract InvokeSmartContract error:%s", err)
		return false
	}
	hexstr, err := common.HexToBytes(res.(string))
	if err != nil {
		ctx.LogError("TestGetContract HexToBytes error:%s", err)
		return false
	}
	contract := new(states.ContractState)
	bf := bytes.NewBuffer(hexstr)
	if err := contract.Deserialize(bf); err != nil {
		ctx.LogError("TestGetContract HexToBytes error:%s", err)
		return false
	}
	ctx.LogError("TestGetContract :%+v ", contract)
	return true
}