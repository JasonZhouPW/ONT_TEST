package testcase

import (
	"github.com/ONT_TEST/testcase/ontid"
	"github.com/ONT_TEST/testframework"
)

func init() {
	testframework.TFramework.RegTestCase("TestONTID", ontid.TestONTID)
}
