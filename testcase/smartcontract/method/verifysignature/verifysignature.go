package verifysignature
//
//import (
//	"github.com/ONT_TEST/testframework"
//	"github.com/Ontology/core/contract"
//	"github.com/Ontology/crypto"
//	"github.com/Ontology/executionengine/types"
//	"time"
//)
//
//func TestVerifySingleSignature(ctx *testframework.TestFrameworkContext) bool {
//	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C3AC616C7566"
//	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
//		code,
//		[]contract.ContractParameterType{contract.ByteArray, contract.ByteArray},
//		contract.ContractParameterType(contract.Boolean),
//		"TestVerifySingleSignature",
//		"1.0",
//		"",
//		"",
//		"",
//		types.NEOVM,
//	)
//	if err != nil {
//		ctx.LogError("TestVerifySingleSignature DeploySmartContract error:%s", err)
//		return false
//	}
//	//等待出块
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("TestVerifySingleSignature WaitForGenerateBlock error:%s", err)
//		return false
//	}
//
//	account := ctx.OntClient.Account1
//	pk , err := account.PublicKey.EncodePoint(true)
//	if err != nil {
//		ctx.LogError("TestVerifySingleSignature EncodePoint PublicKey error:%s", err)
//		return false
//	}
//
//
//	if !testVerifySingleSignature(ctx, code, data, sig, pk, *account.PublicKey) {
//		return false
//	}
//
//	//account = ctx.OntClient.Account2
//	//sig, err = crypto.Sign(account.PrivateKey, []byte("Hello world"))
//	//if err != nil {
//	//	ctx.LogError("TestVerifySingleSignature Account2 sign error:%s", err)
//	//	return false
//	//}
//	//if !testVerifySingleSignature(ctx, code, data, sig, pk, *account.PublicKey) {
//	//	return false
//	//}
//	return true
//}
//
//func testVerifySingleSignature(ctx *testframework.TestFrameworkContext, code string, data, sig, pk []byte, publicKey crypto.PubKey) bool {
//	account := ctx.OntClient.Account1
//	res, err := ctx.Ont.InvokeSmartContract(
//		account,
//		code,
//		[]interface{}{sig, pk},
//	)
//
//	if err != nil {
//		ctx.LogError("testVerifySingleSignature InvokeSmartContract error:%s", err)
//		return false
//	}
//	err = ctx.AssertToBoolean(res, verifySingleSignature(data, sig, publicKey))
//	if err != nil {
//		ctx.LogError("testVerifySingleSignature test failed %s", err)
//		return false
//	}
//	return true
//}
//
//func verifySingleSignature(data, sig []byte, pk crypto.PubKey) bool {
//	err := crypto.Verify(pk, data, sig)
//	return err == nil
//}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static bool Main(byte[] sig, byte[] pk)
    {
        return VerifySignature(sig, pk);
    }
}
*/
