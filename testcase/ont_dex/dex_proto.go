package ont_dex

import (
	"fmt"
	. "github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
	"reflect"
)

func init() {
	fmt.Printf("-------> DexProto CodeHash:%x Reverse:%x\n", DexProto.CodeHash().ToArray(), DexProto.CodeHash().ToArrayReverse())
}

var DexProtoCode = "0126c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e52756e74696d652e47657454726967676572009c6c766b53527ac46c766b53c3641d00616c766b52c30002d007c46c766b52c36c766b54527ac462e4046168164e656f2e52756e74696d652e47657454726967676572609c6c766b55527ac46c766b55c364a104616c766b00c304696e6974876c766b56527ac46c766b56c3646c00616c766b51c3c0529c009c6c766b59527ac46c766b59c3641d00616c766b52c30002d107c46c766b52c36c766b54527ac4626b046c766b51c300c36c766b57527ac46c766b51c351c36c766b58527ac46c766b57c36c766b58c3617c6549046c766b54527ac46236046c766b00c30b6f6e6d616b656f72646572876c766b5a527ac46c766b5ac3648000616c766b51c3c0539c009c6c766b5e527ac46c766b5ec3641d00616c766b52c30002d107c46c766b52c36c766b54527ac462e1036c766b51c300c36c766b5b527ac46c766b51c351c36c766b5c527ac46c766b51c352c36c766b5d527ac46c766b5bc36c766b5cc36c766b5dc3615272650f056c766b54527ac46298036c766b00c30f6f6e6f72646572636f6d706c657465876c766b5f527ac46c766b5fc3648600616c766b51c3c0539c009c6c766b0113527ac46c766b0113c3641d00616c766b52c30002d107c46c766b52c36c766b54527ac4623d036c766b51c300c36c766b60527ac46c766b51c351c36c766b0111527ac46c766b51c352c36c766b0112527ac46c766b60c36c766b0111c36c766b0112c36152726593056c766b54527ac462f0026c766b00c30d6f6e6f7264657263616e63656c876c766b0114527ac46c766b0114c3648800616c766b51c3c0539c009c6c766b0118527ac46c766b0118c3641d00616c766b52c30002d107c46c766b52c36c766b54527ac46295026c766b51c300c36c766b0115527ac46c766b51c351c36c766b0116527ac46c766b51c352c36c766b0117527ac46c766b0115c36c766b0116c36c766b0117c36152726525076c766b54527ac46246026c766b00c30b6368616e676561646d696e876c766b0119527ac46c766b0119c3645c00616c766b51c3c0519c009c6c766b011b527ac46c766b011bc3641d00616c766b52c30002d107c46c766b52c36c766b54527ac462ed016c766b51c300c36c766b011a527ac46c766b011ac361650c0b6c766b54527ac462ca016c766b00c30867657461646d696e876c766b011c527ac46c766b011cc3641200616165090c6c766b54527ac4629b016c766b00c30961646463616c6c6572876c766b011d527ac46c766b011dc3645c00616c766b51c3c0519c009c6c766b011f527ac46c766b011fc3641d00616c766b52c30002d107c46c766b52c36c766b54527ac46244016c766b51c300c36c766b011e527ac46c766b011ec3616530076c766b54527ac46221016c766b00c30c64656c65746563616c6c6572876c766b0120527ac46c766b0120c3645c00616c766b51c3c0519c009c6c766b0122527ac46c766b0122c3641d00616c766b52c30002d107c46c766b52c36c766b54527ac462c7006c766b51c300c36c766b0121527ac46c766b0121c36165f1076c766b54527ac462a4006c766b00c316636865636b63616c6c65727065726d69737373696f6e876c766b0123527ac46c766b0123c3645c00616c766b51c3c0519c009c6c766b0125527ac46c766b0125c3641d00616c766b52c30002d107c46c766b52c36c766b54527ac46240006c766b51c300c36c766b0124527ac46c766b0124c36165a5086c766b54527ac4621d00616c766b52c30002d207c46c766b52c36c766b54527ac46203006c766b54c3616c756659c56b6c766b00527ac46c766b51527ac46151c56c766b52527ac46c766b52c30000c46168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e617c680f4e656f2e53746f726167652e4765746c766b53527ac46c766b53c3c000a06c766b56527ac46c766b56c3641d00616c766b52c30002d907c46c766b52c36c766b57527ac462cc006168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e6c766b00c3615272680f4e656f2e53746f726167652e5075746101006c766b51c36c766b54c39c6c766b58527ac46c766b58c3641300616c766b52c36c766b57527ac46262000663616c6c65726c766b51c3617c65bc096c766b55527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b55c36c766b51c3615272680f4e656f2e53746f726167652e507574616c766b52c36c766b57527ac46203006c766b57c3616c75665bc56b6c766b00527ac46c766b51527ac46c766b52527ac46152c56c766b53527ac46c766b53c30000c461682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b54527ac46c766b53c3516c766b54c3c46c766b54c36165e406009c6c766b58527ac46c766b58c3641d00616c766b53c30002d407c46c766b53c36c766b59527ac462870052c576006c766b00c3c476516c766b52c3c46c766b55527ac4046c6f636b6c766b55c3617c673a587e9ea0d89f41ff8cbc260675b3336b5555e36c766b56527ac46c766b56c300c36c766b57527ac46c766b57c3009c009c6c766b5a527ac46c766b5ac3641100616c766b53c3006c766b57c3c4616c766b53c36c766b59527ac46203006c766b59c3616c75665fc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c461682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b54527ac46c766b54c36165c405009c6c766b5a527ac46c766b5ac3641d00616c766b53c30002d407c46c766b53c36c766b5b527ac462a30152c576006c766b00c3c476516c766b52c3c46c766b55527ac406756e6c6f636b6c766b55c3617c673a587e9ea0d89f41ff8cbc260675b3336b5555e36c766b56527ac46c766b56c300c36c766b57527ac46c766b57c3009c009c6c766b5c527ac46c766b5cc3641f00616c766b53c3006c766b57c3c46c766b53c36c766b5b527ac4621e0152c576006c766b00c3c476516c766b52c3c46c766b58527ac4077061796d656e746c766b58c3617c673a587e9ea0d89f41ff8cbc260675b3336b5555e36c766b56527ac46c766b56c300c36c766b57527ac46c766b57c3009c009c6c766b5d527ac46c766b5dc3641f00616c766b53c3006c766b57c3c46c766b53c36c766b5b527ac462980052c576006c766b51c3c476516c766b52c3c46c766b59527ac407726563656970746c766b59c3617c673a587e9ea0d89f41ff8cbc260675b3336b5555e36c766b56527ac46c766b56c300c36c766b57527ac46c766b57c3009c009c6c766b5e527ac46c766b5ec3641f00616c766b53c3006c766b57c3c46c766b53c36c766b5b527ac46212006c766b53c36c766b5b527ac46203006c766b5bc3616c75665bc56b6c766b00527ac46c766b51527ac46c766b52527ac46151c56c766b53527ac46c766b53c30000c461682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e67536372697074486173686c766b54527ac46c766b54c361658803009c6c766b58527ac46c766b58c3641d00616c766b53c30002d407c46c766b53c36c766b59527ac462970052c576006c766b00c3c476516c766b52c3c46c766b55527ac406756e6c6f636b6c766b55c3617c673a587e9ea0d89f41ff8cbc260675b3336b5555e36c766b56527ac46c766b56c300c36c766b57527ac46c766b57c3009c009c6c766b5a527ac46c766b5ac3641f00616c766b53c3006c766b57c3c46c766b53c36c766b59527ac46212006c766b53c36c766b59527ac46203006c766b59c3616c756657c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b54527ac46c766b54c3641d00616c766b51c30002da07c46c766b51c36c766b55527ac462ad006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b56527ac46c766b56c3641d00616c766b51c30002d707c46c766b51c36c766b55527ac46262000663616c6c65726c766b00c3617c65e6036c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c36c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b55527ac46203006c766b55c3616c756657c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b54527ac46c766b54c3641d00616c766b51c30002da07c46c766b51c36c766b55527ac462aa006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b56527ac46c766b56c3641d00616c766b51c30002d707c46c766b51c36c766b55527ac4625f000663616c6c65726c766b00c3617c65a8026c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b53c3617c68124e656f2e53746f726167652e44656c657465616c766b51c36c766b55527ac46203006c766b55c3616c756653c56b6c766b00527ac46152c56c766b51527ac46c766b51c30000c46c766b51c3516c766b00c361651c00c46c766b51c36c766b52527ac46203006c766b52c3616c756654c56b6c766b00527ac4610663616c6c65726c766b00c3617c65f4016c766b51527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b51c3617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c000a06c766b53527ac46203006c766b53c3616c756656c56b6c766b00527ac46151c56c766b51527ac46c766b51c30000c46168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e617c680f4e656f2e53746f726167652e4765746c766b52527ac46c766b52c3c0009c6c766b53527ac46c766b53c3641d00616c766b51c30002da07c46c766b51c36c766b54527ac4629b006c766b52c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b55527ac46c766b55c3641d00616c766b51c30002d707c46c766b51c36c766b54527ac46250006168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e6c766b00c3615272680f4e656f2e53746f726167652e507574616c766b51c36c766b54527ac46203006c766b54c3616c756652c56b6152c56c766b00527ac46c766b00c30000c46c766b00c3516168164e656f2e53746f726167652e476574436f6e746578740a70726f746f61646d696e617c680f4e656f2e53746f726167652e476574c46c766b00c36c766b51527ac46203006c766b51c3616c756653c56b6c766b00527ac46c766b51527ac4616c766b00c36c766b51c37e6c766b52527ac46203006c766b52c3616c7566"
var DexProto = NewDexProto()

