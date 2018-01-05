package smartcontract

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestExecutingScriptHash(ctx *TestFrameworkContext) bool {
	code := "51c56b6161682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b00527ac46203006c766b00c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestExecutingScriptHash",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestExecutingScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestExecutingScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestExecutingScriptHash error:%s", err)
		return false
	}
	ctx.LogInfo("TestExecutingScriptHash res:%s", res)

	c, _ := common.HexToBytes(code)
	codeHash, _ := common.ToCodeHash(c)

	err = ctx.AssertToByteArray(res, codeHash.ToArray())
	if err != nil {
		ctx.LogError("AssertToByteArray error:%s", err)
		return false
	}

	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class A : SmartContract
{
    public static byte[] Main()
    {
        return ExecutionEngine.ExecutingScriptHash;
    }
}
*/
