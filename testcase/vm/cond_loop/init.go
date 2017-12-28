package cond_loop

import (
	. "github.com/ONT_TEST/testframework"
)

func TestCondLoop(){
	TFramework.RegTestCase("TestIfElse", TestIfElse)
	TFramework.RegTestCase("TestSwitch", TestSwitch)
	TFramework.RegTestCase("TestFor", TestFor)
}