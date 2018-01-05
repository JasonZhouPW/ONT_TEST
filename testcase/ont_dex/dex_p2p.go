package ont_dex

import (
	"fmt"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	//"github.com/Ontology/crypto"
	//"math/rand"
	"reflect"
)

func init() {
	fmt.Printf("-------> DexP2P CodeHash:%x Reverse:%x\n", DexP2P.CodeHash().ToArray(), DexP2P.CodeHash().ToArrayReverse())
}

var DexP2PCode = "012ac56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e52756e74696d652e47657454726967676572009c6c766b53527ac46c766b53c3641d00616c766b52c30002e803c46c766b52c36c766b54527ac4628b056168164e656f2e52756e74696d652e47657454726967676572609c6c766b55527ac46c766b55c3644805616c766b00c304696e6974876c766b56527ac46c766b56c3648000616c766b51c3c0539c009c6c766b5a527ac46c766b5ac3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46212056c766b51c300c36c766b57527ac46c766b51c351c36c766b58527ac46c766b51c352c36c766b59527ac46c766b57c36c766b58c36c766b59c361527265dc046c766b54527ac462c9046c766b00c30b6368616e676561646d696e876c766b5b527ac46c766b5bc3645800616c766b51c3c0519c009c6c766b5d527ac46c766b5dc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46274046c766b51c300c36c766b5c527ac46c766b5cc36165ec056c766b54527ac46253046c766b00c30867657461646d696e876c766b5e527ac46c766b5ec3641200616165e9066c766b54527ac46226046c766b00c30973657463616c6c6572876c766b5f527ac46c766b5fc3645a00616c766b51c3c0519c009c6c766b0111527ac46c766b0111c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462d1036c766b51c300c36c766b60527ac46c766b60c36165dd066c766b54527ac462b0036c766b00c30967657463616c6c6572876c766b0112527ac46c766b0112c3641200616165ca076c766b54527ac46280036c766b00c3076465706f736974876c766b0113527ac46c766b0113c364120061616507086c766b54527ac46252036c766b00c3046c6f636b876c766b0114527ac46c766b0114c3647200616c766b51c3c0529c009c6c766b0117527ac46c766b0117c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46200036c766b51c300c36c766b0115527ac46c766b51c351c36c766b0116527ac46c766b0115c36c766b0116c3617c653b0f6c766b54527ac462c7026c766b00c306756e6c6f636b876c766b0118527ac46c766b0118c3647200616c766b51c3c0529c009c6c766b011b527ac46c766b011bc3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46273026c766b51c300c36c766b0119527ac46c766b51c351c36c766b011a527ac46c766b0119c36c766b011ac3617c6519106c766b54527ac4623a026c766b00c30962616c616e63656f66876c766b011c527ac46c766b011cc3645c00616c766b51c3c0519c009c6c766b011e527ac46c766b011ec3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462e3016c766b51c300c36c766b011d527ac46c766b011dc361655f116c766b54527ac462c0016c766b00c30772656365697074876c766b011f527ac46c766b011fc3647200616c766b51c3c0529c009c6c766b0122527ac46c766b0122c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac4626b016c766b51c300c36c766b0120527ac46c766b51c351c36c766b0121527ac46c766b0120c36c766b0121c3617c65ea096c766b54527ac46232016c766b00c3077061796d656e74876c766b0123527ac46c766b0123c3647200616c766b51c3c0529c009c6c766b0126527ac46c766b0126c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac462dd006c766b51c300c36c766b0124527ac46c766b51c351c36c766b0125527ac46c766b0124c36c766b0125c3617c65200b6c766b54527ac462a4006c766b00c316636865636b63616c6c65727065726d69737373696f6e876c766b0127527ac46c766b0127c3645c00616c766b51c3c0519c009c6c766b0129527ac46c766b0129c3641d00616c766b52c30002e903c46c766b52c36c766b54527ac46240006c766b51c300c36c766b0128527ac46c766b0128c36165d9106c766b54527ac4621d00616c766b52c30002ea03c46c766b52c36c766b54527ac46203006c766b54c3616c756659c56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c000a06c766b56527ac46c766b56c3641d00616c766b53c30002f203c46c766b53c36c766b57527ac462e80001006c766b52c36c766b55c39c009c6c766b58527ac46c766b58c3644300616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b52c3615272680f4e656f2e53746f726167652e50757461616168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b51c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e7465787407617373657469646c766b00c3615272680f4e656f2e53746f726167652e507574616c766b53c36c766b57527ac46203006c766b57c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4629a006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3641d00616c766b51c30002f003c46c766b51c36c766b54527ac4624f006168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740966756e6461646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002f303c46c766b51c36c766b54527ac4628d006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3640f00616c766b51c30002f003c4616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c65726c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c75660114c56b6154c56c766b00527ac46c766b00c30000c46168164e656f2e53746f726167652e476574436f6e746578740761737365746964617c680f4e656f2e53746f726167652e4765746c766b51527ac46c766b51c3c0009c6c766b5c527ac46c766b5cc3641d00616c766b00c30002f303c46c766b00c36c766b5d527ac4627b0361682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b52527ac46c766b52c361681d4e656f2e5472616e73616374696f6e2e4765745265666572656e63657300c36c766b53527ac46c766b53c36168154e656f2e4f75747075742e476574417373657449646c766b51c39c009c6c766b5e527ac46c766b5ec3641d00616c766b00c30002ef03c46c766b00c36c766b5d527ac462cc026c766b53c36168184e656f2e4f75747075742e476574536372697074486173686c766b54527ac461682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b55527ac46c766b52c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b56527ac4006c766b57527ac4616c766b56c36c766b5f527ac4006c766b60527ac46289006c766b5fc36c766b60c3c36c766b0111527ac4616c766b0111c36168184e656f2e4f75747075742e476574536372697074486173686c766b55c39c6c766b0112527ac46c766b0112c3642e00616c766b57c36c766b0111c36168134e656f2e4f75747075742e47657456616c7565936c766b57527ac461616c766b60c351936c766b60527ac46c766b60c36c766b5fc3c09f636eff6c766b57c300a0009c6c766b0113527ac46c766b0113c3641d00616c766b00c30002f103c46c766b00c36c766b5d527ac4625c016c766b00c3536c766b57c3c405746f74616c6c766b54c3617c65cc0a6c766b58527ac405617661696c6c766b54c3617c65b50a6c766b59527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b58c3617c680f4e656f2e53746f726167652e4765746c766b5a527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b59c3617c680f4e656f2e53746f726167652e4765746c766b5b527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b58c36c766b5ac36c766b57c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b59c36c766b5bc36c766b57c393615272680f4e656f2e53746f726167652e507574616c766b00c3516c766b5bc36c766b57c393c46c766b00c3526c766b5ac36c766b57c393c46c766b00c36c766b5d527ac46203006c766b5dc3616c75665bc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac4626b01616563086c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac4622c0105746f74616c6c766b00c3617c65e4086c766b54527ac405617661696c6c766b00c3617c65cd086c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b56c36c766b51c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b57c36c766b51c393615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c75665cc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac4629f0161659f066c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac462600105617661696c6c766b00c3617c6520076c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac46c766b55c36c766b51c39f6c766b5b527ac46c766b5bc3641d00616c766b52c30002ed03c46c766b52c36c766b59527ac462dd006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b55c36c766b51c394615272680f4e656f2e53746f726167652e5075746105746f74616c6c766b00c3617c655f066c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c36c766b57c36c766b51c394615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c75665ac56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b56527ac46c766b56c3641d00616c766b52c30002e903c46c766b52c36c766b57527ac46212016165a7046c766b53527ac46c766b53c3009c009c6c766b58527ac46c766b58c3641f00616c766b52c3006c766b53c3c46c766b52c36c766b57527ac462d30005617661696c6c766b00c3617c6528056c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac46c766b55c36c766b51c39f6c766b59527ac46c766b59c3641d00616c766b52c30002ed03c46c766b52c36c766b57527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b55c36c766b51c394615272680f4e656f2e53746f726167652e507574616c766b52c36c766b57527ac46203006c766b57c3616c75665cc56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c300a16c766b58527ac46c766b58c3641d00616c766b52c30002e903c46c766b52c36c766b59527ac462670161653c036c766b53527ac46c766b53c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b52c3006c766b53c3c46c766b52c36c766b59527ac462280105746f74616c6c766b00c3617c65bd036c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac405617661696c6c766b00c3617c656e036c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b56c3617c680f4e656f2e53746f726167652e4765746c766b57527ac46c766b55c36c766b57c3946c766b51c39f6c766b5b527ac46c766b5bc3641d00616c766b52c30002ee03c46c766b52c36c766b59527ac46250006168164e656f2e53746f726167652e476574436f6e746578746c766b56c36c766b57c36c766b51c393615272680f4e656f2e53746f726167652e507574616c766b52c36c766b59527ac46203006c766b59c3616c756656c56b6c766b00527ac46153c56c766b51527ac46c766b51c30000c46c766b00c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b54527ac46c766b54c3641d00616c766b51c30002f003c46c766b51c36c766b55527ac462b00005617661696c6c766b00c3617c6528026c766b52527ac46c766b51c3516168164e656f2e53746f726167652e476574436f6e746578746c766b52c3617c680f4e656f2e53746f726167652e476574c405746f74616c6c766b00c3617c65d9016c766b53527ac46c766b51c3526168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e476574c46c766b51c36c766b55527ac46203006c766b55c3616c756655c56b6c766b00527ac46152c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641b00616c766b51c35151c46c766b51c36c766b54527ac46224006c766b51c3516c766b52c36c766b00c39cc46c766b51c36c766b54527ac46203006c766b54c3616c756654c56b616168164e656f2e53746f726167652e476574436f6e746578740a66756e6463616c6c6572617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c3c0009c6c766b51527ac46c766b51c3640f0061006c766b52527ac462610061682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b00c39c009c6c766b53527ac46c766b53c36411006102ec036c766b52527ac4620e00006c766b52527ac46203006c766b52c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c37e6c766b52527ac46203006c766b52c3616c7566"
var DexP2P = NewDexP2PContract()

