package deployinvoke

import (
	"github.com/ONT_TEST/testframework"
)

func TestInvokeSmartContract(ctx *testframework.TestFrameworkContext) bool {
	codeString := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C393616C7566"

	var a = 1
	var b = 5
	var c = a + b

	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		codeString,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("SimpleSmartContract InvokeSmartContract error:%s", err)
		return false
	}
	resi, ok := res.(float64)
	if !ok {
		ctx.LogError("SimpleSmartContract SmartContractRes:%v asset failed", res)
		return false
	}
	if int(resi) != c {
		ctx.LogError("SimpleSmartContract SmartContractRes:%v != %v", res, c)
		return false
	}
	return true
}
