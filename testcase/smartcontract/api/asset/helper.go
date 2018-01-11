package asset

import (
	"encoding/hex"
	"fmt"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	as "github.com/Ontology/core/asset"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"github.com/Ontology/core/transaction/utxo"
)

var (
	isInit         = false
	assetId        common.Uint256
)

func GetAssetId()*common.Uint256{
	return &assetId
}

func InitAsset(ctx *testframework.TestFrameworkContext) error {
	if isInit {
		return nil
	}

	assetType := as.Token
	assetName := "tst"
	assetAmount := 1000000.0
	assetPrecision := byte(8)
	assetOwner := ctx.OntClient.Admin
	assetAdmin := ctx.OntClient.Admin
	assetIssuer := ctx.OntClient.Account1

	var err error
	assetId, err = CreateAsset(
		ctx,
		assetType,
		assetName,
		assetAmount,
		assetPrecision,
		assetOwner,
		assetAdmin,
		assetIssuer)
	if err != nil {
		return fmt.Errorf("CreateAsset error:%s", err)
	}

	isInit = true
	return nil
}

/**
using Neo.SmartContract.Framework;
using System.Numerics;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static byte[] Main(BigInteger type, string name, BigInteger amount, BigInteger precision, byte[] owner, byte[] admin, byte[] issuer)
    {
        Asset asset = Asset.Create((byte)type, name, (long)amount, (byte)precision, owner, admin, issuer);
        return asset.AssetId;
    }
}
*/

var createAssetCode = "59c56b6c766b00527ac46c766b51527ac46c766b52527ac46c766b53527ac46c766b54527ac46c766b55527ac46c766b56527ac4616c766b00c36c766b51c36c766b52c36c766b53c36c766b54c36c766b55c36c766b56c36156795179587275517275557952795772755272755479537956727553727568104e656f2e41737365742e4372656174656c766b57527ac46c766b57c36168144e656f2e41737365742e476574417373657449646c766b58527ac46203006c766b58c3616c7566"

func CreateAsset(ctx *testframework.TestFrameworkContext,
	assetType as.AssetType,
	name string,
	amount float64,
	precision byte,
	owner, admin, issuer *account.Account) (assetId common.Uint256, err error) {

	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		createAssetCode,
		[]contract.ContractParameterType{
			contract.Integer,
			contract.String,
			contract.Integer,
			contract.Integer,
			contract.ByteArray,
			contract.ByteArray,
			contract.ByteArray},
		contract.ContractParameterType(contract.ByteArray),
		"TestAssetCreate",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("DeploySmartContract error:%s", err)
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}

	ownerPk, err := owner.PublicKey.EncodePoint(true)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("PublicKey.EncodePoint error:%s", err)
	}
	res, err := ctx.Ont.InvokeSmartContract(
		owner,
		createAssetCode,
		[]interface{}{
			int(assetType),
			name,
			ctx.Ont.MakeAssetAmount(amount, precision),
			int(precision),
			ownerPk,
			admin.ProgramHash.ToArray(),
			issuer.ProgramHash.ToArray(),
		},
	)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("InvokeSmartContract error:%s", err)
	}

	data, ok := res.(string)
	if !ok {
		return common.Uint256{}, fmt.Errorf("asset to string failed")
	}

	c, _ := hex.DecodeString(data)
	assetId, err = common.Uint256ParseFromBytes(c)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("Uint256ParseFromBytes error:%s", err)
	}

	return
}

func IssueAsset(ctx *testframework.TestFrameworkContext, assetId common.Uint256, amount float64, issuer *account.Account, toAccount *account.Account) error {
	output := &utxo.TxOutput{
		Value:       ctx.Ont.MakeAssetAmount(float64(amount)),
		AssetID:     assetId,
		ProgramHash: toAccount.ProgramHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ctx.Ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		return fmt.Errorf("NewIssueAssetTransaction error:%s", err)
	}
	_, err = ctx.Ont.SendTransaction(issuer, issueTx)
	if err != nil {
		return fmt.Errorf("SendTransaction error:%s", err)
	}
	_, err = ctx.Ont.WaitForGenerateBlock(time.Second * 30)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return nil
}