type DexP2PContract struct{}

func NewDexP2PContract() *DexP2PContract {
	return &DexP2PContract{}
}
func (this *DexP2PContract) CodeHash() *common.Uint160 {
	c, _ := common.HexToBytes(DexP2PCode)
	hashCode, _ := common.ToCodeHash(c)
	return &hashCode
}

func (this *DexP2PContract) Deploy(ctx *TestFrameworkContext) error {
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		DexP2PCode,
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"DexP2PContract",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		return fmt.Errorf("DeploySmartContract error:%s", err)
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	}
	return nil
}

func (this *DexP2PContract) Init(ctx *TestFrameworkContext, admin *account.Account, lockTime int) error {
	if lockTime == 0 {
		lockTime = 24*3600
	}
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexP2PCode,
		[]interface{}{"init", []interface{}{admin.ProgramHash.ToArray(), lockTime}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract Init res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexP2PContract) ChangeAdmin(ctx *TestFrameworkContext, admin, newAdmin *account.Account) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexP2PCode,
		[]interface{}{"changeadmin", []interface{}{newAdmin.ProgramHash.ToArray()}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract ChangeAdmin res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexP2PContract) GetAdmin(ctx *TestFrameworkContext) (string, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		DexP2PCode,
		[]interface{}{"getadmin", []interface{}{}},
	)
	if err != nil {
		return "", fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract GetAdmin res:%s", res)
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

func (this *DexP2PContract) MakeBuyOrder(ctx *TestFrameworkContext,
	orderSig, orderId []byte,
	buyer, seller *account.Account,
	amount float64) error {
	buyerPk, err := buyer.PublicKey.EncodePoint(true)
	if err != nil {
		return fmt.Errorf("PublicKey.EncodePoint error:%s", err)
	}
	res, err := ctx.Ont.InvokeSmartContract(
		seller,
		DexP2PCode,
		[]interface{}{"makebuyorder", []interface{}{
			orderSig, orderId,
			buyer.ProgramHash.ToArray(),
			seller.ProgramHash.ToArray(),
			buyerPk, amount}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("makeBuyOrder res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexP2PContract) BuyOrderComplete(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account) error {
	res, err := ctx.Ont.InvokeSmartContract(
		buyer,
		DexP2PCode,
		[]interface{}{"buyordercomplete", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("buyOrderComplete res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexP2PContract)BuyOrderCancel (ctx *TestFrameworkContext, orderId []byte, buyer *account.Account)error{
	res, err := ctx.Ont.InvokeSmartContract(
		buyer,
		DexP2PCode,
		[]interface{}{"buyordercancel", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("BuyOrderCancel res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexP2PContract) GetOrderLockTime(ctx *TestFrameworkContext)(int, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		DexP2PCode,
		[]interface{}{"getorderlocktime", []interface{}{}},
	)
	if err != nil {
		return 0, fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract GetAdmin res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return 0, fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return 0, fmt.Errorf("ErrorCode:%v", errorCode)
	}
	lockTime, err := GetRetValue(res, 1, reflect.Int)
	if err != nil {
		return 0, fmt.Errorf("GetRetValue error:%s", err)
	}
	return lockTime.(int), err
}

func (this *DexP2PContract) SetOrderLockTime(ctx *TestFrameworkContext, admin *account.Account, newLockTime int) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexP2PCode,
		[]interface{}{"changeadmin", []interface{}{newLockTime}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract SetOrderLockTime res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}
//
//func deployDexP2P(ctx *TestFrameworkContext) bool {
//	code := DEXP2PCode
//	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
//		code,
//		[]contract.ContractParameterType{contract.String, contract.Array},
//		contract.ContractParameterType(contract.Array),
//		"TestDExP2P",
//		"1.0",
//		"",
//		"",
//		"",
//		types.NEOVM,
//	)
//	if err != nil {
//		ctx.LogError("deployDexP2P DeploySmartContract error:%s", err)
//		return false
//	}
//	//等待出块
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("deployDexP2P WaitForGenerateBlock error:%s", err)
//		return false
//	}
//	//
//	//if !testDexP2PInit(ctx){
//	//	return false
//	//}
//	//buyer := ctx.OntClient.Account1
//	//seller := ctx.OntClient.Account2
//	//orderId := []byte(fmt.Sprint("%d", rand.Int31()))
//	//orderSig, err := crypto.Sign(buyer.PrivateKey, orderId)
//	//if err != nil {
//	//	ctx.LogError("TestDexP2P crypto.Sign error:%s", err)
//	//	return false
//	//}
//	//amount := 10
//	//if !testReceipt(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testMakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
//	//	return false
//	//}
//	//
//	//if !testBuyOrderComplete(ctx, orderId, buyer){
//	//	return false
//	//}
//	//
//	//orderId = []byte(fmt.Sprint("%d", rand.Int31()))
//	//orderSig, err = crypto.Sign(buyer.PrivateKey, orderId)
//	//if err != nil {
//	//	ctx.LogError("TestDexP2P crypto.Sign error:%s", err)
//	//	return false
//	//}
//	//amount = 11
//	//if !testReceipt(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testMakeBuyOrder(ctx, orderSig, orderId, buyer, seller, amount) {
//	//	return false
//	//}
//	//
//	//if !testBuyOrderCancel(ctx, orderId, buyer){
//	//	return false
//	//}
//	return true
//}
//
//func dexP2PInit(ctx *TestFrameworkContext) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		ctx.OntClient.Account1,
//		DEXP2PCode,
//		[]interface{}{"init", []interface{}{}},
//	)
//	if err != nil {
//		ctx.LogError("dexP2PInit error:%s", err)
//		return false
//	}
//	if err != nil {
//		ctx.LogError("dexP2PInit error:%s", err)
//		return false
//	}
//	ctx.LogInfo("dexP2PInit res:%v", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("dexP2PInit getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 && errorCode != 3009 {
//		ctx.LogError("dexP2PInit failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func makeBuyOrder(ctx *TestFrameworkContext, orderSig, orderId []byte, buyer, seller *account.Account, amount int) bool {
//	buyerPk, err := buyer.PublicKey.EncodePoint(true)
//	if err != nil {
//		ctx.LogError("makeBuyOrder PublicKey.EncodePoint error:%s", err)
//		return false
//	}
//	res, err := ctx.Ont.InvokeSmartContract(
//		seller,
//		DEXP2PCode,
//		[]interface{}{"makebuyorder", []interface{}{orderSig, orderId, buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), buyerPk, amount}},
//	)
//	if err != nil {
//		ctx.LogError("makeBuyOrder error:%s", err)
//		return false
//	}
//	ctx.LogInfo("makeBuyOrder res:%v", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("makeBuyOrder getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("makeBuyOrder failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func buyOrderComplete(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DEXP2PCode,
//		[]interface{}{"buyordercomplete", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
//	)
//	if err != nil {
//		ctx.LogError("buyOrderComplete error:%s", err)
//		return false
//	}
//	ctx.LogInfo("buyOrderComplete res:%v", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("buyOrderComplete getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("buyOrderComplete failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func buyOrderCancel(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DEXP2PCode,
//		[]interface{}{"buyordercancel", []interface{}{orderId, buyer.ProgramHash.ToArray()}},
//	)
//	if err != nil {
//		ctx.LogError("buyOrderComplete error:%s", err)
//		return false
//	}
//	ctx.LogInfo("buyOrderComplete res:%v", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("buyOrderComplete getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("buyOrderComplete failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
