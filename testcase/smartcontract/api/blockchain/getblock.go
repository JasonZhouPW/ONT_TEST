package blockchain


import (
	"time"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"github.com/ONT_TEST/testframework"
)

/**
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static object[] Main(int height)
    {
        object[] ret = new object[8];
        Block block = Blockchain.GetBlock((uint)height);
        ret[0] = block.ConsensusData;
        ret[1] = block.Hash;
        ret[2] = block.Index;
        ret[3] = block.MerkleRoot;
        ret[4] = block.NextConsensus;
        ret[5] = block.PrevHash;
        ret[6] = block.Timestamp;
        ret[7] = block.Version;
        return ret;
    }
}
*/

func TestGetBlock(ctx *testframework.TestFrameworkContext) bool {
	code := "54c56b6c766b00527ac46158c56c766b51527ac46c766b00c36168174e656f2e426c6f636b636861696e2e476574426c6f636b6c766b52527ac46c766b51c3006c766b52c361681b4e656f2e4865616465722e476574436f6e73656e73757344617461c46c766b51c3516c766b52c36168124e656f2e4865616465722e47657448617368c46c766b51c3526c766b52c36168134e656f2e4865616465722e476574496e646578c46c766b51c3536c766b52c36168184e656f2e4865616465722e4765744d65726b6c65526f6f74c46c766b51c3546c766b52c361681b4e656f2e4865616465722e4765744e657874436f6e73656e737573c46c766b51c3556c766b52c36168164e656f2e4865616465722e4765745072657648617368c46c766b51c3566c766b52c36168174e656f2e4865616465722e47657454696d657374616d70c46c766b51c3576c766b52c36168154e656f2e4865616465722e47657456657273696f6ec46c766b51c36c766b53527ac46203006c766b53c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Integer},
		contract.ContractParameterType(contract.Array),
		"TestGetBlock",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestGetBlock DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestGetBlock WaitForGenerateBlock error:%s", err)
		return false
	}

	height, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("TestGetBlock GetBlockCount error:%s", err)
		return false
	}

	height -= 1
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{height},
	)
	if err != nil {
		ctx.LogError("TestGetBlock InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestGetBlock res:%s", res)

	return true
}
