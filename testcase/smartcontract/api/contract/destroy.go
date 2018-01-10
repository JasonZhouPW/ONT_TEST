package contract

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/*
 using Neo.SmartContract.Framework;
 using Neo.SmartContract.Framework.Services.Neo;
 using Neo.SmartContract.Framework.Services.System;

public class Contract1:SmartContract
{
    public static void Main()
    {
        Neo.SmartContract.Framework.Services.Neo.Contract.Destroy();
    }
}

*/

func TestContractDestroy(ctx *testframework.TestFrameworkContext) bool {
	code := "00c56b616168144e656f2e436f6e74726163742e44657374726f7961616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Void},
		contract.ContractParameterType(contract.Void),
		"TestContractDestroy",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)

	if err != nil {
		ctx.LogError("TestContractDestroy DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 2)
	if err != nil {
		ctx.LogError("TestContractDestroy WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		nil,
	)
	if err != nil {
		ctx.LogError("TestContractDestroy InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestContractDestroy1 :%+v ", res)

	res, err = ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		nil,
	)

	if err == nil {
		ctx.LogError("TestContractDestroy InvokeSmartContract error should fail, return:%v", res)
		return false
	}

	return true
}
