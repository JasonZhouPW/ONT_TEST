package sys

import (
	. "github.com/ONT_TEST/testframework"
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
)

func TestTriggerType(ctx *TestFrameworkContext)bool{
	code := "55c56b616168164e656f2e52756e74696d652e47657454726967676572609c6c766b00527ac46c766b00c3640f0061516c766b51527ac462b2006168164e656f2e52756e74696d652e47657454726967676572009c6c766b52527ac46c766b52c3640f0061526c766b51527ac4627c006168164e656f2e52756e74696d652e4765745472696767657201119c6c766b53527ac46c766b53c3640f0061536c766b51527ac46245006168164e656f2e52756e74696d652e47657454726967676572519c6c766b54527ac46c766b54c3640f0061546c766b51527ac4620f0061556c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{ },
		contract.ContractParameterType(contract.Integer),
		"TestTriggerType",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestTriggerType DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestTriggerType WaitForGenerateBlock error:%s", err)
		return false
	}

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestTriggerType error:%s", err)
		return false
	}
	ctx.LogInfo("TestTriggerType type:%d", int(res.(float64)))
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
using System.Numerics;

public class HelloWorld : SmartContract
{
    public static int Main()
    {
        if (Runtime.Trigger == TriggerType.Application)
        {
            return 1;
        }
        else if (Runtime.Trigger == TriggerType.Verification)
        {
            return 2;
        }
        else if (Runtime.Trigger == TriggerType.ApplicationR)
        {
            return 3;
        }
        else if (Runtime.Trigger == TriggerType.VerificationR)
        {
            return 4;
        }
        else {
            return 5;
        }
    }
}
*/