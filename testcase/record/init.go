package record

import (
	. "github.com/ONT_TEST/testframework"
)

func TestRecord() {
	//Record test
	TFramework.RegTestCase("TestRecordTransactionByRecord", TestRecordTransactionByRecord)
	TFramework.RegTestCase("TestRecordTransactionByTransfer", TestRecordTransactionByTransfer)
}

