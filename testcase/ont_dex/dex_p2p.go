package ont_dex

import (
	"fmt"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"reflect"
)

func init() {
	fmt.Printf("-------> DexP2P CodeHash:%x Reverse:%x\n", DexP2P.CodeHash().ToArray(), DexP2P.CodeHash().ToArrayReverse())
}

var DexP2PCode = "0126c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e52756e74696d652e47657454726967676572009c6c766b53527ac46c766b53c3641d00616c766b52c30002b80bc46c766b52c36c766b54527ac462ea046168164e656f2e52756e74696d652e47657454726967676572609c6c766b55527ac46c766b55c364a704616c766b00c304696e6974876c766b56527ac46c766b56c3646c00616c766b51c3c0529c009c6c766b59527ac46c766b59c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac46271046c766b51c300c36c766b57527ac46c766b51c351c36c766b58527ac46c766b57c36c766b58c3617c654f046c766b54527ac4623c046c766b00c30c6d616b656275796f72646572876c766b5a527ac46c766b5ac364d700616c766b51c3c0569c009c6c766b0111527ac46c766b0111c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac462e4036c766b51c300c36c766b5b527ac46c766b51c351c36c766b5c527ac46c766b51c352c36c766b5d527ac46c766b51c353c36c766b5e527ac46c766b51c354c36c766b5f527ac46c766b51c355c36c766b60527ac46c766b5bc36c766b5cc36c766b5dc36c766b5ec36c766b5fc36c766b60c361557951795772755172755479527956727552727553795379557275537275657e046c766b54527ac46246036c766b00c3106275796f72646572636f6d706c657465876c766b0112527ac46c766b0112c3647200616c766b51c3c0529c009c6c766b0115527ac46c766b0115c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac462e8026c766b51c300c36c766b0113527ac46c766b51c351c36c766b0114527ac46c766b0113c36c766b0114c3617c6528076c766b54527ac462af026c766b00c30e6275796f7264657263616e63656c876c766b0116527ac46c766b0116c3647200616c766b51c3c0529c009c6c766b0119527ac46c766b0119c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac46253026c766b51c300c36c766b0117527ac46c766b51c351c36c766b0118527ac46c766b0117c36c766b0118c3617c65ea096c766b54527ac4621a026c766b00c31373656c6c6572747279636c6f73656f72646572876c766b011a527ac46c766b011ac3647200616c766b51c3c0529c009c6c766b011d527ac46c766b011dc3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac462b9016c766b51c300c36c766b011b527ac46c766b51c351c36c766b011c527ac46c766b011bc36c766b011cc3617c65a50c6c766b54527ac46280016c766b00c30b6368616e676561646d696e876c766b011e527ac46c766b011ec3645c00616c766b51c3c0519c009c6c766b0120527ac46c766b0120c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac46227016c766b51c300c36c766b011f527ac46c766b011fc3616571106c766b54527ac46204016c766b00c30867657461646d696e876c766b0121527ac46c766b0121c36412006161656a116c766b54527ac462d5006c766b00c3107365746f726465726c6f636b74696d65876c766b0122527ac46c766b0122c3645c00616c766b51c3c0519c009c6c766b0124527ac46c766b0124c3641d00616c766b52c30002b90bc46c766b52c36c766b54527ac46277006c766b51c300c36c766b0123527ac46c766b0123c3616552116c766b54527ac46254006c766b00c3106765746f726465726c6f636b74696d65876c766b0125527ac46c766b0125c364120061616578126c766b54527ac4621d00616c766b52c30002ba0bc46c766b52c36c766b54527ac46203006c766b54c3616c756656c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e617c680f4e656f2e53746f726167652e4765746c766b53527ac46c766b53c3c000a06c766b54527ac46c766b54c3641d00616c766b52c30002c10bc46c766b52c36c766b55527ac4628f006168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578740d6f726465726c6f636b74696d656c766b51c3615272680f4e656f2e53746f726167652e507574616c766b52c36c766b55527ac46203006c766b55c3616c75660115c56b6c766b00527ac46c766b51527ac46c766b52527ac46c766b53527ac46c766b54527ac46c766b55527ac46151c56c766b56527ac46c766b56c30000c46c766b53c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b0111527ac46c766b0111c3641e00616c766b56c30002c00bc46c766b56c36c766b0112527ac462ac026168184e656f2e426c6f636b636861696e2e4765744865696768746168184e656f2e426c6f636b636861696e2e4765744865616465726c766b57527ac46c766b57c36168174e656f2e4865616465722e47657454696d657374616d706c766b58527ac4066f62757965726c766b51c3617c6592106c766b59527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b59c3617c680f4e656f2e53746f726167652e4765746c766b5a527ac46c766b5ac3c000a06c766b0113527ac46c766b0113c3641e00616c766b56c30002bc0bc46c766b56c36c766b0112527ac462c501076f616d6f756e746c766b51c3617c650d106c766b5b527ac4076f73656c6c65726c766b51c3617c65f40f6c766b5c527ac4056f74696d656c766b51c3617c65dd0f6c766b5d527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b5bc36c766b55c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b59c36c766b52c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b5cc36c766b53c3615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b5dc36c766b58c3615272680f4e656f2e53746f726167652e5075746153c576006c766b52c3c476516c766b53c3c476526c766b55c3c46c766b5e527ac40b6f6e6d616b656f726465726c766b5ec3617c67423ecabefbaf2f639321934d68e2568bb04dd7966c766b5f527ac46c766b5fc300c36c766b60527ac46c766b60c3009c009c6c766b0114527ac46c766b0114c3641400616c766b5fc36c766b0112527ac46213006c766b56c36c766b0112527ac46203006c766b0112c3616c75660112c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b5d527ac46c766b5dc3641d00616c766b52c30002c00bc46c766b52c36c766b5e527ac462e202066f62757965726c766b00c3617c65d30d6c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c0009c6c766b5f527ac46c766b5fc3641d00616c766b52c30002bd0bc46c766b52c36c766b5e527ac46261026c766b54c36c766b51c39c009c6c766b60527ac46c766b60c3641d00616c766b52c30002bf0bc46c766b52c36c766b5e527ac4622b02076f616d6f756e746c766b00c3617c651b0d6c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac4076f73656c6c65726c766b00c3617c65ca0c6c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c680f4e656f2e53746f726167652e4765746c766b58527ac453c576006c766b54c3c476516c766b58c3c476526c766b56c3c46c766b59527ac40f6f6e6f72646572636f6d706c6574656c766b59c3617c67423ecabefbaf2f639321934d68e2568bb04dd7966c766b5a527ac46c766b5ac300c36c766b5b527ac46c766b5bc3009c009c6c766b0111527ac46c766b0111c3641300616c766b5ac36c766b5e527ac462fd00056f74696d656c766b00c3617c65ef0b6c766b5c527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b5cc3617c68124e656f2e53746f726167652e44656c657465616c766b52c36c766b5e527ac46203006c766b5ec3616c75660112c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b5d527ac46c766b5dc3641d00616c766b52c30002c00bc46c766b52c36c766b5e527ac462e002066f62757965726c766b00c3617c657c0a6c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c0009c6c766b5f527ac46c766b5fc3641d00616c766b52c30002bd0bc46c766b52c36c766b5e527ac4625f026c766b54c36c766b51c39c009c6c766b60527ac46c766b60c3641d00616c766b52c30002bf0bc46c766b52c36c766b5e527ac4622902076f616d6f756e746c766b00c3617c65c4096c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac4076f73656c6c65726c766b00c3617c6573096c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c680f4e656f2e53746f726167652e4765746c766b58527ac453c576006c766b54c3c476516c766b58c3c476526c766b56c3c46c766b59527ac40d6f6e6f7264657263616e63656c6c766b59c3617c67423ecabefbaf2f639321934d68e2568bb04dd7966c766b5a527ac46c766b5ac300c36c766b5b527ac46c766b5bc3009c009c6c766b0111527ac46c766b0111c3641300616c766b5ac36c766b5e527ac462fd00056f74696d656c766b00c3617c659a086c766b5c527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b5cc3617c68124e656f2e53746f726167652e44656c657465616c766b52c36c766b5e527ac46203006c766b5ec3616c75660117c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46c766b51c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b0111527ac46c766b0111c3641e00616c766b52c30002c00bc46c766b52c36c766b0112527ac462cf03066f62757965726c766b00c3617c6524076c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c680f4e656f2e53746f726167652e4765746c766b54527ac46c766b54c3c0009c6c766b0113527ac46c766b0113c3641e00616c766b52c30002bd0bc46c766b52c36c766b0112527ac4624b03076f73656c6c65726c766b00c3617c659f066c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46c766b56c36c766b51c39c009c6c766b0114527ac46c766b0114c3641e00616c766b52c30002bf0bc46c766b52c36c766b0112527ac462c102056f74696d656c766b00c3617c6517066c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c680f4e656f2e53746f726167652e4765746c766b58527ac46168184e656f2e426c6f636b636861696e2e4765744865696768746168184e656f2e426c6f636b636861696e2e4765744865616465726c766b59527ac46c766b59c36168174e656f2e4865616465722e47657454696d657374616d706c766b5a527ac46165030551c36c766b5b527ac46c766b5ac36c766b58c3946c766b5bc39f6c766b0115527ac46c766b0115c3641e00616c766b52c30002be0bc46c766b52c36c766b0112527ac462c501076f616d6f756e746c766b00c3617c6519056c766b5c527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b5cc3617c680f4e656f2e53746f726167652e4765746c766b5d527ac453c576006c766b54c3c476516c766b56c3c476526c766b5dc3c46c766b5e527ac40f6f6e6f72646572636f6d706c6574656c766b5ec3617c67423ecabefbaf2f639321934d68e2568bb04dd7966c766b5f527ac46c766b5fc300c36c766b60527ac46c766b60c3009c009c6c766b0116527ac46c766b0116c3641400616c766b5fc36c766b0112527ac462e7006168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b55c3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b5cc3617c68124e656f2e53746f726167652e44656c657465616168164e656f2e53746f726167652e476574436f6e746578746c766b57c3617c68124e656f2e53746f726167652e44656c657465616c766b52c36c766b0112527ac46203006c766b0112c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002c20bc46c766b51c36c766b54527ac46299006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3641d00616c766b51c30002c00bc46c766b51c36c766b54527ac4624e006168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756657c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46c766b00c3009f6c766b53527ac46c766b53c3641d00616c766b51c30002b90bc46c766b51c36c766b54527ac4620b016168164e656f2e53746f726167652e476574436f6e746578740870327061646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b55527ac46c766b55c3641d00616c766b51c30002c20bc46c766b51c36c766b54527ac4629e006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b56527ac46c766b56c3641d00616c766b51c30002c00bc46c766b51c36c766b54527ac46253006168164e656f2e53746f726167652e476574436f6e746578740d6f726465726c6f636b74696d656c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740d6f726465726c6f636b74696d65617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c37e6c766b52527ac46203006c766b52c3616c7566"
var DexP2P = NewDexP2PContract()

