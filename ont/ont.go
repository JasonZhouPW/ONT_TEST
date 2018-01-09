package ont

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ONT_TEST/utils"
	"github.com/Ontology/account"
	. "github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/core/code"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/core/ledger"
	"github.com/Ontology/core/signature"
	"github.com/Ontology/core/transaction"
	"github.com/Ontology/core/transaction/payload"
	"github.com/Ontology/core/transaction/utxo"
	"github.com/Ontology/crypto"
	"github.com/Ontology/smartcontract/types"
	"github.com/Ontology/vm/neovm"
	log4 "github.com/alecthomas/log4go"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	crypto.SetAlg("P256R1")
}

type Ontology struct {
	qid          uint64
	rpcAddresses []string
	wsAddresses  []string
	client       *http.Client
}

func NewOntology(rpcAddresses, wsAddresses []string) *Ontology {
	return &Ontology{
		rpcAddresses: rpcAddresses,
		wsAddresses:  wsAddresses,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   50,
				DisableKeepAlives:     false, //启动keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
			},
			Timeout: time.Second * 300,
		},
	}
}

func (this *Ontology) GetVersion() (string, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETVERSION, []interface{}{})
	if err != nil {
		return "", fmt.Errorf("SendRpcRequest error:%s", err)
	}
	return string(data), nil
}

func (this *Ontology) CreateAsset(
	name string,
	precision byte,
	assetType asset.AssetType,
	recordType asset.AssetRecordType) *asset.Asset {
	return &asset.Asset{
		Name:       name,
		Precision:  precision,
		AssetType:  assetType,
		RecordType: recordType,
	}
}

func (this *Ontology) GetBlockByHash(hash Uint256) (*ledger.Block, error) {
	blockHash := Uint256ToString(hash)
	data, err := this.sendRpcRequest(ONT_RPC_GETBLOCK, []interface{}{Uint256ToString(hash)})
	if err != nil {
		return nil, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	blockInfo := &BlockInfo{}
	err = json.Unmarshal(data, blockInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal BlockInfo:%s error:%s", blockInfo, err)
	}
	block, err := ParseBlock(blockInfo)
	if err != nil {
		return nil, fmt.Errorf("ParseBlock Hash:%x error:%s", blockHash, err)
	}
	return block, nil
}

func (this *Ontology) GetBlockByHeight(height uint32) (*ledger.Block, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETBLOCK, []interface{}{height})
	if err != nil {
		return nil, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	blockInfo := &BlockInfo{}
	err = json.Unmarshal(data, blockInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal BlockInfo:%s error:%s", blockInfo, err)
	}
	block, err := ParseBlock(blockInfo)
	if err != nil {
		return nil, fmt.Errorf("ParseBlock Hright:%v error:%s", height, err)
	}
	return block, nil
}

func (this *Ontology) GetBlockHash(height uint32) (Uint256, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETBLOCKHASH, []interface{}{height})
	if err != nil {
		return Uint256{}, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	hash, err := ParseUint256FromString(string(data))
	if err != nil {
		return Uint256{}, fmt.Errorf("ParseUint256FromString Hash:%s error:%s", data, err)
	}
	return hash, nil
}

func (this *Ontology) GetCurrentBlockHash() (Uint256, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETCURRENTBLOCKHASH, []interface{}{})
	if err != nil {
		return Uint256{}, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	hash, err := ParseUint256FromString(string(data))
	if err != nil {
		return Uint256{}, fmt.Errorf("ParseUint256FromString:%s error:%s", hash, err)
	}
	return hash, nil
}

func (this *Ontology) GetBlockCount() (uint32, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETBLOCKCOUNT, []interface{}{})
	if err != nil {
		return 0, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	count := uint32(0)
	err = json.Unmarshal(data, &count)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal Count:%s error:%s", data, err)
	}
	return count, nil
}

