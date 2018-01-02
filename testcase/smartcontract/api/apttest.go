package api

import (
	. "github.com/ONT_TEST/testframework"
	//"github.com/ONT_TEST/testcase/smartcontract/api/header"
	//"github.com/ONT_TEST/testcase/smartcontract/api/block"
	//"github.com/ONT_TEST/testcase/smartcontract/api/block"
	//"github.com/ONT_TEST/testcase/smartcontract/api/account"
	//"github.com/ONT_TEST/testcase/smartcontract/api/asset"
	//"github.com/ONT_TEST/testcase/smartcontract/api/attribute"
	//"github.com/ONT_TEST/testcase/smartcontract/api/input"
	//"github.com/ONT_TEST/testcase/smartcontract/api/blockchain"
	//"github.com/ONT_TEST/testcase/smartcontract/api/account"
	"github.com/ONT_TEST/testcase/smartcontract/api/sys"
)

func TestAPI() {
	//blockchain
	//TFramework.RegTestCase("TestGetCurrentHeight", blockchain.TestGetCurrentHeight)
	//TFramework.RegTestCase("TestGetHeader", blockchain.TestGetHeader)
	//TFramework.RegTestCase("TestGetBlock", blockchain.TestGetBlock)
	//TFramework.RegTestCase("TestGetTransaction", blockchain.TestGetTransaction)
	//TFramework.RegTestCase("TestGetAccount", blockchain.TestGetAccount)
	//TFramework.RegTestCase("TestGetAsset", blockchain.TestGetAsset)
	//TFramework.RegTestCase("TestGetContract", blockchain.TestGetContract)
	//TFramework.RegTestCase("TestValidators", blockchain.TestValidators)

	//header
	//TFramework.RegTestCase("TestGetHeaderHash", header.TestGetHeaderHash)
	//TFramework.RegTestCase("TestGetHeaderVersion", header.TestGetHeaderVersion)
	//TFramework.RegTestCase("TestGetPrevHash", header.TestGetPrevHash)
	//TFramework.RegTestCase("TestGetMerkleRoot", header.TestGetMerkleRoot)
	//TFramework.RegTestCase("TestGetHeaderTimeStamp", header.TestGetHeaderTimeStamp)
	//TFramework.RegTestCase("TestGetConsensusData", header.TestGetConsensusData)
	//TFramework.RegTestCase("TestGetNextConsus", header.TestGetNextConsus)

	//block
	//TFramework.RegTestCase("TestGetBlockTransctionCount", block.TestGetBlockTransctionCount)
	//TFramework.RegTestCase("TestGetBlockTransactions", block.TestGetBlockTransactions)
	//TFramework.RegTestCase("TestGetBlockTransaction", block.TestGetBlockTransaction)

	//account
	//TFramework.RegTestCase("TestGetBalance", account.TestGetBalance)
	//TFramework.RegTestCase("TestGetScriptHash", account.TestGetScriptHash)

	//asset
	//TFramework.RegTestCase("TestGetAdmin", asset.TestGetAdmin)
	//TFramework.RegTestCase("TestGetAmount", asset.TestGetAmount)
	//TFramework.RegTestCase("TestGetAssetId", asset.TestGetAssetId)
	//TFramework.RegTestCase("TestGetAssetType", asset.TestGetAssetType)
	//TFramework.RegTestCase("TestGetAvailable", asset.TestGetAvailable)
	//TFramework.RegTestCase("TestGetOwner", asset.TestGetOwner)
	//TFramework.RegTestCase("TestGetPrecision", asset.TestGetPrecision)

	//attribute
	//TFramework.RegTestCase("TestGetAttributeData", attribute.TestGetAttributeData)

	//inputs
	//TFramework.RegTestCase("TestGetInputHash", input.TestGetInputHash)

	//triggetype
	TFramework.RegTestCase("TestTriggerType",sys.TestTriggerType)

}

