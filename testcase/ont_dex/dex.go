package ont_dex

import (
	"github.com/ONT_TEST/testframework"
	"fmt"
	"github.com/Ontology/crypto"
	"math/rand"
)

func TestOntDex(){
	testframework.TFramework.RegTestCase("TestOntDex", TestOntDexInter)
}

func TestOntDexInter(ctx *testframework.TestFrameworkContext)bool{
	if !deployDexFund(ctx){
		return false
	}
	admin := ctx.OntClient.Admin
	assetId := []byte("")
	if !dexFundInit(ctx, assetId, admin){
		return false
	}
	//if !setFundCaller(ctx, admin, DExProtoCode_Hash){
	//	return false
	//}
	if !deployDexProto(ctx){
		return false
	}
	if !dexProtoInit(ctx, admin){
		return false
	}
	if !addProtoCaller(ctx,admin, DEXP2PCodeHashReverse){
		return false
	}
	if !deployDexP2P(ctx){
		return false
	}
	if !dexP2PInit(ctx){
		return false
	}

	buyer := ctx.OntClient.Account1
	seller := ctx.OntClient.Account2
	orderId := []byte(fmt.Sprint("%d", rand.Int31()))
	orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
	if err != nil {
		ctx.LogError("TestOntDexInter crypto.Sign error:%s", err)
		return false
	}
	amount := 10
	if !fundReceipt(ctx, buyer, amount){
		return false
	}
	if !makeBuyOrder(ctx,orderSig,orderId,buyer,seller, amount ){
		return false;
	}
	if !buyOrderComplete(ctx, orderId, buyer){
		return false;
	}

	amount = 11
	if !fundReceipt(ctx, buyer, amount){
		return false
	}
	if !makeBuyOrder(ctx,orderSig,orderId,buyer,seller, amount ){
		return false;
	}
	if !buyOrderCancel(ctx, orderId, buyer){
		return false;
	}
	return true
}
