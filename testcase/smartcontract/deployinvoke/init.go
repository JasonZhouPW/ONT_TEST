package deployinvoke

import (
	. "github.com/ONT_TEST/testframework"
)

func TestDeployInvoke(){
	TFramework.RegTestCase("TestDeploySmartContract", TestDeploySmartContract)
	TFramework.RegTestCase("TestInvokeSmartContract", TestInvokeSmartContract)
}

