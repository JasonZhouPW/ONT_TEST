package ont_dex

import (
	"fmt"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/core/transaction/utxo"
	"github.com/Ontology/smartcontract/types"
	"reflect"
	//"time"
)

func init() {
	fmt.Printf("-------> DexFund CodeHash:%x Reverse:%x\n", DexFund.CodeHash().ToArray(), DexFund.CodeHash().ToArrayReverse())
}

var DexFundCode = "012ec56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e52756e74696d652e47657454726967676572009c6c766b53527ac46c766b53c3641d00616c766b52c30002e803c46c766b52c36c766b54527ac4627f066168164e656f2e52756e74696d652e47657454726967676572609c6c766b55527ac46c766b55c3643c06616c766b00c304696e6974876c766b56527ac46c766b56c3648000616c766b51c3c0539c009c6c766b5a527ac46c766b5ac3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46206066c766b51c300c36c766b57527ac46c766b51c351c36c766b58527ac46c766b51c352c36c766b59527ac46c766b57c36c766b58c36c766b59c361527265d0056c766b54527ac462bd056c766b00c30b6368616e676561646d696e876c766b5b527ac46c766b5bc3645800616c766b51c3c0519c009c6c766b5d527ac46c766b5dc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46268056c766b51c300c36c766b5c527ac46c766b5cc36165e0066c766b54527ac46247056c766b00c30867657461646d696e876c766b5e527ac46c766b5ec3641200616165dd076c766b54527ac4621a056c766b00c30973657463616c6c6572876c766b5f527ac46c766b5fc3645a00616c766b51c3c0519c009c6c766b0111527ac46c766b0111c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462c5046c766b51c300c36c766b60527ac46c766b60c36165d1076c766b54527ac462a4046c766b00c30967657463616c6c6572876c766b0112527ac46c766b0112c3641200616165be086c766b54527ac46274046c766b00c3076465706f736974876c766b0113527ac46c766b0113c3641200616165fb086c766b54527ac46246046c766b00c3046c6f636b876c766b0114527ac46c766b0114c364af00616c766b51c3c0529c009c6c766b0118527ac46c766b0118c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462f40361682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b0115527ac46c766b51c300c36c766b0116527ac46c766b51c351c36c766b0117527ac46c766b0115c36c766b0116c36c766b0117c361527265fc0f6c766b54527ac4627e036c766b00c306756e6c6f636b876c766b0119527ac46c766b0119c364af00616c766b51c3c0529c009c6c766b011d527ac46c766b011dc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac4622a0361682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b011a527ac46c766b51c300c36c766b011b527ac46c766b51c351c36c766b011c527ac46c766b011ac36c766b011bc36c766b011cc361527265a9106c766b54527ac462b4026c766b00c30962616c616e63656f66876c766b011e527ac46c766b011ec3645c00616c766b51c3c0519c009c6c766b0120527ac46c766b0120c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac4625d026c766b51c300c36c766b011f527ac46c766b011fc36165fb116c766b54527ac4623a026c766b00c30772656365697074876c766b0121527ac46c766b0121c364af00616c766b51c3c0529c009c6c766b0125527ac46c766b0125c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462e50161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b0122527ac46c766b51c300c36c766b0123527ac46c766b51c351c36c766b0124527ac46c766b0122c36c766b0123c36c766b0124c361527265190a6c766b54527ac4626f016c766b00c3077061796d656e74876c766b0126527ac46c766b0126c364af00616c766b51c3c0529c009c6c766b012a527ac46c766b012ac3641d00616c766b52c30002e903c46c766b52c36c766b54527ac4621a0161682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b0127527ac46c766b51c300c36c766b0128527ac46c766b51c351c36c766b0129527ac46c766b0127c36c766b0128c36c766b0129c3615272651e0b6c766b54527ac462a4006c766b00c316636865636b63616c6c65727065726d69737373696f6e876c766b012b527ac46c766b012bc3645c00616c766b51c3c0519c009c6c766b012d527ac46c766b012dc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46240006c766b51c300c36c766b012c527ac46c766b012cc36165fb106c766b54527ac4621d00616c766b52c30002ea03c46c766b52c36c766b54527ac46203006c766b54c3616c756659c56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c000a06c766b56527ac46c766b56c3641d00616c766b53c30002f203c46c766b53c36c766b57527ac462e80001006c766b52c36c766b55c39c009c6c766b58527ac46c766b58c3644300616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b52c3615272680f4e656f2e53746f726167652e50757461616168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b51c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e7465787407617373657469646c766b00c3615272680f4e656f2e53746f726167652e507574616c766b53c36c766b57527ac46203006c766b57c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4629a006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3641d00616c766b51c30002f003c46c766b51c36c766b54527ac4624f006168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4628d006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3640f00616c766b51c30002f003c4616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c75660114c56b6153c56c766b00527ac46c766b00c30000c46168164e656f2e53746f726167652e476574436f6e746578740761737365746964617c680f4e656f2e53746f726167652e4765746c766b51527ac46c766b51c3c0009c6c766b5c527ac46c766b5cc3641d00616c766b00c30002f303c46c766b00c36c766b5d527ac4626d0361682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b52527ac46c766b52c361681d4e656f2e5472616e73616374696f6e2e4765745265666572656e63657300c36c766b53527ac46c766b53c36168154e656f2e4f75747075742e476574417373657449646c766b51c39c009c6c766b5e527ac46c766b5ec3641d00616c766b00c30002ef03c46c766b00c36c766b5d527ac462be026c766b53c36168184e656f2e4f75747075742e476574536372697074486173686c766b54527ac461682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b55527ac46c766b52c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b56527ac4006c766b57527ac4616c766b56c36c766b5f527ac4006c766b60527ac46289006c766b5fc36c766b60c3c36c766b0111527ac4616c766b55c36c766b0111c36168184e656f2e4f75747075742e47657453637269707448617368876c766b0112527ac46c766b0112c3642e00616c766b57c36c766b0111c36168134e656f2e4f75747075742e47657456616c7565936c766b57527ac461616c766b60c351936c766b60527ac46c766b60c36c766b5fc3c09f636eff6c766b57c300a16c766b0113527ac46c766b0113c3641d00616c766b00c30002f103c46c766b00c36c766b5d527ac462500105746f74616c6c766b54c3617c65da0a6c766b58527ac405617661696c6c766b54c3617c65c30a6c766b59527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b58c3617c680f4e656f2e53746f726167652e4765746c766b5a527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b59c3617c680f4e656f2e53746f726167652e4765746c766b5b527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b58c36c766b5ac36c766b57c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b59c36c766b5bc36c766b57c393615272680f4e656f2e53746f726167652e507574616c766b00c3516c766b5bc36c766b57c393c46c766b00c3526c766b5ac36c766b57c393c46c766b00c36c766b5d527ac46203006c766b5dc3616c75665cc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46c766b52c300a16c766b59527ac46c766b59c3641d00616c766b53c30002e903c46c766b53c36c766b5a527ac46270016c766b00c3616587086c766b54527ac46c766b54c3009c009c6c766b5b527ac46c766b5bc3641f00616c766b53c3006c766b54c3c46c766b53c36c766b5a527ac4622c0105746f74616c6c766b51c3617c65e6086c766b55527ac405617661696c6c766b51c3617c65cf086c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c3617c680f4e656f2e53746f726167652e4765746c766b58527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b57c36c766b52c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b56c36c766b58c36c766b52c393615272680f4e656f2e53746f726167652e507574616c766b53c36c766b5a527ac46203006c766b5ac3616c75665dc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46c766b52c300a16c766b59527ac46c766b59c3641d00616c766b53c30002e903c46c766b53c36c766b5a527ac462a4016c766b00c36165b7066c766b54527ac46c766b54c3009c009c6c766b5b527ac46c766b5bc3641f00616c766b53c3006c766b54c3c46c766b53c36c766b5a527ac462600105617661696c6c766b51c3617c6516076c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46c766b56c36c766b52c39f6c766b5c527ac46c766b5cc3641d00616c766b53c30002ed03c46c766b53c36c766b5a527ac462dd006168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b56c36c766b52c394615272680f4e656f2e53746f726167652e5075746105746f74616c6c766b51c3617c6555066c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c680f4e656f2e53746f726167652e4765746c766b58527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c36c766b58c36c766b52c394615272680f4e656f2e53746f726167652e507574616c766b53c36c766b5a527ac46203006c766b5ac3616c75665bc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46c766b52c300a16c766b57527ac46c766b57c3641d00616c766b53c30002e903c46c766b53c36c766b58527ac46217016c766b00c36165b3046c766b54527ac46c766b54c3009c009c6c766b59527ac46c766b59c3641f00616c766b53c3006c766b54c3c46c766b53c36c766b58527ac462d30005617661696c6c766b51c3617c6512056c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46c766b56c36c766b52c39f6c766b5a527ac46c766b5ac3641d00616c766b53c30002ed03c46c766b53c36c766b58527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b56c36c766b52c394615272680f4e656f2e53746f726167652e507574616c766b53c36c766b58527ac46203006c766b58c3616c75665dc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46c766b52c300a16c766b59527ac46c766b59c3641d00616c766b53c30002e903c46c766b53c36c766b5a527ac4626c016c766b00c361653c036c766b54527ac46c766b54c3009c009c6c766b5b527ac46c766b5bc3641f00616c766b53c3006c766b54c3c46c766b53c36c766b5a527ac462280105746f74616c6c766b51c3617c659b036c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac405617661696c6c766b51c3617c654c036c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c680f4e656f2e53746f726167652e4765746c766b58527ac46c766b56c36c766b58c3946c766b52c39f6c766b5c527ac46c766b5cc3641d00616c766b53c30002ee03c46c766b53c36c766b5a527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b57c36c766b58c36c766b52c393615272680f4e656f2e53746f726167652e507574616c766b53c36c766b5a527ac46203006c766b5ac3616c756656c56b6c766b00527ac46153c56c766b51527ac46c766b51c30000c46c766b00c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b54527ac46c766b54c3641d00616c766b51c30002f003c46c766b51c36c766b55527ac462b00005617661696c6c766b00c3617c6506026c766b52527ac46c766b51c3516168164e656f2e53746f726167652e476574436f6e746578746c766b52c3617c680f4e656f2e53746f726167652e476574c405746f74616c6c766b00c3617c65b7016c766b53527ac46c766b51c3526168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e476574c46c766b51c36c766b55527ac46203006c766b55c3616c756655c56b6c766b00527ac46152c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641b00616c766b51c35151c46c766b51c36c766b54527ac46224006c766b51c3516c766b52c36c766b00c39cc46c766b51c36c766b54527ac46203006c766b54c3616c756655c56b6c766b00527ac4616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b51527ac46c766b51c3c0009c6c766b52527ac46c766b52c3640f0061006c766b53527ac46238006c766b00c36c766b51c39c009c6c766b54527ac46c766b54c36411006102ec036c766b53527ac4620e00006c766b53527ac46203006c766b53c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c37e6c766b52527ac46203006c766b52c3616c7566"
var DexFund = NewDexFund()

