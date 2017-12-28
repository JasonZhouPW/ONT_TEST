package utils

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"math/big"
	vmtype "github.com/Ontology/vm/neovm/types"
)

func TestAsBigInteger(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b6c766b00527ac4616c766b00c36c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.Integer),
		"TestAsBigInteger",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestAsBigInteger DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestAsBigInteger WaitForGenerateBlock error:%s", err)
		return false
	}

	b := big.NewInt(1233)
	if !testAsBigInteger(ctx, code, b) {
		return false
	}
	b = big.NewInt(0)
	if !testAsBigInteger(ctx, code, b){
		return false
	}
	b = big.NewInt(-1233)
	if !testAsBigInteger(ctx, code, b) {
		return false
	}
	return true
}

func testAsBigInteger(ctx *testframework.TestFrameworkContext, code string, b *big.Int) bool {
	data := vmtype.ConvertBigIntegerToBytes(b)
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{data},
	)
	if err != nil {
		ctx.LogError("TestAsBigInteger InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertBigInteger(res, b)
	if err != nil {
		ctx.LogError("TestAsBigInteger test failed %s", err)
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
    public static BigInteger Main(byte[] v)
    {
        return v.AsBigInteger();
    }
}
*/
