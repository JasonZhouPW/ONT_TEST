package smartcontract

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"encoding/hex"
)

func TestExecutingScriptHash(ctx *TestFrameworkContext) bool {
	code := "52c56b6151c56c766b00527ac46c766b00c30061682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368c46c766b00c36c766b51527ac46203006c766b51c3616c7566"
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

	v, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestExecutingScriptHash assert failed")
		return false
	}

	hc, ok := v[0].(string)
	if !ok {
		ctx.LogError("TestExecutingScriptHash assert to string failed")
		return false
	}

	c, _ := common.HexToBytes(code)
	codeHash, _ := common.ToCodeHash(c)
	
	hcb, _ := hex.DecodeString(hc)

	if string(codeHash.ToArray()) != string(hcb){
		ctx.LogError("TestExecutingScriptHash res: %x != %x", hcb, codeHash.ToArray())
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
    public static object[] Main()
    {
        object[] ret = new object[1];
        ret[0] = ExecutionEngine.ExecutingScriptHash;
        return ret;
    }
}
*/