type DexFundContract struct{}

func NewDexFund() *DexFundContract {
	return &DexFundContract{}
}

func (this *DexFundContract) GetCode()string  {
	return DexFundCode
}

func (this *DexFundContract) CodeHash() *common.Uint160 {
	c, _ := common.HexToBytes(this.GetCode())
	hashCode, _ := common.ToCodeHash(c)
	return &hashCode
}

func (this *DexFundContract) Deploy(ctx *TestFrameworkContext, admin *account.Account) error {
	ctx.LogInfo("DexFundContract Deploy")
	_, err := ctx.Ont.DeploySmartContract(admin,
		this.GetCode(),
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"DexFundContract",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		return err
	}
	//等待出块
	//_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	//if err != nil {
	//	return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	//}
	return nil
}

func (this *DexFundContract) Init(ctx *TestFrameworkContext, assetId []byte, admin *account.Account, caller []byte) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		this.GetCode(),
		[]interface{}{"init", []interface{}{assetId, admin.ProgramHash.ToArray(), caller}},
	)
	if err != nil {
		return err
	}
	ctx.LogInfo("DexFundContract init res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%d", errorCode)
	}
	return nil
}

func (this *DexFundContract) Deposit(ctx *TestFrameworkContext, assetId common.Uint256, user *account.Account, amount float64) error {
	unspents, err := ctx.Ont.GetUnspendOutput(assetId, user.ProgramHash)
	if err != nil {
		return fmt.Errorf("GetUnspendOutput error:%s", err)
	}
	if unspents == nil {
		return fmt.Errorf("GetUnspendOutput return nil")
	}

	ctx.LogInfo("Receier:%x %x", this.CodeHash(), user.ProgramHash)

	assAmount := ctx.Ont.MakeAssetAmount(amount)
	txInputs := make([]*utxo.UTXOTxInput, 0, 1)
	txOutputs := make([]*utxo.TxOutput, 0, 2)

	for _, unspent := range unspents {
		if unspent.Value < assAmount {
			continue
		}
		input := &utxo.UTXOTxInput{
			ReferTxID:          unspent.Txid,
			ReferTxOutputIndex: uint16(unspent.Index),
		}
		txInputs = append(txInputs, input)
		output := &utxo.TxOutput{
			AssetID:     assetId,
			Value:       assAmount,
			ProgramHash: *this.CodeHash(),
		}
		txOutputs = append(txOutputs, output)
		//dibs output
		dibs := unspent.Value - assAmount
		if dibs > 0 {
			output2 := &utxo.TxOutput{
				AssetID:     output.AssetID,
				Value:       dibs,
				ProgramHash: user.ProgramHash,
			}
			txOutputs = append(txOutputs, output2)
		}
		break
	}
	if len(txInputs) == 0 {
		return fmt.Errorf("TxInput is nil")
	}
	ctx.LogInfo("deposit amount:%v", assAmount)
	tx, err := ctx.Ont.BuildSmartContractInvokerTx(
		this.GetCode(),
		[]interface{}{"deposit", []interface{}{}},
	)
	if err != nil {
		return fmt.Errorf("BuildSmartContractInvokerTx error:%s", err)
	}

	tx.UTXOInputs = txInputs
	tx.Outputs = txOutputs

	res, err := ctx.Ont.InvokeSmartContractWithTx(user, tx)
	if err != nil {
		return fmt.Errorf("InvokeSmartContractWithTx error:%s", err)
	}
	ctx.LogInfo("fundDeposit res:%s", res)

	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%d", errorCode)
	}
	return nil
}

