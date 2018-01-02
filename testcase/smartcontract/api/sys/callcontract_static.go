package sys

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestCallContractStatic(ctx *TestFrameworkContext) bool {
	codeA := "52c56b6c766b00527ac4616c766b00c36c766b51527ac46203006c766b51c3616c7566"
	codeHash, err := common.ToCodeHash([]byte(codeA))
	if err != nil {
		ctx.LogError("TestCallContractStatic ToCodeHash error:%s", err)
		return false
	}
	ctx.LogInfo("CodeHash: %x", codeHash)
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

	codeB := "52c56b6c766b00527ac4616c766b00c3616742bc967b05492676273682412439e5866e3a088d6c766b51527ac46203006c766b51c3616c7566"
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
    [Appcall("8d083a6e86e5392441823627762649057b96bc42")]
    public static extern int OtherContract(int input);
    public static int Main(int input)
    {
        return OtherContract(input);
    }
}

Code:52c56b6c766b00527ac4616c766b00c3616742bc967b05492676273682412439e5866e3a088d6c766b51527ac46203006c766b51c3616c7566
*/
