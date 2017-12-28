package asset

import (
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	. "github.com/Ontology/core/asset"
	"github.com/Ontology/core/transaction/utxo"
	. "github.com/ONT_TEST/ont"
	"github.com/ONT_TEST/testframework"
	"errors"
	"fmt"
	"time"
)

var BenchmarkClient = 50
var BenchmarkTime = time.Second * 30

func BenchmarkTransaction(ctx *testframework.TestFrameworkContext) bool {
	//assetId, asset, err := registAsset(ctx.Ont, 1000000, "BS01", 4, ctx.OntClient.Admin, ctx.OntClient.Admin)
	//if err != nil {
	//	ctx.LogError("RegisterAsset error:%s", err)
	//	return false
	//}

	clientNum := BenchmarkClient
	testClient := func(id int, exitCh chan interface{}) {
		for {
			select {
			case <-exitCh:
				return
			default:
				//err := issueAsset(ctx.Ont, 0.1, assetId, asset, ctx.OntClient.Admin, ctx.OntClient.Account1)
			err := record(ctx.Ont, ctx.OntClient.Account1)
				if err != nil {
					ctx.LogError("Client:%v IssueAsset error:%s", id, err)
				}
			}
		}
	}

	exitCh := make(chan interface{}, 0)
	benchmarkTime := BenchmarkTime
	for i := 0; i < clientNum; i++ {
		go testClient(i, exitCh)
	}

	startBlockHeight, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("GetBlockCount error:%s", err)
		return false
	}
	startBlockHeight--
	ctx.LogInfo("StartBlockHeight:%v", startBlockHeight)
	benchmarkTimer := time.NewTimer(benchmarkTime)
	<-benchmarkTimer.C

	endBlockHeight, err := ctx.Ont.GetBlockCount()
	if err != nil {
		ctx.LogError("GetBlockCount error:%s", err)
		return false
	}
	endBlockHeight--
	//Stop
	close(exitCh)
	ctx.LogInfo("EndBlockHeight:%v", endBlockHeight)
	tps, err := calculateTPS(ctx, startBlockHeight, endBlockHeight)
	if err != nil {
		ctx.LogError("calculateTPS error:%s", err)
		return false
	}

	ctx.LogInfo("TPS:%v.", tps)
	return true
}

func calculateTPS(ctx *testframework.TestFrameworkContext, startBlockHeight, endBlockHeight uint32) (int, error) {
	var startTime, endTime uint32
	transactionNum := 0
	for i := startBlockHeight; i <= endBlockHeight; i++ {
		block, err := ctx.Ont.GetBlockByHeight(i)
		if err != nil {
			ctx.LogError("GetBlockByHeight:%v error:%s", i, err)
			continue
		}
		switch i {
		case startBlockHeight:
			startTime = block.Blockdata.Timestamp
		default:
			endTime = block.Blockdata.Timestamp
		}
		transactionNum += len(block.Transactions) - 1
	}
	timeCost := endTime - startTime
	if timeCost == 0 {
		return 0, errors.New("Timecose is 0.")
	}
	return int(transactionNum / int(timeCost)), nil
}

func registAsset(ont *Ontology,
	amount float64,
	assetName string,
	precision byte,
	issuer, controller *account.Account) (assetId common.Uint256, asset *Asset, err error) {

	assetPrecise := precision
	assetType := Token
	recordType := UTXO
	asset = ont.CreateAsset(assetName, assetPrecise, assetType, recordType)
	assetAmount := ont.MakeAssetAmount(amount)
	regTx, err := ont.NewAssetRegisterTransaction(asset, assetAmount, issuer, controller)
	if err != nil {
		return
	}

	assetId, err = ont.SendTransaction(issuer, regTx)
	if err != nil {
		err = fmt.Errorf("SendTransaction RegTx error:%s", err)
		return
	}

	_, err = ont.WaitForGenerateBlock(time.Second * 10)
	if err != nil {
		err = fmt.Errorf("WaitForGenerateBlock RegTx error:%s", err)
		return
	}

	_, err = ont.GetTransaction(assetId)
	if err != nil {
		err = fmt.Errorf("GetTransaction RegTx Hash:%x error:%s", assetId, err)
		return
	}
	return
}

func issueAsset(ont *Ontology,
	amount float64,
	assetId common.Uint256,
	asset *Asset,
	controller, toAccount *account.Account) error {
	programHash, err := ont.GetAccountProgramHash(toAccount)
	if err != nil {
		return fmt.Errorf("GetAccountProgramHash toAccount error:%s", err)
	}
	output := &utxo.TxOutput{
		Value:       ont.MakeAssetAmount(amount),
		AssetID:     assetId,
		ProgramHash: programHash,
	}
	txOutputs := []*utxo.TxOutput{output}
	issueTx, err := ont.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		return fmt.Errorf("NewIssueAssetTransaction error:%s", err)
	}
	_, err = ont.SendTransaction(controller, issueTx)
	if err != nil {
		return fmt.Errorf("SendTransaction error:%s", err)
	}
	return nil
}

func record(ont *Ontology, account *account.Account) error {
	recordType := "TestRecord"
	recordData := make([]byte, 1024)
	recordTx, err := ont.NewRecordTransaction(recordType, recordData)
	if err != nil {
		return err
	}

	_, err = ont.SendTransaction(account, recordTx)
	return err
}
