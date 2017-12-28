package storage


import (
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/core/contract"
	"github.com/Ontology/smartcontract/types"
	"time"
)

func TestStorageGetAndPutAndDelete(ctx *testframework.TestFrameworkContext)bool{
	code := "53C56B0548656C6C6F6C766B00527AC46168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C68124E656F2E53746F726167652E44656C6574656168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C680F4E656F2E53746F726167652E4765746C766B51527AC46C766B51C3C06409000131616C756605576F726C646C766B51527AC46168164E656F2E53746F726167652E476574436F6E746578746C766B00C36C766B51C3615272680F4E656F2E53746F726167652E5075746168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C680F4E656F2E53746F726167652E4765746C766B51C39C6309000132616C75660548656C6C6F6C766B51527AC46168164E656F2E53746F726167652E476574436F6E746578746C766B00C36C766B51C3615272680F4E656F2E53746F726167652E5075746168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C680F4E656F2E53746F726167652E4765746C766B51C39C6309000133616C75666168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C68124E656F2E53746F726167652E44656C6574656168164E656F2E53746F726167652E476574436F6E746578746C766B00C3617C680F4E656F2E53746F726167652E4765746C766B52527AC46C766B52C3C0640C006C766B52C3616C75660130616C7566"
	_, err := ctx.Ont.DeploySmartContract(ctx.OntClient.Account1,
		code,
		[]contract.ContractParameterType{},
		contract.ContractParameterType(contract.String),
		"TestStorage",
		"1.0",
		"",
		"",
		"",
		types.NEOVM,
	)
	if err != nil {
		ctx.LogError("TestStorage DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestStorage WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.InvokeSmartContract(
		ctx.OntClient.Account1,
		code,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestStorage InvokeSmartContract error:%s", err)
		return false
	}
	err = ctx.AssertToString(res, "0")
	if err != nil {
		ctx.LogError("TestStorage test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

public class HelloWorld : SmartContract
{
    public static string Main()
    {
        byte[] key = "Hello".AsByteArray();
        Storage.Delete(Storage.CurrentContext, key);
        byte[] value = Storage.Get(Storage.CurrentContext, key);
        if(value.Length !=0)
        {
            return "1";
        }
        value = "World".AsByteArray();
        Storage.Put(Storage.CurrentContext, key, value);
        if(Storage.Get(Storage.CurrentContext, key) != value)
        {
            return "2";
        }
        value = "Hello".AsByteArray();
		Storage.Put(Storage.CurrentContext, key, value);
        if(Storage.Get(Storage.CurrentContext, key) != value)
        {
            return "3";
        }
        Storage.Delete(Storage.CurrentContext, key);
		byte[] v = Storage.Get(Storage.CurrentContext, key);
		if(v.Length != 0)
        {
            return v.AsString();
        }
        return "0";
    }
}
*/