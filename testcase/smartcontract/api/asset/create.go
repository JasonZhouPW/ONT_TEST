package asset

import (
	"github.com/ONT_TEST/testframework"
)

func TestAssetCreate(ctx *testframework.TestFrameworkContext) bool {
	err := InitAsset(ctx)
	if err != nil {
		ctx.LogError("TestAssetCreate CreateAsset error:%s", err)
		return false
	}
	ctx.LogError("TestAssetCreate AssetId:%x ", GetAssetId().ToArray())
	return true
}
