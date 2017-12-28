package deployinvoke

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : FunctionCode
{
    public static int Main(int a, int b)
    {
        return a + b;
    }
}
----------------------------------------
AVMHexString: 52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C393616C7566
*/


func TestDeploySmartContract(ctx *testframework.TestFrameworkContext) bool {
	codeString := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C393616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		codeString,
		[]contract.ContractParameterType{contract.Integer, contract.Integer},
		contract.ContractParameterType(contract.Integer),
		"SimpleSmartContract",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("SimpleSmartContract DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("SimpleSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}
	return true
}