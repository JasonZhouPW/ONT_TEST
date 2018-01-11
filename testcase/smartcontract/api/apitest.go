package api

import (
	. "github.com/ONT_TEST/testframework"
	//"github.com/ONT_TEST/testcase/executionengine/api/header"
	//"github.com/ONT_TEST/testcase/executionengine/api/block"
	//"github.com/ONT_TEST/testcase/executionengine/api/account"
	//"github.com/ONT_TEST/testcase/executionengine/api/asset"
	//"github.com/ONT_TEST/testcase/executionengine/api/attribute"
	//"github.com/ONT_TEST/testcase/executionengine/api/input"
	//"github.com/ONT_TEST/testcase/executionengine/api/blockchain"
	//"github.com/ONT_TEST/testcase/executionengine/api/account"
	//"github.com/ONT_TEST/testcase/executionengine/api/executionengine"
	//"github.com/ONT_TEST/testcase/executionengine/api/transaction"
	//"github.com/ONT_TEST/testcase/executionengine/api/block"
	//"github.com/ONT_TEST/testcase/executionengine/api/blockchain"
	//"github.com/ONT_TEST/testcase/executionengine/api/header"
	//"github.com/ONT_TEST/testcase/executionengine/api/blockchain"
	//"github.com/ONT_TEST/testcase/executionengine/api/header"
	//"github.com/ONT_TEST/testcase/smartcontract/api/blockchain"
	"github.com/ONT_TEST/testcase/smartcontract/api/event"
)

func TestAPI() {
	//blockchain
	//TFramework.RegTestCase("TestGetCurrentHeight", blockchain.TestGetCurrentHeight)
	//TFramework.RegTestCase("TestGetHeader", blockchain.TestGetHeader)
	//TFramework.RegTestCase("TestGetBlock", blockchain.TestGetBlock)
	//TFramework.RegTestCase("TestGetTransaction", blockchain.TestGetTransaction)
	//TFramework.RegTestCase("TestGetContract", blockchain.TestGetContract)
	//TFramework.RegTestCase("TestValidators", blockchain.TestValidators)

	//header
	//TFramework.RegTestCase("TestGetHeader", header.TestGetHeader)

	//block
	//TFramework.RegTestCase("TestGetBlockTransctionCount", block.TestGetBlockTransctionCount)
	//TFramework.RegTestCase("TestGetBlockTransactions", block.TestGetBlockTransactions)
	//TFramework.RegTestCase("TestGetBlockTransaction", block.TestGetBlockTransaction)

	//account
	//TFramework.RegTestCase("TestGetBalance", account.TestGetBalance)
	//TFramework.RegTestCase("TestGetScriptHash", account.TestGetScriptHash)

	//asset
	//TFramework.RegTestCase("TestAssetCreate", asset.TestAssetCreate)
	//TFramework.RegTestCase("TestAsset", asset.TestAsset)
	//TFramework.RegTestCase("TestRenew", asset.TestRenew)

	//attribute
	//TFramework.RegTestCase("TestGetAttributeData", attribute.TestGetAttributeData)

	//inputs
	//TFramework.RegTestCase("TestGetInputHash", input.TestGetInputHash)
	TFramework.RegTestCase("TestLog", event.TestLog)
	TFramework.RegTestCase("TestNotify", event.TestNotify)



	//triggetype
	//TFramework.RegTestCase("TestTriggerType",executionengine.TestTriggerType)
	//TFramework.RegTestCase("TestCallContractStatic", executionengine.TestCallContractStatic)
	//TFramework.RegTestCase("TestCallingScriptHash", executionengine.TestCallingScriptHash)
	//TFramework.RegTestCase("TestCheckWitness", executionengine.TestCheckWitness)
	//TFramework.RegTestCase("TestExecutingScriptHash", executionengine.TestExecutingScriptHash)
	//TFramework.RegTestCase("TestEntryScriptHash", executionengine.TestEntryScriptHash)
	//TFramework.RegTestCase("TestExecutionEngine", executionengine.TestExecutionEngine)

	//Transction
	//TFramework.RegTestCase("TestGetAttributes", transaction.TestGetAttributes)
	//TFramework.RegTestCase("TestGetInputs", transaction.TestGetInputs)
	//TFramework.RegTestCase("TestGetOutputs", transaction.TestGetOutputs)
	//TFramework.RegTestCase("TestGetReference", transaction.TestGetReference)
	//TFramework.RegTestCase("TestGetTxType", transaction.TestGetTxType)
	//TFramework.RegTestCase("TestGetTxHash", transaction.TestGetTxHash)
}


