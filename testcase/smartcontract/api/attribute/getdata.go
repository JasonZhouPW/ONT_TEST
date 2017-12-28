package attribute

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
    public static byte[] Main()
    {
        Block block = Blockchain.GetBlock(0);
        Transaction trans = block.GetTransaction(0);
        TransactionAttribute[] attributs =  trans.GetAttributes();
        TransactionAttribute attribute = attributs[0];
        return attribute.Data;
    }
}
 */

func TestGetAttributeData(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b61006168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b00527ac46c766b00c3006168184e656f2e426c6f636b2e4765745472616e73616374696f6e6c766b51527ac46c766b51c361681d4e656f2e5472616e73616374696f6e2e476574417474726962757465736c766b52527ac46c766b52c300c36c766b53527ac46c766b53c36168154e656f2e4174747269627574652e476574446174616c766b54527ac46203006c766b54c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetAttributeData",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetAttributeData DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetAttributeData WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestGetAttributeData InvokeSmartContract error:%s", err)
		return false
	}
	ctx.LogError("TestGetAttributeData :%+v ", res)
	return true
}