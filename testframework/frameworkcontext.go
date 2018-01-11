package testframework

import (
	"github.com/ONT_TEST/ont"
	"fmt"
	log4 "github.com/alecthomas/log4go"
	"time"
	"encoding/hex"
	"math/big"
	"bytes"

)

type TestFrameworkContext struct {
	Ont            *ont.Ontology
	OntClient      *ont.OntClient
	OntAsset       *ont.OntAsset
	BenchThreadNum int
	BenchLastTime  time.Duration
	failNowCh      chan interface{}
}

func NewTestFrameworkContext(ont *ont.Ontology,
	ontClient *ont.OntClient,
	ontAsset *ont.OntAsset,
	failNowCh chan interface{},
	benchThreadNum int,
	benchLastTime time.Duration) *TestFrameworkContext {
	return &TestFrameworkContext{
		Ont:            ont,
		OntClient:      ontClient,
		OntAsset:       ontAsset,
		failNowCh:      failNowCh,
		BenchThreadNum: benchThreadNum,
		BenchLastTime:  benchLastTime,
	}
}

func (this *TestFrameworkContext) LogInfo(arg0 interface{}, args ...interface{}) {
	log4.Info(arg0, args...)
}

func (this *TestFrameworkContext) LogError(arg0 interface{}, args ...interface{}) {
	log4.Error(arg0, args...)
}

func (this *TestFrameworkContext) FailNow() {
	select {
	case <-this.failNowCh:
	default:
		close(this.failNowCh)
	}
}

func (this *TestFrameworkContext) AssertToInt(value interface{}, expect int) error {
	v, ok := value.(float64)
	if !ok {
		return fmt.Errorf("Assert:%v to float failed", value)
	}
	if int(v) != expect {
		return fmt.Errorf("%v not equal:%v", value, expect)
	}
	return nil
}

func (this *TestFrameworkContext) AssertToUint(value interface{}, expect uint) error {
	v, ok := value.(float64)
	if !ok {
		return fmt.Errorf("Assert:%v to float failed", value)
	}
	if uint(v) != expect {
		return fmt.Errorf("%v not equal:%v", value, expect)
	}
	return nil
}

func (this *TestFrameworkContext) AssertToBoolean(value interface{}, expect bool) error {
	v, ok := value.(bool)
	if !ok {
		return fmt.Errorf("Assert:%v to boolean failed", value)
	}
	if v != expect {
		return fmt.Errorf("%v not equal:%v", value, expect)
	}
	return nil
}

func (this *TestFrameworkContext) AssertToString(value interface{}, expect string) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("Assert:%v to string failed", value)
	}
	if v != expect {
		return fmt.Errorf("%v not equal:%v", value, expect)
	}
	return nil
}

func (this *TestFrameworkContext) AssertToByteArray(value interface{}, expect []byte) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("Assert:%v to string failed", value)
	}
	d, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("hex.DecodeString:%s error:%s", v, err)
	}
	if !bytes.EqualFold(d, expect) {
		return fmt.Errorf("%x not equal:%x", d, expect)
	}
	return nil
}

func (this *TestFrameworkContext) AssertBigInteger(value interface{}, expect *big.Int) error{
	v, ok := value.(float64)
	if !ok {
		return  fmt.Errorf("Assert:%v to big.int failed", value)
	}

	if big.NewInt(int64(v)).Cmp(expect) != 0{
		return fmt.Errorf("%v not equal:%v", v, expect)
	}
	return nil
}
