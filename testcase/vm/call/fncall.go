package call

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestFnCall(ctx *testframework.TestFrameworkContext) bool {
	{
		code := contract_add
		_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
			code,
			[]contract.ContractParameterType{contract.Integer},
			contract.ContractParameterType(contract.Integer),
			"TestFnCall_contract_add",
			"1.0",
			"",
			"",
			"",
			types.NEOVM,
		)
		if err != nil {
			ctx.LogError("TestFnCall_contract_add DeploySmartContract error:%s", err)
			return false
		}
	}
	{
		code := app_call
		_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
			code,
			[]contract.ContractParameterType{contract.Integer},
			contract.ContractParameterType(contract.Integer),
			"TestFnCall_app_call",
			"1.0",
			"",
			"",
			"",
			types.NEOVM,
		)
		if err != nil {
			ctx.LogError("TestFnCall_app_call DeploySmartContract error:%s", err)
			return false
		}
	}
	//等待出块
	_, err := ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestFnCall WaitForGenerateBlock error:%s", err)
		return false
	}

	// if !testFnCall(ctx, contract_add, 2, 3) {
	// 	return false
	// }

	// if !testFnCall(ctx, contract_add, -23, -1) {
	// 	return false
	// }

	// if !testFnCall(ctx, contract_add, -1, -1) {
	// 	return false
	// }

	if !testAppCall(ctx, app_call, -1) {
		return false
	}

	// if !testAppCall(ctx, app_call, 10) {
	// 	return false
	// }
	// if !testAppCall(ctx, app_call, -10) {
	// 	return false
	// }

	return true
}

func testFnCall(ctx *testframework.TestFrameworkContext, code string, a int, b int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("testFnCall InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, a+b)
	if err != nil {
		ctx.LogError("testFnCall test for %d failed %s", a, err)
		return false
	}
	return true
}

func testAppCall(ctx *testframework.TestFrameworkContext, code string, a int) bool {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("testAppCall InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToInt(res, 9-a)
	if err != nil {
		ctx.LogError("testAppCall test for %d failed %s", a, err)
		return false
	}
	return true
}

var (
	/*
		using Neo.SmartContract.Framework;
		using Neo.SmartContract.Framework.Services.Neo;

		public class HelloWorld : SmartContract
		{
		    public static int Main(int a, int b)
		    {
		        return add(a, b);
		    }

		    public static int add(int a, int b) {
		        return a + b;
		    }
		}
	*/
	contract_add      = `53c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c3617c6516006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c3936c766b52527ac46203006c766b52c3616c7566`
	contract_add_hash = `27f84b01bbe52463c0487f59d0f2cb7b76b37978`

	/*
	   using Neo.SmartContract.Framework;
	   using Neo.SmartContract.Framework.Services.Neo;

	   public class Lock : SmartContract
	   {
	       [Appcall("xxxxxxxxxxxxxxxxx")]
	       public static extern int contract_add(int a, int b); // add two ints
	       public static int Main(int d)
	       {
	           int c = contract_add(-1, 10);
	           return sub(c, d);
	       }

	       public static int sub(int a, b) {
	           return a - b;
	       }
	   }
	*/

	app_call = `53c56b6c766b00527ac4614f5a617c6727f84b01bbe52463c0487f59d0f2cb7b76b379786c766b51527ac46c766b51c36c766b00c3617c6516006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c3946c766b52527ac46203006c766b52c3616c7566`
)
