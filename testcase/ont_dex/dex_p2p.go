package ont_dex

import (
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"github.com/Ontology/crypto"
)

func TestDexP2P(ctx *TestFrameworkContext) bool {
	code := "0117c56b6c766b00527ac4610c6d616b656275796f726465726c766b51527ac46c766b51c30c6d616b656275796f72646572876c766b52527ac46c766b52c364c700616c766b00c3c0569c009c6c766b59527ac46c766b59c3640f0061516c766b5a527ac4623c026c766b00c300c36c766b53527ac46c766b00c351c36c766b54527ac46c766b00c352c36c766b55527ac46c766b00c353c36c766b56527ac46c766b00c354c36c766b57527ac46c766b00c355c36c766b58527ac46c766b53c36c766b54c36c766b55c36c766b56c36c766b57c36c766b58c36155795179577275517275547952795672755272755379537955727553727565b1016c766b5a527ac4629e016c766b51c3106275796f72646572636f6d706c657465876c766b5b527ac46c766b5bc3645e00616c766b00c3c0529c009c6c766b5e527ac46c766b5ec3640f0061516c766b5a527ac46252016c766b00c300c36c766b5c527ac46c766b00c351c36c766b5d527ac46c766b5cc36c766b5dc3617c6572016c766b5a527ac4621d016c766b51c30e6275796f7264657263616e63656c876c766b5f527ac46c766b5fc3646200616c766b00c3c0529c009c6c766b0112527ac46c766b0112c3640f0061516c766b5a527ac462d1006c766b00c300c36c766b60527ac46c766b00c351c36c766b0111527ac46c766b60c36c766b0111c3617c6515016c766b5a527ac4629a006c766b51c31373656c6c6572747279636c6f73656f72646572876c766b0113527ac46c766b0113c3646400616c766b00c3c0529c009c6c766b0116527ac46c766b0116c3640f0061516c766b5a527ac46247006c766b00c300c36c766b0114527ac46c766b00c351c36c766b0115527ac46c766b0114c36c766b0115c3617c65af006c766b5a527ac4620e00526c766b5a527ac46203006c766b5ac3616c756657c56b6c766b00527ac46c766b51527ac46c766b52527ac46c766b53527ac46c766b54527ac46c766b55527ac461006c766b56527ac46203006c766b56c3616c756653c56b6c766b00527ac46c766b51527ac461006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac461006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac461006c766b52527ac46203006c766b52c3616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{ contract.Array},
		contract.ContractParameterType(contract.Integer),
		"TestDExP2P",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestHash160 DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestHash160 WaitForGenerateBlock error:%s", err)
		return false
	}

	buyer := ctx.OntClient.Account1
	seller := ctx.OntClient.Account2
	orderId := []byte("223342534534543545")
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestDexP2P crypto.Sign error:%s", err)
		return false
	}
	amount := 10
	if !testMakeBuyOrder(ctx, code, orderSig, orderId, buyer, seller, amount) {
		return false
	}
	return true
}

func testMakeBuyOrder(ctx *TestFrameworkContext, code string, orderSig, orderId []byte, buyer, seller *account.Account, amount int) bool {
	//operation := "makebuyorder"
	//buyerPk, err := buyer.PublicKey.EncodePoint(true)
	//if err != nil {
	//	ctx.LogError("PublicKey.EncodePoint error:%s", err)
	//	return false
	//}
	res, err := ctx.Ont.InvokeSmartContract(
		seller,
		code,
		[]interface{}{[]interface{}{[]byte{3,4,5}, []byte{1,2,3}}},
		//[]interface{}{[]interface{}{orderSig, orderId, buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), buyerPk, amount}},
	)
	if err != nil {
		ctx.LogError("TestMakeBuyOrder error:%s", err)
		return false
	}
	ctx.LogInfo("TestMakeBuyOrder res:%v", res)
	return true
}
