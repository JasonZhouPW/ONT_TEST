package header

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object[] Main(int height)
    {
        object[] ret = new object[8];
        Header header = Blockchain.GetHeader((uint)height);
        ret[0] = header.ConsensusData;
        ret[1] = header.Hash;
        ret[2] = header.Index;
        ret[3] = header.MerkleRoot;
        ret[4] = header.NextConsensus;
        ret[5] = header.PrevHash;
        ret[6] = header.Timestamp;
        ret[7] = header.Version;
        return ret;
    }
}
*/

func TestGetHeader(ctx *testframework.TestFrameworkContext) bool {
	code := "54c56b6c766b00527ac46158c56c766b51527ac46c766b00c36168184e656f2e426c6f636b636861696e2e4765744865616465726c766b52527ac46c766b51c3006c766b52c361681b4e656f2e4865616465722e476574436f6e73656e73757344617461c46c766b51c3516c766b52c36168124e656f2e4865616465722e47657448617368c46c766b51c3526c766b52c36168134e656f2e4865616465722e476574496e646578c46c766b51c3536c766b52c36168184e656f2e4865616465722e4765744d65726b6c65526f6f74c46c766b51c3546c766b52c361681b4e656f2e4865616465722e4765744e657874436f6e73656e737573c46c766b51c3556c766b52c36168164e656f2e4865616465722e4765745072657648617368c46c766b51c3566c766b52c36168174e656f2e4865616465722e47657454696d657374616d70c46c766b51c3576c766b52c36168154e656f2e4865616465722e47657456657273696f6ec46c766b51c36c766b53527ac46203006c766b53c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Array),
		"TestGetHeader",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetHeader DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetHeader WaitForGenerateBlock error:%s", err)
		return false
	}

	height, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("TestGetHeader GetBlockCount error:%s", err)
		return false
	}

	height -= 1
	block, err := ctx.Ont.GetBlockByHeight(height)
	if err != nil {
		ctx.LogError("TestGetHeader GetBlockByHeight error:%s", err)
		return false
	}

	header := block.Blockdata
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{int(height)},
	)
	if err != nil {
		ctx.LogError("TestGetHeader InvokeSmartContract error:%s", err)
		return false
	}

	ret, ok := res.([]interface{})
	if !ok {
		ctx.LogError("TestGetHeader asset ret to []interface{} failed")
		return false
	}

	//err = ctx.AssertToUint(ret[0], uint(heard.ConsensusData))
	//if err != nil {
	//	ctx.LogError("TestGetHeader consensusData AssertToUint error:%s", err)
	//	return false
	//}

	hash := header.Hash()
	err = ctx.AssertToByteArray(ret[1], hash.ToArray())
	if err != nil {
		ctx.LogError("TestGetHeader Hash AssertToByteArray error:%s", err)
		return false
	}

	err = ctx.AssertToInt(ret[2], int(header.Height))
	if err != nil {
		ctx.LogError("TestGetHeader Height AssertToInt error:%s", err)
		return false
	}

	err = ctx.AssertToByteArray(ret[3], header.TransactionsRoot.ToArray())
	if err != nil {
		ctx.LogError("TestGetHeader TransactionsRoot AssertToByteArray error:%s", err)
		return false
	}

	err = ctx.AssertToByteArray(ret[4], header.NextBookKeeper.ToArray())
	if err != nil {
		ctx.LogError("TestGetHeader NextBookKeeper AssertToByteArray error:%s", err)
		return false
	}

	err = ctx.AssertToByteArray(ret[5], header.PrevBlockHash.ToArray())
	if err != nil {
		ctx.LogError("TestGetHeader PrevBlockHash AssertToByteArray error:%s", err)
		return false
	}

	err = ctx.AssertToInt(ret[6], int(header.Timestamp))
	if err != nil {
		ctx.LogError("TestGetHeader Timestamp AssertToInt error:%s", err)
		return false
	}

	err = ctx.AssertToInt(ret[7], int(header.Version))
	if err != nil {
		ctx.LogError("TestGetHeader Version AssertToInt error:%s", err)
		return false
	}
	return true
}
