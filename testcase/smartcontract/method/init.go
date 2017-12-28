package method

import (
	. "github.com/ONT_TEST/testframework"
	."github.com/ONT_TEST/testcase/smartcontract/method/utils"
	. "github.com/ONT_TEST/testcase/smartcontract/method/hash"
)

func TestMethod() {
	//Utils
	TFramework.RegTestCase("TestAsBigInteger", TestAsBigInteger)
	TFramework.RegTestCase("TestAsByteArrayBigInteger", TestAsByteArrayBigInteger)
	TFramework.RegTestCase("TestAsByteArrayString", TestAsByteArrayString)
	TFramework.RegTestCase("TestAsString", TestAsString)
	TFramework.RegTestCase("TestRange", TestRange)
	TFramework.RegTestCase("TestTake", TestTake)

	//hash
	TFramework.RegTestCase("TestHash160", TestHash160)
	TFramework.RegTestCase("TestHash256", TestHash256)
	TFramework.RegTestCase("TestSha1", TestSha1)
	TFramework.RegTestCase("TestSha256", TestSha256)
}
