package transaction

import (
	. "github.com/ONT_TEST/testcase/smartcontract/api/helper"
	"github.com/ONT_TEST/testframework"
	"github.com/Ontology/account"
	"github.com/Ontology/common"
)

var (
	isTransctionInit = false
	txHash           common.Uint256
)

func getTransferTransaction(ctx *testframework.TestFrameworkContext, form, to *account.Account) (common.Uint256, error) {
	if isTransctionInit {
		return txHash, nil
	}

	err := InitAsset(ctx, form)
	if err != nil {
		return common.Uint256{}, err
	}

	tx, err := Transfer(ctx, AssetId, ctx.OntClient.Account1, ctx.OntClient.Account2, 1000.0)
	if err != nil {
		return common.Uint256{}, err
	}
	isTransctionInit = true
	txHash = tx
	return txHash, nil
}
