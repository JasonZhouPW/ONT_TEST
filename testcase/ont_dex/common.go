package ont_dex

import (
	"encoding/hex"
	"fmt"
	vmtype "github.com/Ontology/vm/neovm/types"
	"reflect"
)

func GetErrorCode(res interface{}) (int, error) {
	v, err := GetRetValue(res, 0, reflect.Int)
	if err != nil {
		return 0, err
	}
	return v.(int), nil
}

func GetRetValue(res interface{}, index int, vType reflect.Kind) (interface{}, error) {
	rt, ok := res.([]interface{})
	if !ok {
		return 0, fmt.Errorf("%s assert to array failed.", res)
	}
	vs, ok := rt[index].(string)
	if !ok {
		return 0, fmt.Errorf("%s assert string")
	}
	v, err := hex.DecodeString(vs)
	if err != nil {
		return 0, fmt.Errorf("hex.DecodeString:%s error:%s", err)
	}
	switch vType {
	case reflect.Int:
		return int(vmtype.ConvertBytesToBigInteger(v).Int64()), nil
	case reflect.String:
		return vs, nil
	}
	return nil, fmt.Errorf("unsupport type:%v", vType)
}
