package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bm-commit/blog/testutil/keeper"
	"github.com/bm-commit/blog/x/blog/keeper"
	"github.com/bm-commit/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
