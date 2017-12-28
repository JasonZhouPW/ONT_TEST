package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	vmtypes "github.com/Ontology/vm/neovm/types"
	"math/big"
)

func TestAsByteArrayBigInteger(ctx *testframework.TestFrameworkContext) bool {
	code := "51C56B6C766B00527AC46C766B00C3616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.ByteArray),
		"TestAsByteArrayBigInteger",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestAsByteArrayBigInteger DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestAsByteArrayBigInteger WaitForGenerateBlock error:%s", err)
		return false
	}

	input := -233545554
	if !testAsArray_BigInteger(ctx, code, input){
		return false
	}
	input = -3434
	if !testAsArray_BigInteger(ctx, code, input){
		return false
	}
	input = 1
	if !testAsArray_BigInteger(ctx, code, input){
		return false
	}
	return true
}

func testAsArray_BigInteger(ctx *testframework.TestFrameworkContext, code string, input int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{input},
	)
	if err != nil {
		ctx.LogError("TestAsByteArrayBigInteger InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToByteArray(res, vmtypes.ConvertBigIntegerToBytes(big.NewInt(int64(input))))

	if err != nil {
		ctx.LogError("TestAsByteArrayBigInteger test failed %s", err)
		return false
	}
	return true
}

/*
using System.Numerics;
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static byte[] Main(BigInteger arg)
    {
        return arg.AsByteArray();
    }
}
*/