func (this *DexFundContract) BalanceOf(ctx *TestFrameworkContext, user *account.Account) (avail, total float64, err error) {
	res, err := ctx.Ont.InvokeSmartContract(
		user,
		this.GetCode(),
		[]interface{}{"balanceof", []interface{}{user.ProgramHash.ToArray()}},
	)
	if err != nil {
		return 0, 0, fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("BalanceOf res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return 0, 0, fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return 0, 0, fmt.Errorf("ErrorCode:%d", errorCode)
	}
	va, err := GetRetValue(res, 1, reflect.Int)
	if err != nil {
		return 0, 0, fmt.Errorf("GetRetValue error:%s", err)
	}
	vt, err := GetRetValue(res, 2, reflect.Int)
	if err != nil {
		return 0, 0, fmt.Errorf("GetRetValue error:%s", err)
	}

	avail = ctx.Ont.GetRawAssetAmount(common.Fixed64(va.(int)))
	total = ctx.Ont.GetRawAssetAmount(common.Fixed64(vt.(int)))
	return avail, total, nil
}

func (this *DexFundContract) SetCaller(ctx *TestFrameworkContext, admin *account.Account, caller []byte) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		this.GetCode(),
		[]interface{}{"setcaller", []interface{}{caller}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("setFundCaller res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%d", errorCode)
	}
	return nil
}

