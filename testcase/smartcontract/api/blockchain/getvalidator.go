package blockchain

import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
	"fmt"
	"reflect"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object Main()
    {
        byte[][] ast = Blockchain.GetValidators();
        return ast;
    }
}
 */

func TestValidators(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b6161681c4e656f2e426c6f636b636861696e2e47657456616c696461746f72736c766b00527ac46c766b00c36c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.Array),
		"TestValidators",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestValidators DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestValidators WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestValidators InvokeSmartContract error:%s", err)
		return false
	}
	fmt.Println("res", reflect.TypeOf(res))
	//hexstr, err := common.HexToBytes(res.([]interface {}))
	//if err != nil {
	//	ctx.LogError("TestValidators HexToBytes error:%s", err)
	//	return false
	//}
	//state := new(states.ValidatorState)
	//bf := bytes.NewBuffer(hexstr)
	//if err := state.Deserialize(bf); err != nil {
	//	ctx.LogError("TestValidators HexToBytes error:%s", err)
	//	return false
	//}
	ctx.LogError("TestValidators :%+v ", res)
	return true
}