func (this *Ontology) NewAssetRegisterTransaction(asset *asset.Asset,
	amount Fixed64,
	issuer,
	controllerAccount *account.Account) (*transaction.Transaction, error) {
	controller, err := contract.CreateSignatureContract(controllerAccount.PubKey())
	if err != nil {
		return nil, fmt.Errorf("CreateSignatureContract error:%s", err)
	}
	tx, err := transaction.NewRegisterAssetTransaction(asset, amount, issuer.PubKey(), controller.ProgramHash)
	if err != nil {
		return nil, fmt.Errorf("NewRegisterAssetTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) NewIssueAssetTransaction(txOutputs []*utxo.TxOutput) (*transaction.Transaction, error) {
	tx, err := transaction.NewIssueAssetTransaction(txOutputs)
	if err != nil {
		return nil, fmt.Errorf("NewIssueAssetTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) NewTransferAssetTransaction(inputs []*utxo.UTXOTxInput,
	outputs []*utxo.TxOutput) (*transaction.Transaction, error) {
	tx, err := transaction.NewTransferAssetTransaction(inputs, outputs)
	if err != nil {
		return nil, fmt.Errorf("NewTransferAssetTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) NewRecordTransaction(recordType string, recordData []byte) (*transaction.Transaction, error) {
	tx, err := transaction.NewRecordTransaction(recordType, recordData)
	if err != nil {
		return nil, fmt.Errorf("NewRecordTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) setNonce(tx *transaction.Transaction) {
	attr := transaction.NewTxAttribute(transaction.Nonce, []byte(fmt.Sprintf("%d", rand.Int63())))
	tx.Attributes = append(tx.Attributes, &attr)
}

func (this *Ontology) DoSendTransaction(tx *transaction.Transaction) (Uint256, error) {
	var buffer bytes.Buffer
	err := tx.Serialize(&buffer)
	if err != nil {
		return Uint256{}, fmt.Errorf("Serialize error:%s", err)
	}

	txData := hex.EncodeToString(buffer.Bytes())
	data, err := this.sendRpcRequest(ONT_RPC_SENDTRANSACTION, []interface{}{txData})
	if err != nil {
		return Uint256{}, err
	}
	hash, err := ParseUint256FromString(string(data))
	if err != nil {
		return Uint256{}, errors.New(string(data))
	}
	return hash, nil
}

func (this *Ontology) SendTransaction(signer *account.Account, tx *transaction.Transaction) (Uint256, error) {
	err := this.SignTransaction(tx, []*account.Account{signer})
	if err != nil {
		return Uint256{}, fmt.Errorf("SignTransaction error:%s", err)
	}

	return this.DoSendTransaction(tx)
}

func (this *Ontology) SignTransaction(tx *transaction.Transaction, signers []*account.Account) error {
	programHashes, err := this.GetTransactionProgramHashes(tx)
	if err != nil {
		return fmt.Errorf("GetTransactionProgramHashes error:%s", err)
	}
	if len(programHashes) == 0 {
		return nil
	}
	ctx, err := this.NewContractContext(tx, programHashes)
	if err != nil {
		return fmt.Errorf("NewContractContext error:%s", err)
	}

	for _, signer := range signers {
		sig, err := signature.SignBySigner(tx, signer)
		if err != nil {
			return fmt.Errorf("SignBySigner error:%s", err)
		}
		transactionContract, err := contract.CreateSignatureContract(signer.PubKey())
		if err != nil {
			return fmt.Errorf("CreateSignatureContract error:%s", err)
		}

		err = ctx.AddContract(transactionContract, signer.PubKey(), sig)
		if err != nil {
			return fmt.Errorf("AddContract error:%s", err)
		}
	}

	tx.SetPrograms(ctx.GetPrograms())
	return nil
}

func (this *Ontology) SendMultiSigTransction(owner *account.Account, m int, singers []*account.Account, tx *transaction.Transaction) (Uint256, error) {
	err := this.MultiSignTransaction(owner, m, singers, tx)
	if err != nil {
		return Uint256{}, fmt.Errorf("MultiSignTransaction error:%s", err)
	}

	var buffer bytes.Buffer
	err = tx.Serialize(&buffer)
	if err != nil {
		return Uint256{}, fmt.Errorf("Serialize error:%s", err)
	}

	txData := hex.EncodeToString(buffer.Bytes())
	data, err := this.sendRpcRequest(ONT_RPC_SENDTRANSACTION, []interface{}{txData})
	if err != nil {
		return Uint256{}, err
	}

	hash, err := ParseUint256FromString(string(data))
	if err != nil {
		return Uint256{}, fmt.Errorf("ParseUint256FromString Hash:%s error:%s", data, err)
	}
	return hash, nil
}

func (this *Ontology) MultiSignTransaction(owner *account.Account, m int, signers []*account.Account, tx *transaction.Transaction) error {
	if len(signers) == 0 {
		return fmt.Errorf("not enough signer")
	}
	pubKeys := make([]*crypto.PubKey, 0, len(signers))
	signatures := make([][]byte, 0, len(signers))
	for _, signer := range signers {
		sig, err := signature.SignBySigner(tx, signer)
		if err != nil {
			return fmt.Errorf("SignBySigner error:%s", err)
		}
		signatures = append(signatures, sig)
		pubKeys = append(pubKeys, signer.PubKey())
	}
	transactionContract, err := contract.CreateMultiSigContract(owner.ProgramHash, m, pubKeys)
	if err != nil {
		return fmt.Errorf("CreateMultiSigContract error:%s", err)
	}
	programHashes, err := this.GetTransactionProgramHashes(tx)
	if err != nil {
		return fmt.Errorf("GetTransactionProgramHashes error:%s", err)
	}
	ctx, err := this.NewContractContext(tx, programHashes)
	if err != nil {
		return fmt.Errorf("NewContractContext error:%s", err)
	}
	for _, sig := range signatures {
		err = ctx.AddContract(transactionContract, owner.PubKey(), sig)
		if err != nil {
			return fmt.Errorf("AddContract error:%s", err)
		}
	}
	tx.SetPrograms(ctx.GetPrograms())
	return nil
}

func (this *Ontology) GetTransactionProgramHashes(tx *transaction.Transaction) ([]Uint160, error) {
	hashs := []Uint160{}
	uniqHashes := []Uint160{}
	// add inputUTXO's transaction
	referenceWithUTXO_Output, err := this.GetTransactionReference(tx)
	if err != nil {
		return nil, fmt.Errorf("Transction GetReference error:%s", err)
	}
	for _, output := range referenceWithUTXO_Output {
		programHash := output.ProgramHash
		hashs = append(hashs, programHash)
	}
	for _, attribute := range tx.Attributes {
		if attribute.Usage != transaction.Script {
			continue
		}
		dataHash, err := Uint160ParseFromBytes(attribute.Data)
		if err != nil {
			return nil, fmt.Errorf("Uint160ParseFromBytes error:%s", err)
		}
		hashs = append(hashs, Uint160(dataHash))
	}
	switch tx.TxType {
	case transaction.RegisterAsset:
		issuer := tx.Payload.(*payload.RegisterAsset).Issuer
		signatureRedeemScript, err := contract.CreateSignatureRedeemScript(issuer)
		if err != nil {
			return nil, fmt.Errorf("CreateSignatureRedeemScript error:%s", err)
		}
		astHash, err := ToCodeHash(signatureRedeemScript)
		if err != nil {
			return nil, fmt.Errorf("ToCodeHash error:%s", err)
		}
		hashs = append(hashs, astHash)
	case transaction.IssueAsset:
		result := tx.GetMergedAssetIDValueFromOutputs()
		if err != nil {
			return nil, fmt.Errorf("GetMergedAssetIDValueFromOutputs error:%s", err)
		}
		for k := range result {
			regTx, err := this.GetTransaction(k)
			if err != nil {
				return nil, fmt.Errorf("GetTransaction TxHash:%x error:%s", k, err)
			}
			if regTx.TxType != transaction.RegisterAsset {
				return nil, errors.New("Transaction is not RegisterAsset")
			}

			regPayload := regTx.Payload.(*payload.RegisterAsset)
			hashs = append(hashs, regPayload.Controller)
		}
	case transaction.TransferAsset:
	case transaction.Record:
	case transaction.BookKeeper:
	default:
	}
	//remove dupilicated hashes
	uniq := make(map[Uint160]bool)
	for _, v := range hashs {
		uniq[v] = true
	}
	for k := range uniq {
		uniqHashes = append(uniqHashes, k)
	}
	sort.Sort(ByProgramHashes(uniqHashes))
	return uniqHashes, nil
}

func (this *Ontology) NewContractContext(data signature.SignableData, programHashes ...[]Uint160) (*contract.ContractContext, error) {
	var proHashes []Uint160
	var err error
	if len(programHashes) > 0 {
		proHashes = programHashes[0]
	} else {
		proHashes, err = data.GetProgramHashes()
		if err != nil {
			return nil, fmt.Errorf("GetProgramHashes error:%s", err)
		}
	}
	hashLen := len(proHashes)
	return &contract.ContractContext{
		Data:            data,
		ProgramHashes:   proHashes,
		Codes:           make([][]byte, hashLen),
		Parameters:      make([][][]byte, hashLen),
		MultiPubkeyPara: make([][]contract.PubkeyParameter, hashLen),
	}, nil
}

func (this *Ontology) GetTransaction(txHash Uint256) (*transaction.Transaction, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETTRANSACTION, []interface{}{Uint256ToString(txHash)})
	if err != nil {
		return nil, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	txStr := &Transactions{}
	err = json.Unmarshal(data, txStr)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal Transactions:%s error:%s", data, err)
	}
	tx, err := ParseTransaction(txStr)
	if err != nil {
		return nil, fmt.Errorf("ParseTransaction:%+v error:%s", txStr, err)
	}
	return tx, nil
}

func (this *Ontology) GetUnspendOutput(assetHash Uint256, programHash Uint160) ([]*utxo.UTXOUnspent, error) {
	data, err := this.sendRpcRequest(ONT_RPC_GETUNSPENDOUTPUT, []interface{}{Uint160ToString(programHash), Uint256ToString(assetHash)})
	if err != nil {
		return nil, fmt.Errorf("sendRpcRequest error:%s", err)
	}
	if string(data) == "{}" {
		return nil, nil
	}
	outputs := make([]*UnspendUTXOInfo, 0)
	err = json.Unmarshal(data, &outputs)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal []*UnspendUTXOInfo:%s error:%s", data, err)
	}

	unspents := make([]*utxo.UTXOUnspent, 0, len(outputs))
	for _, output := range outputs {
		txid, err := ParseUint256FromString(output.Txid)
		if err != nil {
			return nil, fmt.Errorf("ParseUint256FromString:%x error:%s", output.Txid, err)
		}
		index, err := strconv.ParseInt(output.Index, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("strconv.ParseInt:%s error:%s", output.Index, err)
		}
		value, err := ParseFixed64FromString(output.Value)
		if err != nil {
			return nil, fmt.Errorf("ParseFixed64FromString value:%s error:%s", value, err)
		}
		unspent := &utxo.UTXOUnspent{
			Txid:  txid,
			Index: uint32(index),
			Value: value,
		}
		unspents = append(unspents, unspent)
	}
	return unspents, nil
}

func (this *Ontology) WaitForGenerateBlock(timeout time.Duration, blockCount ...uint32) (bool, error) {
	count := uint32(2)
	if len(blockCount) > 0 && blockCount[0] > 0 {
		count = blockCount[0]
	}
	blockHeight, err := this.GetBlockCount()
	if err != nil {
		return false, fmt.Errorf("GetBlockCount error:%s", err)
	}
	secs := int(timeout / time.Second)
	if secs <= 0 {
		secs = 1
	}
	ok := false
	for i := 0; i < secs; i++ {
		time.Sleep(time.Second)
		curBlockHeigh, err := this.GetBlockCount()
		if err != nil {
			continue
		}
		if curBlockHeigh-blockHeight >= count {
			ok = true
			break
		}
	}
	return ok, nil
}

func (this *Ontology) MakeAssetAmount(rawAmont float64) Fixed64 {
	return Fixed64(rawAmont * 100000000)
}

func (this *Ontology) GetRawAssetAmount(assetAmount Fixed64) float64 {
	return float64(assetAmount) / 100000000
}

func (this *Ontology) GetAccountProgramHash(account *account.Account) (Uint160, error) {
	ctr, err := contract.CreateSignatureContract(account.PubKey())
	if err != nil {
		return Uint160{}, fmt.Errorf("CreateSignatureContract error:%s", err)
	}
	return ctr.ProgramHash, nil
}

func (this *Ontology) GetAccountsProgramHash(owner *account.Account, m int, accounts []*account.Account) (Uint160, error) {
	if m > len(accounts) {
		return Uint160{}, fmt.Errorf("m:%v should not larger then count of accounts:%v", m, len(accounts))
	}
	pubKeys := make([]*crypto.PubKey, 0, len(accounts))
	for _, ac := range accounts {
		pubKeys = append(pubKeys, ac.PubKey())
	}
	ctr, err := contract.CreateMultiSigContract(owner.ProgramHash, m, pubKeys)
	if err != nil {
		return Uint160{}, fmt.Errorf("CreateMultiSigContract error:%s", err)
	}
	return ctr.ProgramHash, nil
}

func (this *Ontology) getQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}

func (this *Ontology) getRpcAddress() string {
	if len(this.rpcAddresses) == 0 {
		return ""
	}
	return this.rpcAddresses[0]
}

func (this *Ontology) getWSAddress() string {
	if len(this.wsAddresses) == 0 {
		return ""
	}
	return this.wsAddresses[rand.Intn(len(this.wsAddresses))]
}

func (this *Ontology) GetTransactionReference(tx *transaction.Transaction) (map[*utxo.UTXOTxInput]*utxo.TxOutput, error) {
	if tx.TxType == transaction.RegisterAsset {
		return nil, nil
	}
	//UTXO input /  Outputs
	reference := make(map[*utxo.UTXOTxInput]*utxo.TxOutput)
	// Key index，v UTXOInput
	for _, item := range tx.UTXOInputs {
		referTx, err := this.GetTransaction(item.ReferTxID)
		if err != nil {
			return nil, fmt.Errorf("GetTransaction refer txHash:%x", item.ReferTxID)
		}
		index := item.ReferTxOutputIndex
		reference[item] = referTx.Outputs[index]
	}
	return reference, nil
}

func (this *Ontology) sendRpcRequest(method string, params []interface{}) ([]byte, error) {
	data, err := this.Call(this.getRpcAddress(), method, this.getQid(), params)
	if method == ONT_RPC_SENDTRANSACTION {
		//log4.Debug("Call:%s params:%+v", method, params)
		log4.Debug("Res:%s", data)
	}
	if err != nil {
		return nil, fmt.Errorf("Call %s error:%s", method, err)
	}
	if err != nil {
		return nil, fmt.Errorf("Call %s error:%s", method, err)
	}
	if data == nil {
		return nil, fmt.Errorf("Call %s return nil.", method)
	}
	res := &ONTJsonRpcRes{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal ONTJsonRpcRes:%s error:%s", res, err)
	}
	data, err = res.HandleResult()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Call sends RPC request to server
func (this *Ontology) Call(address string, method string, id interface{}, params []interface{}) ([]byte, error) {
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Marshal JSON request: %v\n", err)
		return nil, err
	}
	resp, err := this.client.Post(address, "application/json", strings.NewReader(string(data)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "POST request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GET response: %v\n", err)
		return nil, err
	}

	return body, nil
}

func (this *Ontology) NewDeployCodeTransaction(fc *code.FunctionCode,
	programHash Uint160,
	name, codeversion, author, email, desp string,
	vmType types.VmType) (*transaction.Transaction, error) {
	tx, err := transaction.NewDeployTransaction(fc, programHash, name, codeversion, author, email, desp, vmType, false)
	if err != nil {
		return nil, fmt.Errorf("NewDeployTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) NewInvokeTransaction(fc []byte, codeHash Uint160) (*transaction.Transaction, error) {
	tx, err := transaction.NewInvokeTransaction(fc, codeHash)
	if err != nil {
		return nil, fmt.Errorf("NewInvokeTransaction error:%s", err)
	}
	this.setNonce(tx)
	return tx, nil
}

func (this *Ontology) DeploySmartContract(
	account *account.Account,
	smartContractCode string,
	smartContractParams []contract.ContractParameterType,
	smartContractReturnType contract.ContractParameterType,
	smartContractName,
	smartContractVersion,
	smartContractAuthor,
	smartContractEmail,
	smartContractDesc string,
	smartContractVMType types.VmType) (Uint256, error) {

	c, err := hex.DecodeString(smartContractCode)
	if err != nil {
		return Uint256{}, fmt.Errorf("hex.DecodeString code:%s error:%s", smartContractCode, err)
	}
	//fmt.Println("code:", smartContractCode, c)
	fc := &code.FunctionCode{
		Code:           c,
		ParameterTypes: smartContractParams,
		ReturnType:     smartContractReturnType,
	}
	tx, err := this.NewDeployCodeTransaction(
		fc,
		account.ProgramHash,
		smartContractName,
		smartContractVersion,
		smartContractAuthor,
		smartContractEmail,
		smartContractDesc,
		smartContractVMType,
	)
	if err != nil {
		return Uint256{}, fmt.Errorf("NewDeployCodeTransaction error:%s", err)
	}

	txHash, err := this.SendTransaction(account, tx)
	if err != nil {
		return Uint256{}, fmt.Errorf("SendTransaction tx:%+v error:%s", tx, err)
	}
	return txHash, nil
}

func (this *Ontology) buildSmartContractParamInter(builder *neovm.ParamsBuilder, smartContractParams []interface{}) error {
	//虚拟机参数入栈时会反序
	for i := len(smartContractParams) - 1; i >= 0; i-- {
		switch v := smartContractParams[i].(type) {
		case bool:
			builder.EmitPushBool(v)
		case int:
			builder.EmitPushInteger(big.NewInt(int64(v)))
		case uint:
			builder.EmitPushInteger(big.NewInt(int64(v)))
		case int32:
			builder.EmitPushInteger(big.NewInt(int64(v)))
		case uint32:
			builder.EmitPushInteger(big.NewInt(int64(v)))
		case int64:
			builder.EmitPushInteger(big.NewInt(int64(v)))
		case Fixed64:
			builder.EmitPushInteger(big.NewInt(int64(v.GetData())))
		case uint64:
			val := big.NewInt(0)
			builder.EmitPushInteger(val.SetUint64(uint64(v)))
		case string:
			builder.EmitPushByteArray([]byte(v))
		case *big.Int:
			builder.EmitPushInteger(v)
		case []byte:
			builder.EmitPushByteArray(v)
		case []interface{}:
			err := this.buildSmartContractParamInter(builder, v)
			if err != nil {
				return err
			}
			builder.EmitPushInteger(big.NewInt(int64(len(v))))
			builder.Emit(neovm.PACK)
		default:
			return fmt.Errorf("unsupported param:%s", v)
		}
	}
	return nil
}

func (this *Ontology) BuildSmartContractParam(smartContractParams []interface{}) ([]byte, error) {
	builder := neovm.NewParamsBuilder(new(bytes.Buffer))
	err := this.buildSmartContractParamInter(builder, smartContractParams)
	if err != nil {
		return nil, err
	}
	return builder.ToArray(), nil
}

func (this *Ontology) InvokeSmartContract(
	account *account.Account,
	smartContractCode string,
	smartContractParams []interface{}) (interface{}, error) {
	tx, err := this.BuildSmartContractInvokerTx(smartContractCode, smartContractParams)
	if err != nil {
		return nil, fmt.Errorf("buildSmartContractInvokerTx error:%s", err)
	}

	return this.InvokeSmartContractWithTx(account, tx)
}

func (this *Ontology) BuildSmartContractInvokerTx(
	smartContractCode string,
	smartContractParams []interface{}) (*transaction.Transaction, error) {
	c, err := hex.DecodeString(smartContractCode)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString code:%s error:%s", smartContractCode, err)
	}
	codeHash, err := ToCodeHash(c)
	if err != nil {
		return nil, fmt.Errorf("ToCodeHash Code:%x error:%s", c, err)
	}

	param, err := this.BuildSmartContractParam(smartContractParams)
	if err != nil {
		return nil, err
	}

	tx, err := this.NewInvokeTransaction(param, codeHash)
	if err != nil {
		return nil, fmt.Errorf("NewInvokeTransaction error:%s", err)
	}
	return tx, nil
}

func (this *Ontology) InvokeSmartContractWithTx(account *account.Account, tx *transaction.Transaction) (interface{}, error) {
	wsClient := utils.NewWebSocketClient(this.getWSAddress())
	recvCh, err := wsClient.Connet()
	if err != nil {
		return nil, fmt.Errorf("NewWebSocketClient error:%s", err)
	}
	defer wsClient.Close()

	err = this.WSSendTransaction(wsClient, account, tx)
	if err != nil {
		return nil, fmt.Errorf("WSSendTransaction error:%s", err)
	}

	timeout := 30 * time.Second
	timer := time.NewTimer(timeout)
	for {
		select {
		case <-timer.C:
			return nil, fmt.Errorf("WaitSmartContractRes Timeout after:%v secs.", timeout.Seconds())
		case data := <-recvCh:
			if data == nil {
				return nil, fmt.Errorf("SmartContractResp is nil")
			}
			resp := make(map[string]interface{}, 0)
			err := json.Unmarshal(data, &resp)
			if err != nil {
				return nil, fmt.Errorf("SmartContractResp json.Unmarshal:%s error:%s", data, err)
			}
			log4.Info("==>WS:%s", data)
			action := resp["Action"]
			if action == ONT_HEARTBEAT {
				continue
			}
			if action == ONT_SMARTCONTRACTINVOKE {
				timer.Stop()
				scErr, ok := resp["Error"].(float64)
				if !ok {
					return nil, fmt.Errorf("SmartContract Error:%v assert to float64 failed", resp["Error"])
				}
				if int(scErr) != DNA_ERR_OK {
					return nil, fmt.Errorf("InvokeSmartContract failed. Error:%v Desc:%s Res:%+v", resp["Error"], resp["Desc"], resp["Result"])
				}
				res := resp["Result"]
				if res == nil {
					return nil, fmt.Errorf("InvokeSmartContract return nil")
				}
				return res, nil
			}
		}
	}
	return nil, nil
}

func (this *Ontology) WSSendTransaction(ws *utils.WebSocketClient, signer *account.Account, tx *transaction.Transaction) error {
	attr := &transaction.TxAttribute{Usage: transaction.Script, Data: signer.ProgramHash.ToArray()}
	tx.Attributes = append(tx.Attributes, attr)
	err := this.SignTransaction(tx, []*account.Account{signer})
	if err != nil {
		return fmt.Errorf("SignTransaction error:%s", err)
	}

	var buffer bytes.Buffer
	err = tx.Serialize(&buffer)
	if err != nil {
		return fmt.Errorf("Serialize error:%s", err)
	}

	txData := hex.EncodeToString(buffer.Bytes())

	req := map[string]interface{}{
		"Action": ONT_SENDTRANSACTION,
		"Data":   txData,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal Req:%+v error:%s", req, err)
	}
	return ws.Send(data)
}

func (this *Ontology) WaitSmartContractRes(exitCh chan interface{}, timeout ...time.Duration) error {
	var t time.Duration
	if len(timeout) == 0 {
		t = time.Second * 12
	} else {
		t = timeout[0]
	}
	timer := time.NewTimer(t)
	select {
	case <-timer.C:
		return fmt.Errorf("Timeout after %vsecs.", t.Seconds())
	case <-exitCh:
		timer.Stop()
	}
	return nil
}
