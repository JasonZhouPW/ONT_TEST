package runtime


import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestCheckWitness(ctx *TestFrameworkContext) bool {
	code := "52c56b6c766b00527ac4616c766b00c36168184e656f2e52756e74696d652e436865636b5769746e6573736c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Boolean),
		"TestCheckWitness",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestCheckWitness DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestCheckWitness WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testCheckWitness(ctx, code, ctx.OntClient.Account1.ProgramHash.ToArray(), true){
		return false
	}

	if !testCheckWitness(ctx, code, ctx.OntClient.Account2.ProgramHash.ToArray(), false){
		return false
	}

	return true
}

func testCheckWitness(ctx *TestFrameworkContext,code string, input []byte, expect bool) bool{
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestCheckWitness error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(res, expect)
	if err != nil {
		ctx.LogError("TestCheckWitness AssertToBoolean error:%s", err)
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
    public static bool Main(byte[] input)
    {
        return Runtime.CheckWitness(input);
    }
}
*/