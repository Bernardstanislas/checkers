package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/bernardstanislas/checkers/x/checkers/types"
    "github.com/bernardstanislas/checkers/x/checkers/keeper"
    keepertest "github.com/bernardstanislas/checkers/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
