package blockchain

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"encoding/hex"
	"bytes"
)

/**

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static string Main()
    {
        return "Hello World!";
    }
}

code = 51c56b610c48656c6c6f20576f726c64216c766b00527ac46203006c766b00c3616c7566
------------------------------------------------------------

using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;
class B : SmartContract
{
    public static byte[] Main(byte[] codeHash)
    {
        return Blockchain.GetContract(codeHash).Script;
    }
}
*/

func TestGetContract(ctx *testframework.TestFrameworkContext) bool {
	code := "52c56b6c766b00527ac4616c766b00c361681a4e656f2e426c6f636b636861696e2e476574436f6e74726163746168164e656f2e436f6e74726163742e4765745363726970746c766b51527ac46203006c766b51c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetContract",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetContract DeploySmartContract code error:%s", err)
		return false
	}

	codeA := "51c56b610c48656c6c6f20576f726c64216c766b00527ac46203006c766b00c3616c7566"
	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		codeA,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.ByteArray),
		"TestGetContract",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetContract DeploySmartContract codeA error:%s", err)
		return false
	}

	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetContract WaitForGenerateBlock error:%s", err)
		return false
	}

	c, _ := common.HexToBytes(codeA)
	codeHash, _ := common.ToCodeHash(c)
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{codeHash.ToArray()},
	)
	if err != nil {
		ctx.LogError("TestGetContract InvokeSmartContract error:%s", err)
		return false
	}

	script, ok := res.(string)
	if !ok{
		ctx.LogError("TestGetContract asset res to string failed")
		return false
	}

	data, _ := hex.DecodeString(script)
	codeAHash, _ := common.ToCodeHash(data)
	if !bytes.EqualFold(codeAHash.ToArray(), codeHash.ToArray()){
		ctx.LogError("TestGetContract %x != %x", data, codeHash.ToArray())
		return false
	}
	return true
}
