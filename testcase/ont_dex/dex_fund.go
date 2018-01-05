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
	"time"
)

func init() {
	fmt.Printf("-------> DexFund CodeHash:%x Reverse:%x\n", DexFund.CodeHash().ToArray(), DexFund.CodeHash().ToArrayReverse())
}

var DexFundCode = "012ac56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e52756e74696d652e47657454726967676572009c6c766b53527ac46c766b53c3641d00616c766b52c30002e803c46c766b52c36c766b54527ac4628b056168164e656f2e52756e74696d652e47657454726967676572609c6c766b55527ac46c766b55c3644805616c766b00c304696e6974876c766b56527ac46c766b56c3648000616c766b51c3c0539c009c6c766b5a527ac46c766b5ac3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46212056c766b51c300c36c766b57527ac46c766b51c351c36c766b58527ac46c766b51c352c36c766b59527ac46c766b57c36c766b58c36c766b59c361527265dc046c766b54527ac462c9046c766b00c30b6368616e676561646d696e876c766b5b527ac46c766b5bc3645800616c766b51c3c0519c009c6c766b5d527ac46c766b5dc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46274046c766b51c300c36c766b5c527ac46c766b5cc36165ec056c766b54527ac46253046c766b00c30867657461646d696e876c766b5e527ac46c766b5ec3641200616165e9066c766b54527ac46226046c766b00c30973657463616c6c6572876c766b5f527ac46c766b5fc3645a00616c766b51c3c0519c009c6c766b0111527ac46c766b0111c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462d1036c766b51c300c36c766b60527ac46c766b60c36165dd066c766b54527ac462b0036c766b00c30967657463616c6c6572876c766b0112527ac46c766b0112c3641200616165ca076c766b54527ac46280036c766b00c3076465706f736974876c766b0113527ac46c766b0113c364120061616507086c766b54527ac46252036c766b00c3046c6f636b876c766b0114527ac46c766b0114c3647200616c766b51c3c0529c009c6c766b0117527ac46c766b0117c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46200036c766b51c300c36c766b0115527ac46c766b51c351c36c766b0116527ac46c766b0115c36c766b0116c3617c65c00f6c766b54527ac462c7026c766b00c306756e6c6f636b876c766b0118527ac46c766b0118c3647200616c766b51c3c0529c009c6c766b011b527ac46c766b011bc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46273026c766b51c300c36c766b0119527ac46c766b51c351c36c766b011a527ac46c766b0119c36c766b011ac3617c659e106c766b54527ac4623a026c766b00c30962616c616e63656f66876c766b011c527ac46c766b011cc3645c00616c766b51c3c0519c009c6c766b011e527ac46c766b011ec3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462e3016c766b51c300c36c766b011d527ac46c766b011dc36165e4116c766b54527ac462c0016c766b00c30772656365697074876c766b011f527ac46c766b011fc3647200616c766b51c3c0529c009c6c766b0122527ac46c766b0122c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac4626b016c766b51c300c36c766b0120527ac46c766b51c351c36c766b0121527ac46c766b0120c36c766b0121c3617c656f0a6c766b54527ac46232016c766b00c3077061796d656e74876c766b0123527ac46c766b0123c3647200616c766b51c3c0529c009c6c766b0126527ac46c766b0126c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462dd006c766b51c300c36c766b0124527ac46c766b51c351c36c766b0125527ac46c766b0124c36c766b0125c3617c65a50b6c766b54527ac462a4006c766b00c316636865636b63616c6c65727065726d69737373696f6e876c766b0127527ac46c766b0127c3645c00616c766b51c3c0519c009c6c766b0129527ac46c766b0129c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46240006c766b51c300c36c766b0128527ac46c766b0128c361655e116c766b54527ac4621d00616c766b52c30002ea03c46c766b52c36c766b54527ac46203006c766b54c3616c756659c56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c000a06c766b56527ac46c766b56c3641d00616c766b53c30002f203c46c766b53c36c766b57527ac462e80001006c766b52c36c766b55c39c009c6c766b58527ac46c766b58c3644300616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b52c3615272680f4e656f2e53746f726167652e50757461616168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b51c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e7465787407617373657469646c766b00c3615272680f4e656f2e53746f726167652e507574616c766b53c36c766b57527ac46203006c766b57c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4629a006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3641d00616c766b51c30002f003c46c766b51c36c766b54527ac4624f006168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4628d006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3640f00616c766b51c30002f003c4616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c75660115c56b615fc56c766b00527ac46c766b00c30000c46168164e656f2e53746f726167652e476574436f6e746578740761737365746964617c680f4e656f2e53746f726167652e4765746c766b51527ac46c766b51c3c0009c6c766b5d527ac46c766b5dc3641d00616c766b00c30002f303c46c766b00c36c766b5e527ac462000461682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b52527ac46c766b52c361681d4e656f2e5472616e73616374696f6e2e4765745265666572656e63657300c36c766b53527ac46c766b53c36168154e656f2e4f75747075742e476574417373657449646c766b51c39c009c6c766b5f527ac46c766b5fc3641d00616c766b00c30002ef03c46c766b00c36c766b5e527ac46251036c766b53c36168184e656f2e4f75747075742e476574536372697074486173686c766b54527ac461682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b55527ac46c766b52c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b56527ac4006c766b57527ac46c766b00c3536c766b55c3c4546c766b58527ac4616c766b56c36c766b60527ac4006c766b0111527ac46206016c766b60c36c766b0111c3c36c766b0112527ac4616c766b00c36c766b58c36c766b0112c36168184e656f2e4f75747075742e47657453637269707448617368c46c766b58c351936c766b58527ac46c766b00c36c766b58c36c766b55c36c766b0112c36168184e656f2e4f75747075742e4765745363726970744861736887c46c766b58c351936c766b58527ac46c766b55c36c766b0112c36168184e656f2e4f75747075742e47657453637269707448617368876c766b0113527ac46c766b0113c3642e00616c766b57c36c766b0112c36168134e656f2e4f75747075742e47657456616c7565936c766b57527ac461616c766b0111c351936c766b0111527ac46c766b0111c36c766b60c3c09f63f0fe6c766b57c300a16c766b0114527ac46c766b0114c3641d00616c766b00c30002f103c46c766b00c36c766b5e527ac462500105746f74616c6c766b54c3617c65cc0a6c766b59527ac405617661696c6c766b54c3617c65b50a6c766b5a527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b59c3617c680f4e656f2e53746f726167652e4765746c766b5b527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b5ac3617c680f4e656f2e53746f726167652e4765746c766b5c527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b59c36c766b5bc36c766b57c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b5ac36c766b5cc36c766b57c393615272680f4e656f2e53746f726167652e507574616c766b00c3516c766b5cc36c766b57c393c46c766b00c3526c766b5bc36c766b57c393c46c766b00c36c766b5e527ac46203006c766b5ec3616c75665bc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac4626b01616563086c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac4622c0105746f74616c6c766b00c3617c65e4086c766b54527ac405617661696c6c766b00c3617c65cd086c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b56c36c766b51c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b57c36c766b51c393615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c75665cc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac4629f0161659f066c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac462600105617661696c6c766b00c3617c6520076c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac46c766b55c36c766b51c39f6c766b5b527ac46c766b5bc3641d00616c766b52c30002ed03c46c766b52c36c766b59527ac462dd006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b55c36c766b51c394615272680f4e656f2e53746f726167652e5075746105746f74616c6c766b00c3617c655f066c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c36c766b57c36c766b51c394615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c75665ac56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b56527ac46c766b56c3641d00616c766b52c30002e903c46c766b52c36c766b57527ac46212016165a7046c766b53527ac46c766b53c3009c009c6c766b58527ac46c766b58c3641f00616c766b52c3006c766b53c3c46c766b52c36c766b57527ac462d30005617661696c6c766b00c3617c6528056c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac46c766b55c36c766b51c39f6c766b59527ac46c766b59c3641d00616c766b52c30002ed03c46c766b52c36c766b57527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b55c36c766b51c394615272680f4e656f2e53746f726167652e507574616c766b52c36c766b57527ac46203006c766b57c3616c75665cc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac462670161653c036c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac462280105746f74616c6c766b00c3617c65bd036c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac405617661696c6c766b00c3617c656e036c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46c766b55c36c766b57c3946c766b51c39f6c766b5b527ac46c766b5bc3641d00616c766b52c30002ee03c46c766b52c36c766b59527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b56c36c766b57c36c766b51c393615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c756656c56b6c766b00527ac46153c56c766b51527ac46c766b51c30000c46c766b00c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b54527ac46c766b54c3641d00616c766b51c30002f003c46c766b51c36c766b55527ac462b00005617661696c6c766b00c3617c6528026c766b52527ac46c766b51c3516168164e656f2e53746f726167652e476574436f6e746578746c766b52c3617c680f4e656f2e53746f726167652e476574c405746f74616c6c766b00c3617c65d9016c766b53527ac46c766b51c3526168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e476574c46c766b51c36c766b55527ac46203006c766b55c3616c756655c56b6c766b00527ac46152c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641b00616c766b51c35151c46c766b51c36c766b54527ac46224006c766b51c3516c766b52c36c766b00c39cc46c766b51c36c766b54527ac46203006c766b54c3616c756654c56b616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c3c0009c6c766b51527ac46c766b51c3640f0061006c766b52527ac462610061682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b00c39c009c6c766b53527ac46c766b53c36411006102ec036c766b52527ac4620e00006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c37e6c766b52527ac46203006c766b52c3616c7566"
var DexFund = NewDexFund()

