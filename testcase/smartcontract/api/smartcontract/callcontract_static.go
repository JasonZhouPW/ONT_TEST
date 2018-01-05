package smartcontract

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"encoding/hex"
)

func TestCallContractStatic(ctx *TestFrameworkContext) bool {
	codeA := "52c56b6c766b00527ac4616c766b00c36c766b51527ac46203006c766b51c3616c7566"
	c, _ := hex.DecodeString(codeA)
	codeHash, err := common.ToCodeHash(c)
	if err != nil {
		ctx.LogError("TestCallContractStatic ToCodeHash error:%s", err)
		return false
	}
	ctx.LogInfo("CodeHash: %x R: %x", codeHash, codeHash.ToArrayReverse())
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		codeA,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestCallContractStaticA",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestCallContractStatic DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCallContractStatic WaitForGenerateBlock error:%s", err)
		return false
	}

	codeB := "52c56b6c766b00527ac4616c766b00c361673d711163a4da8a8e37fd469a37e6cc04d37df3696c766b51527ac46203006c766b51c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		codeB,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Integer),
		"TestCallContractStaticB",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestCallContractStatic DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCallContractStatic WaitForGenerateBlock error:%s", err)
		return false
	}

	input := 12
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		codeB,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestTriggerType error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, input)
	if err != nil {
		ctx.LogError("TestCallContractStatic res AssertToInt error:%s", err)
		return false;
	}
	return true
}

/*
SmartContract1

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static int Main()
    {
        return 0;
    }
}

Code:52c56b6c766b00527ac4616c766b00c36c766b51527ac46203006c766b51c3616c7566

SmartContract2
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System.Numerics;

public class HelloWorld : SmartContract
{
    [Appcall("69f37dd304cce6379a46fd378e8adaa46311713d")]
    public static extern int OtherContract(int input);
    public static int Main(int input)
    {
        return OtherContract(input);
    }
}

Code:52c56b6c766b00527ac4616c766b00c361673d711163a4da8a8e37fd469a37e6cc04d37df3696c766b51527ac46203006c766b51c3616c7566


using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using Neo.SmartContract.Framework.Services.System;
using System.Numerics;

public class B : SmartContract
{
    [Appcall("69f37dd304cce6379a46fd378e8adaa46311713d")]
    public static extern byte[] OtherContract();
    public static byte[] Main()
    {
        return OtherContract();
    }

}

Code:
*/
