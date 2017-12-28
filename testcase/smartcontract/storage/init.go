package storage

import (
	. "github.com/ONT_TEST/testframework"
)

func TestStorage(){
	TFramework.RegTestCase("TestStorageGetAndPutAndDelete",TestStorageGetAndPutAndDelete)
}