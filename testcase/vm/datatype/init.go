package datatype

import (
	. "github.com/ONT_TEST/testframework"
)

func TestDataType() {
	TFramework.RegTestCase("TestBoolean", TestBoolean)
	TFramework.RegTestCase("TestInteger", TestInteger)
	TFramework.RegTestCase("TestString", TestString)
	TFramework.RegTestCase("TestArray", TestArray)
	TFramework.RegTestCase("TestReturnType", TestReturnType)
	TFramework.RegTestCase("TestByteArray", TestByteArray)
}
