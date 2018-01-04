package smartcontract

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestCallingScriptHash(ctx *TestFrameworkContext)bool{
	code := "51c56b6161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b00527ac46203006c766b00c3616c7566"
	c, _ := common.HexToBytes(code)
	codeHash,_ := common.ToCodeHash(c)
	ctx.LogInfo("CodeA: %x R:%x", codeHash, codeHash.ToArrayReverse())

	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestCallingScriptHashA",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestCallingScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCallingScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}
	code = "51c56b616167aed4b6ec4c8987ff7f79af45af88a3139ed10c7d6c766b00527ac46203006c766b00c3616c7566"
	c, _ = common.HexToBytes(code)
	codeHash,_ = common.ToCodeHash(c)
	ctx.LogInfo("CodeB: %x R:%x", codeHash,codeHash.ToArrayReverse())

	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestCallingScriptHashB",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestCallingScriptHash DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCallingScriptHash WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestCallingScriptHash error:%s", err)
		return false
	}

	ctx.LogInfo("TestCallingScriptHash res:%s", res)
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
        return ExecutionEngine.CallingScriptHash;
    }
}
Code := 51c56b6161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b00527ac46203006c766b00c3616c7566
---------------------------------------------------------------

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class B : SmartContract
{
    [Appcall("70a636aa57ab38cf9da5d47d0b6a563922904c62")]
    public static extern byte[] OtherContract();
    public static byte[] Main()
    {
        return OtherContract();
    }
}

Code := 51c56b616167aed4b6ec4c8987ff7f79af45af88a3139ed10c7d6c766b00527ac46203006c766b00c3616c7566
*/
