package smartcontract

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"github.com/ONT_TEST/testcase/ont_dex"
	"reflect"
	"bytes"
	"encoding/hex"
)

func TestExecutionEngine(ctx *TestFrameworkContext)bool{
	code := "52c56b6153c56c766b00527ac46c766b00c35161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e6753637269707448617368c46c766b00c35261682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368c46c766b00c36c766b51527ac46203006c766b51c3616c7566"
	c, _ := common.HexToBytes(code)
	codeHashA,_ := common.ToCodeHash(c)
	ctx.LogInfo("CodeA: %x R:%x", codeHashA, codeHashA.ToArrayReverse())

	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Array),
		"TestExecutionEngineA",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestExecutionEngine DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestExecutionEngine WaitForGenerateBlock error:%s", err)
		return false
	}
	code = "51c56b616167271a3444931a2081fe0544cf9309a9542b9c67fe6c766b00527ac46203006c766b00c3616c7566"
	c, _ = common.HexToBytes(code)
	codeHashB,_ := common.ToCodeHash(c)
	ctx.LogInfo("CodeB: %x R:%x", codeHashB,codeHashB.ToArrayReverse())

	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Array),
		"TestExecutionEngine",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestExecutionEngine DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestExecutionEngine WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestExecutionEngine error:%s", err)
		return false
	}

	ctx.LogInfo("TestExecutionEngine res:%s", res)

	callingHash, err := ont_dex.GetRetValue(res, 1, reflect.String)
	if err != nil {
		ctx.LogInfo("TestExecutionEngine GetRetValue error:%s", err)
		return false
	}

	data, _ := hex.DecodeString(callingHash.(string))
	if !bytes.EqualFold(data, codeHashB.ToArray()){
		ctx.LogError("TestExecutionEngine callingHash:%s != %x", callingHash, codeHashB)
		return false
	}

	executeHash, err := ont_dex.GetRetValue(res, 2, reflect.String)
	if err != nil {
		ctx.LogInfo("TestExecutionEngine GetRetValue error:%s", err)
		return false
	}

	data, _ = hex.DecodeString(executeHash.(string))
	if !bytes.EqualFold(data, codeHashA.ToArray()){
		ctx.LogError("TestExecutionEngine executinghash:%s != %x", executeHash, codeHashA)
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
		object[] ret = new object[3];
		//ret[0] = ExecutionEngine.EntryScriptHash;
		ret[1] = ExecutionEngine.CallingScriptHash;
		ret[2] = ExecutionEngine.ExecutingScriptHash;
        return ret;
    }
}
Code := 52c56b6153c56c766b00527ac46c766b00c30061682953797374656d2e457865637574696f6e456e67696e652e476574456e74727953637269707448617368c46c766b00c35161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e6753637269707448617368c46c766b00c35261682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368c46c766b00c36c766b51527ac46203006c766b51c3616c7566
---------------------------------------------------------------

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class B : SmartContract
{
    [Appcall("fe679c2b54a90993cf4405fe81201a9344341a27")]
    public static extern object[] OtherContract();
    public static object[] Main()
    {
        return OtherContract();
    }
}

Code := 51c56b616167271a3444931a2081fe0544cf9309a9542b9c67fe6c766b00527ac46203006c766b00c3616c7566
*/