type DexProtoContract struct{}

func NewDexProto() *DexProtoContract {
	return &DexProtoContract{}
}
func (this *DexProtoContract) CodeHash() *common.Uint160 {
	c, _ := common.HexToBytes(DexProtoCode)
	hashCode, _ := common.ToCodeHash(c)
	return &hashCode
}

func (this *DexProtoContract) Deploy(ctx *TestFrameworkContext) error {
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		DexProtoCode,
		[]contract.ContractParameterType{contract.String, contract.Array},
		contract.ContractParameterType(contract.Array),
		"DexProtoContract",
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

func (this *DexProtoContract) Init(ctx *TestFrameworkContext, admin *account.Account, caller []byte) error {
	if caller == nil {
		caller = []byte("")
	}
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexProtoCode,
		[]interface{}{"init", []interface{}{admin.ProgramHash.ToArray(), caller}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("dexProtoInit res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexProtoContract) AddCaller(ctx *TestFrameworkContext, admin *account.Account, caller []byte) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexProtoCode,
		[]interface{}{"addcaller", []interface{}{caller}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("addProtoCaller res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexProtoContract) DeleteCaller(ctx *TestFrameworkContext, admin *account.Account, caller []byte) error {
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexProtoCode,
		[]interface{}{"deletecaller", []interface{}{caller}},
	)
	if err != nil {
		return fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DeleteProtoCaller res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexProtoContract) CheckCallerPermission(ctx *TestFrameworkContext, caller []byte) (bool, error) {
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		DexProtoCode,
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

func (this *DexProtoContract) ChangeAdmin(ctx *TestFrameworkContext, admin, newAdmin *account.Account)error{
	res, err := ctx.Ont.InvokeSmartContract(
		admin,
		DexProtoCode,
		[]interface{}{"changeadmin", []interface{}{newAdmin.ProgramHash.ToArray()}},
	)
	if err != nil {
		return  fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexProtoContract ChangeAdmin res:%s", res)
	errorCode, err := GetErrorCode(res)
	if err != nil {
		return  fmt.Errorf("GetErrorCode error:%s", err)
	}
	if errorCode != 0 {
		return  fmt.Errorf("ErrorCode:%v", errorCode)
	}
	return nil
}

func (this *DexProtoContract) GetAdmin(ctx *TestFrameworkContext)(string, error){
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Admin,
		DexProtoCode,
		[]interface{}{"getadmin", []interface{}{}},
	)
	if err != nil {
		return "", fmt.Errorf("InvokeSmartContract error:%s", err)
	}
	ctx.LogInfo("DexProtoContract GetAdmin res:%s", res)
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

//
//func deployDexProto(ctx *TestFrameworkContext) bool {
//	code := DExProtoCode
//	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
//		code,
//		[]contract.ContractParameterType{contract.String, contract.Array},
//		contract.ContractParameterType(contract.Array),
//		"TestDexProto",
//		"1.0",
//		"",
//		"",
//		"",
//		types.NEOVM,
//	)
//	if err != nil {
//		ctx.LogError("deployDexProto DeploySmartContract error:%s", err)
//		return false
//	}
//	//等待出块
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("deployDexProto WaitForGenerateBlock error:%s", err)
//		return false
//	}
//	//admin := ctx.OntClient.Admin
//	//if !testDexProtoInit(ctx, admin) {
//	//	return false
//	//}
//	//buyer := ctx.OntClient.Account1
//	//seller := ctx.OntClient.Account2
//	//amount := 11
//	//if !testReceipt(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testOnMakeOrder(ctx, buyer, seller, amount) {
//	//	return false
//	//}
//	//if !testOnOrderComplete(ctx, buyer, seller, amount) {
//	//	return false
//	//}
//	//buyer, seller = seller, buyer
//	//amount = 12
//	//if !testReceipt(ctx, buyer, amount) {
//	//	return false
//	//}
//	//if !testOnMakeOrder(ctx, buyer, seller, amount) {
//	//	return false
//	//}
//	//if !testOnOrderCancel(ctx, buyer, seller, amount) {
//	//	return false
//	//}
//	return true
//}
//
//func dexProtoInit(ctx *TestFrameworkContext, admin *account.Account) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		admin,
//		DExProtoCode,
//		[]interface{}{"init", []interface{}{admin.ProgramHash.ToArray(), []byte("")}},
//	)
//	if err != nil {
//		ctx.LogError("dexProtoInit error:%s", err)
//		return false
//	}
//	ctx.LogInfo("dexProtoInit res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("dexProtoInit getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 && errorCode != 2009 {
//		ctx.LogError("dexProtoInit failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func addProtoCaller(ctx *TestFrameworkContext, admin *account.Account, caller []byte) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		admin,
//		DExProtoCode,
//		[]interface{}{"addcaller", []interface{}{caller}},
//	)
//	if err != nil {
//		ctx.LogError("addProtoCaller error:%s", err)
//		return false
//	}
//	ctx.LogInfo("addProtoCaller res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("addProtoCaller getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("addProtoCaller failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
////
//func testOnMakeOrder(ctx *TestFrameworkContext, buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onmakeorder", []interface{}{buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnMakeOrder error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnMakeOrder res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnMakeOrder getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnMakeOrder failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testOnOrderComplete(ctx *TestFrameworkContext, buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onordercomplete", []interface{}{buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnOrderComplete error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnOrderComplete res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnOrderComplete getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnOrderComplete failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
//
//func testOnOrderCancel(ctx *TestFrameworkContext, buyer, seller *account.Account, amount int) bool {
//	res, err := ctx.Ont.InvokeSmartContract(
//		buyer,
//		DExProtoCode,
//		[]interface{}{"onordercancel", []interface{}{buyer.ProgramHash.ToArray(), seller.ProgramHash.ToArray(), amount}},
//	)
//	if err != nil {
//		ctx.LogError("testOnOrderCancel error:%s", err)
//		return false
//	}
//	ctx.LogInfo("testOnOrderCancel res:%s", res)
//	errorCode, err := GetErrorCode(res)
//	if err != nil {
//		ctx.LogError("testOnOrderCancel getErrorCode error:%s", err)
//		return false
//	}
//	if errorCode != 0 {
//		ctx.LogError("testOnOrderCancel failed errorCode:%d", errorCode)
//		return false
//	}
//	return true
//}