type DexFundContract struct{}

func NewDexFund() *DexFundContract {
	return &DexFundContract{}
}

func (this *DexFundContract) CodeHash() *common.Uint160 {
	c, _ := common.HexToBytes(DexFundCode)
	hashCode, _ := common.ToCodeHash(c)
	return &hashCode
}

func (this *DexFundContract) Deploy(ctx *TestFrameworkContext, admin *account.Account) error {
	_, err := ctx.Ont.DeploySmartContract(admin,
		DexFundCode,
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
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return nil
}

func (this *DexFundContract) Init(ctx *TestFrameworkContext, assetId []byte, admin *account.Account, caller []byte) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexFundCode,
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
		DexFundCode,
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
		DexFundCode,
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
		DexFundCode,
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
		DexFundCode,
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
		DexFundCode,
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
		DexFundCode,
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

//
//func deployDexFund(ctx *TestFrameworkContext) bool {
//	code := DExFundCode
//	c, _ := common.HexToBytes(code)
//	codeHash, err := common.ToCodeHash(c)
//	if err != nil {
//		ctx.LogError("TestDexFund ToCodeHash error:%s", err)
//		return false
//	}
//	ctx.LogInfo("TestDexFund CodeHash: %x , RCodeHash: %x", codeHash, codeHash.ToArrayReverse())
//	_, err = ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
//		code,
//		[]contract.ContractParameterType{contract.String, contract.Array},
//		contract.ContractParameterType(contract.Array),
//		"TestDexFund",
//		"1.0",
//		"",
//		"",
//		"",
//		types.NEOVM,
//	)
//	if err != nil {
//		ctx.LogError("deployDexFund DeploySmartContract error:%s", err)
//		return false
//	}
//	//等待出块
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("deployDexFund WaitForGenerateBlock error:%s", err)
//		return false
//	}
//	//assetId := []byte("")
//	//admin := ctx.OntClient.Admin
//	//if !testDexFundInit(ctx, assetId, admin) {
//	//	return false
//	//}
//	//buyer := ctx.OntClient.Account1
//	//amount := 10
//	//if !testReceipt(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testLock(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testUnLock(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testPayment(ctx, buyer, amount) {
//	//	return false
//	//}
//	return true
//}
//
//func dexFundInit(ctx *TestFrameworkContext, assetId []byte, admin *account.Account) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		admin,
//		DExFundCode,
//		[]interface{}{"init", []interface{}{assetId, admin.ProgramHash.ToArray(), []byte("")}},
//	)
//	if err != nil {
//		ctx.LogError("testDexFundInit error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testDexFundInit res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testDexFundInit getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testDexFundInit failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func fundDeposit(ctx *TestFrameworkContext, assetId common.Uint256, buyer *account.Account, amount int) bool {
//	unspents, err := ctx.Ont.GetUnspendOutput(assetId, buyer.ProgramHash)
//	if err != nil {
//		ctx.LogError("fundDeposit GetUnspendOutput error:%s", err)
//		return false
//	}
//	if unspents == nil {
//		ctx.LogError("fundDeposit GetUnspendOutput return nil")
//		return false
//	}
//	assAmount := ctx.Ont.MakeAssetAmount(float64(amount))
//	txInputs := make([]*utxo.UTXOTxInput, 0, 1)
//	txOutputs := make([]*utxo.TxOutput, 0, 2)
//
//	c, _ := common.HexToBytes(DExFundCode)
//	codeHash, _ := common.ToCodeHash(c)
//	for _, unspent := range unspents {
//		if unspent.Value < assAmount {
//			continue
//		}
//		input := &utxo.UTXOTxInput{
//			ReferTxID:          unspent.Txid,
//			ReferTxOutputIndex: uint16(unspent.Index),
//		}
//		txInputs = append(txInputs, input)
//		output := &utxo.TxOutput{
//			AssetID:     assetId,
//			Value:       assAmount,
//			ProgramHash: codeHash,
//		}
//		txOutputs = append(txOutputs, output)
//		//dibs output
//		dibs := unspent.Value - assAmount
//		if dibs > 0 {
//			output2 := &utxo.TxOutput{
//				AssetID:     output.AssetID,
//				Value:       dibs,
//				ProgramHash: buyer.ProgramHash,
//			}
//			txOutputs = append(txOutputs, output2)
//		}
//		break
//	}
//	if len(txInputs) == 0 {
//		ctx.LogError("fundDeposit TxInput is nil")
//		return false
//	}
//
//	tx, err := ctx.Ont.BuildSmartContractInvokerTx(
//		DExFundCode,
//		[]interface{}{"deposit", []interface{}{}},
//	)
//	if err != nil {
//		ctx.LogError("fundDeposit BuildSmartContractInvokerTx error:%s", err)
//		return false
//	}
//
//	tx.UTXOInputs = txInputs
//	tx.Outputs = txOutputs
//
//	res, err := ctx.Ont.InvokeSmartContractWithTx(buyer, tx)
//	if err != nil {
//		ctx.LogError("fundDeposit InvokeSmartContractWithTx error:%s", err)
//		return false
//	}
//	ctx.LogInfo("fundDeposit res:%s", res)
//
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("fundDeposit getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("fundDeposit failed errorCode:%d", errorCode)
//		return false
//	}
//
//	balance, err := fundAvailBalanceOf(ctx, buyer)
//	if err != nil {
//		ctx.LogError("fundAvailBalanceOf error:%s", err)
//		return false
//	}
//	ctx.LogInfo("fundDeposit fundAvailBalance:%v", balance)
//
//	if common.Fixed64(balance) != assAmount {
//		ctx.LogError("fundDeposit fundAvailBalanc: %v != %v", balance, assAmount)
//		return false
//	}
//	return true
//}
//
//func fundAvailBalanceOf(ctx *TestFrameworkContext, buyer *account.Account) (int, error) {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"availbalanceof", []interface{}{buyer.ProgramHash.ToArray()}},
//	)
//	if err != nil {
//		return 0, fmt.Errorf("fundReceipt error:%s", err)
//	}
//	ctx.LogInfo("fundReceipt res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		return 0, fmt.Errorf("fundReceipt getErrorCode error:%s", err)
//	}
//	if errorCode != 0 {
//		return 0, fmt.Errorf("fundReceipt failed errorCode:%d", errorCode)
//	}
//	v, err := GetRetValue(res, 1, reflect.Int)
//	if err != nil {
//		return 0, fmt.Errorf("fundReceipt GetRetValue error:%s", err)
//	}
//	return v.(int), nil
//}
//
//func fundReceipt(ctx *TestFrameworkContext, receiver *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		receiver,
//		DExFundCode,
//		[]interface{}{"receipt", []interface{}{receiver.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("fundReceipt error:%s", err)
//		return false
//	}
//	ctx.LogInfo("fundReceipt res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("fundReceipt getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("fundReceipt failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testPayment(ctx *TestFrameworkContext, buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"payment", []interface{}{buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testPayment error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testPayment res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testPayment getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testPayment failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testLock(ctx *TestFrameworkContext, buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"lock", []interface{}{buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testLock error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testLock res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testLock getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testLock failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testUnLock(ctx *TestFrameworkContext, buyer *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExFundCode,
//		[]interface{}{"unlock", []interface{}{buyer.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testUnLock error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testUnLock res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testUnLock getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testUnLock failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
