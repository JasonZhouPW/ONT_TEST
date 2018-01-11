package executionengine

import (
. "github.com/ONT_TEST/testframework"
"github.com/Ontology/common"
"github.com/Ontology/core/contract"
"github.com/Ontology/smartcontract/types"
"time"
)

func TestEntryScriptHash(ctx *TestFrameworkContext)bool{
	code := "51c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574456e747279536372697074486173686c766b00527ac46203006c766b00c3616c7566"
	c, _ := common.HexToBytes(code)
	codeHash,_ := common.ToCodeHash(c)
	ctx.LogInfo("CodeA: %x R:%x", codeHash, codeHash.ToArrayReverse())

	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestEntryScriptHashA",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestEntryScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestEntryScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}
	code = "51c56b6161671f6a437bed510b4a7d698d98ca192faa2ae634b86c766b00527ac46203006c766b00c3616c7566"
	c, _ = common.HexToBytes(code)
	codeHash,_ = common.ToCodeHash(c)
	ctx.LogInfo("CodeB: %x R:%x", codeHash,codeHash.ToArrayReverse())

	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestEntryScriptHashB",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestEntryScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestEntryScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestEntryScriptHash error:%s", err)
		return false
	}

	ctx.LogInfo("TestEntryScriptHash res:%s", res)
	err = ctx.AssertToByteArray(res, codeHash.ToArray())
	if err != nil {
		ctx.LogError("TestEntryScriptHash AssertToByteArray error:%s", err)
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
        return ExecutionEngine.EntryScriptHash;
    }
}
Code := 51c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574456e747279536372697074486173686c766b00527ac46203006c766b00c3616c7566
---------------------------------------------------------------

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class B : SmartContract
{
    [Appcall("b834e62aaa2f19ca988d697d4a0b51ed7b436a1f")]
    public static extern byte[] OtherContract();
    public static byte[] Main()
    {
        return OtherContract();
    }
}

Code := 51c56b6161671f6a437bed510b4a7d698d98ca192faa2ae634b86c766b00527ac46203006c766b00c3616c7566
*/
