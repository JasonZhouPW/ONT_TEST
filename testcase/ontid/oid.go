package ontid

import (
	"io/ioutil"
	"time"
	"fmt"
	"github.com/Ontology/common"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
)

func TestONTID(ctx *testframework.TestFrameworkContext) bool {
	buf, err := ioutil.ReadFile("./contract.txt")
	if err != nil {
		return false
	}
	code := string(buf)
	code = common.ToHexString(buf)
	fmt.Println(code)
	_, err = ctx.Ont.DeploySmartContract(
		ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Array},
		contract.ContractParameterType(contract.Integer),
		"TestONTID",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)

	if err != nil {
		ctx.LogError("TestONTID DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestONTID WaitForGenerateBlock error:%s", err)
		return false
	}

	// call the contract
	params := []interface{}{"CreateIdentityByPublicKey", []interface{}{[]byte{0x02, 0x01, 0x02}, 
		[]byte{3, 37, 205, 73, 48, 136, 80, 183,
			242, 241, 5, 116, 157, 189, 148, 67,
			197, 158, 206, 42, 35, 68, 2, 199, 34,
			239, 52, 251, 241, 179, 38, 68, 167 }}}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{params},
	)
	if err != nil {
		ctx.LogError("TestONTID InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, len(params))
	if err != nil {
		ctx.LogError("TestONTID test failed %s", err)
		return false
	}

	return true
}
