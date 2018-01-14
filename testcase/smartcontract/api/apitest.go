package api

import (
	"github.com/ONT_TEST/testcase/smartcontract/api/account"
	"github.com/ONT_TEST/testcase/smartcontract/api/asset"
	"github.com/ONT_TEST/testcase/smartcontract/api/runtime"
	"github.com/ONT_TEST/testcase/smartcontract/api/transaction"
	. "github.com/ONT_TEST/testframework"
	"github.com/ONT_TEST/testcase/smartcontract/api/blockchain"
	"github.com/ONT_TEST/testcase/smartcontract/api/header"
	"github.com/ONT_TEST/testcase/smartcontract/api/block"
	"github.com/ONT_TEST/testcase/smartcontract/api/executionengine"
)

func TestAPI() {
	//blockchain
	TFramework.RegTestCase("TestGetCurrentHeight", blockchain.TestGetCurrentHeight)
	TFramework.RegTestCase("TestGetBlock", blockchain.TestGetBlock)
	TFramework.RegTestCase("TestGetContract", blockchain.TestGetContract)
	//TFramework.RegTestCase("TestValidators", blockchain.TestValidators)
	//
	//header
	TFramework.RegTestCase("TestGetHeader", header.TestGetHeader)
	//
	//block
	TFramework.RegTestCase("TestGetBlockTransactions", block.TestGetBlockTransactions)
	//
	////account
	TFramework.RegTestCase("TestGetBalance", account.TestGetBalance)
	TFramework.RegTestCase("TestGetScriptHash", account.TestGetScriptHash)

	//asset
	TFramework.RegTestCase("TestAssetCreate", asset.TestAssetCreate)
	TFramework.RegTestCase("TestAsset", asset.TestAsset)
	TFramework.RegTestCase("TestRenew", asset.TestRenew)

	//runtime
	TFramework.RegTestCase("TestTriggerType", runtime.TestTriggerType)
	TFramework.RegTestCase("TestCheckWitness", runtime.TestCheckWitness)
	TFramework.RegTestCase("TestLog", runtime.TestLog)
	TFramework.RegTestCase("TestNotify", runtime.TestNotify)

	TFramework.RegTestCase("TestCallingScriptHash", executionengine.TestCallingScriptHash)
	TFramework.RegTestCase("TestExecutingScriptHash", executionengine.TestExecutingScriptHash)
	TFramework.RegTestCase("TestEntryScriptHash", executionengine.TestEntryScriptHash)
	TFramework.RegTestCase("TestExecutionEngine", executionengine.TestExecutionEngine)

	//Transction
	TFramework.RegTestCase("TestGetAttributes", transaction.TestGetAttributes)
	TFramework.RegTestCase("TestGetInputs", transaction.TestGetInputs)
	TFramework.RegTestCase("TestGetOutputs", transaction.TestGetOutputs)
	TFramework.RegTestCase("TestGetReference", transaction.TestGetReference)
	TFramework.RegTestCase("TestGetTxType", transaction.TestGetTxType)
	TFramework.RegTestCase("TestGetTxHash", transaction.TestGetTxHash)
}