func (this *DexFundContract) ChangeAdmin(ctx *TestFrameworkContext, admin, newAdmin *account.Account) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		this.GetCode(),
		[]interface{}{"changeadmin", []interface{}{newAdmin.ProgramHash.ToArray()}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexFundContract ChangeAdmin res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexFundContract) GetAdmin(ctx *TestFrameworkContext) (string, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		this.GetCode(),
		[]interface{}{"getadmin", []interface{}{}},
	)
	if err != nil {
		return "", fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexFundContract GetAdmin res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return "", fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return "", fmt.Errorf("ErrorCode:%v", errorCode)
	}
	admin, err := GetRetValue(res, 1, reflect.String)
	if err != nil {
		return "", fmt.Errorf("GetRetValue error:%s", err)
	}
	return admin.(string), err
}

func (this *DexFundContract) CheckCallerPermission(ctx *TestFrameworkContext, caller []byte) (bool, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		this.GetCode(),
		[]interface{}{"checkcallerpermisssion", []interface{}{caller}},
	)
	if err != nil {
		return false, fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("CheckCallerPermission res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return false, fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return false, fmt.Errorf("ErrorCode:%v", errorCode)
	}
	v, err := GetRetValue(res, 1, reflect.Bool)
	if err != nil {
		return false, fmt.Errorf("GetRetValue error:%s", err)
	}
	return v.(bool), nil
}