type DexP2PContract struct{}

func NewDexP2PContract() *DexP2PContract {
	return &DexP2PContract{}
}

func (this *DexP2PContract) GetCode() string {
	return DexP2PCode
}

func (this *DexP2PContract) CodeHash() *common.Uint160 {
	c, _ := common.HexToBytes(this.GetCode())
	hashCode, _ := common.ToCodeHash(c)
	return &hashCode
}

func (this *DexP2PContract) Deploy(ctx *TestFrameworkContext) error {
	ctx.LogInfo("DexP2PContract Deploy")
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		this.GetCode(),
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
	//_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	//if err != nil {
	//	return fmt.Errorf("WaitForGenerateBlock error:%s", err)
	//}
	return nil
}

func (this *DexP2PContract) Init(ctx *TestFrameworkContext, admin *account.Account, lockTime int) error {
	if lockTime == 0 {
		lockTime = 24 * 3600
	}
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		this.GetCode(),
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
		this.GetCode(),
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
		this.GetCode(),
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
		this.GetCode(),
		[]interface{}{"makebuyorder", []interface{}{
			orderSig, orderId,
			buyer.ProgramHash.ToArray(),
			seller.ProgramHash.ToArray(),
			buyerPk, ctx.Ont.MakeAssetAmount(amount)}},
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
		this.GetCode(),
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

func (this *DexP2PContract) BuyOrderCancel(ctx *TestFrameworkContext, orderId []byte, buyer *account.Account) error {
	res, err := ctx.Ont.InvokeSmartContract(
		buyer,
		this.GetCode(),
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

func (this *DexP2PContract) SellerTryCloseOrder(ctx *TestFrameworkContext, orderId []byte, seller *account.Account) error {
	res, err := ctx.Ont.InvokeSmartContract(
		seller,
		this.GetCode(),
		[]interface{}{"sellertrycloseorder", []interface{}{orderId, seller.ProgramHash.ToArray()}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("SellerTryCloseOrder res:%v", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}

	return nil
}

func (this *DexP2PContract) GetOrderLockTime(ctx *TestFrameworkContext) (int, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		this.GetCode(),
		[]interface{}{"getorderlocktime", []interface{}{}},
	)
	if err != nil {
		return 0, fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexP2PContract GetOrderLockTime res:%s", res)
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
	ctx.LogInfo("DexP2PContract GetOrderLockTime:%v", lockTime)
	return lockTime.(int), err
}

func (this *DexP2PContract) SetOrderLockTime(ctx *TestFrameworkContext, admin *account.Account, newLockTime int) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		this.GetCode(),
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
