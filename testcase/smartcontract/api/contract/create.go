package contract

import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

/*

 using Neo.SmartContract.Framework;
 using Neo.SmartContract.Framework.Services.Neo;
 using Neo.SmartContract.Framework.Services.System;

public class Contract1:SmartContract
{
    public static void Main()
    {
        //这里填写合约脚本
        byte[] script = new byte[] { 116, 107, 0, 97, 116, 0, 147, 108, 118, 107, 148, 121, 116, 81, 147, 108, 118, 107, 148, 121, 147, 116, 0, 148, 140, 108, 118, 107, 148, 114, 117, 98, 3, 0, 116, 0, 148, 140, 108, 118, 107, 148, 121, 97, 116, 140, 108, 118, 107, 148, 109, 116, 108, 118, 140, 107, 148, 109, 116, 108, 118, 140, 107, 148, 109, 108, 117, 102 };

        byte[] parameter_list = { 2, 2 };
        byte return_type = 2;
        bool need_storage = false;
        string name = "加法合约示例";
        string version = "1";
        string author = "chris";
        string email = "chris@abc.com";
        string description = "这是一个加法合约，传入两个整型，合约对两个整型进行相加，返回一个整型";

        Contract.Create(script, parameter_list, return_type, need_storage, name, version, author, email, description);
    }
}

*/

func TestContractCreate(ctx *testframework.TestFrameworkContext) bool {
	code := "5ac56b6144746b00617400936c766b94797451936c766b9479937400948c6c766b9472756203007400948c6c766b947961748c6c766b946d746c768c6b946d746c768c6b946d6c75666c766b00527ac40202026c766b51527ac4526c766b52527ac4006c766b53527ac412e58aa0e6b395e59088e7baa6e7a4bae4be8b6c766b54527ac401316c766b55527ac40563687269736c766b56527ac40d6368726973406162632e636f6d6c766b57527ac44c66e8bf99e698afe4b880e4b8aae58aa0e6b395e59088e7baa6efbc8ce4bca0e585a5e4b8a4e4b8aae695b4e59e8befbc8ce59088e7baa6e5afb9e4b8a4e4b8aae695b4e59e8be8bf9be8a18ce79bb8e58aa0efbc8ce8bf94e59b9ee4b880e4b8aae695b4e59e8b6c766b58527ac46c766b00c36c766b51c36c766b52c36c766b53c36c766b54c36c766b55c36c766b56c36c766b57c36c766b58c361587951795a727551727557795279597275527275567953795872755372755579547957727554727568134e656f2e436f6e74726163742e43726561746575616c7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{contract.Void},
		contract.ContractParameterType(contract.Void),
		"TestContractCreate",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestContractCreate DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestContractCreate WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		nil,
	)
	if err != nil {
		ctx.LogError("TestContractCreate InvokeSmartContract error:%s", err)
		return false
	}

	ctx.LogError("TestContractCreate :%+v ", res)
	return true
}
