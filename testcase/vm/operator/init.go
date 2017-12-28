package operator

import (
	. "github.com/ONT_TEST/testframework"
)

func TestOperation() {
	TFramework.RegTestCase("TestOperationAdd", TestOperationAdd)
	TFramework.RegTestCase("TestOperationSub", TestOperationSub)
	TFramework.RegTestCase("TestOperationMulti", TestOperationMulti)
	TFramework.RegTestCase("TestOperationDivide", TestOperationDivide)
	TFramework.RegTestCase("TestOperationSelfAdd", TestOperationSelfAdd)
	TFramework.RegTestCase("TestOperationSelfSub", TestOperationSelfSub)
	TFramework.RegTestCase("TestOperationLarger", TestOperationLarger)
	TFramework.RegTestCase("TestOperationLargerEqual", TestOperationLargerEqual)
	TFramework.RegTestCase("TestOperationSmaller", TestOperationSmaller)
	TFramework.RegTestCase("TestOperationSmallerEqual", TestOperationSmallerEqual)
	TFramework.RegTestCase("TestOperationEqual", TestOperationEqual)
	TFramework.RegTestCase("TestOperationNotEqual", TestOperationNotEqual)
	TFramework.RegTestCase("TestOperationNegative", TestOperationNegative)
	TFramework.RegTestCase("TestOperationOr", TestOperationOr)
	TFramework.RegTestCase("TestOperationAnd", TestOperationAnd)
	TFramework.RegTestCase("TestOperationLeftShift", TestOperationLeftShift)
	TFramework.RegTestCase("TestOperationRightShift", TestOperationRightShift)
}
