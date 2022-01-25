package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/RaspiRepo/demopoa/testutil/keeper"
	"github.com/RaspiRepo/demopoa/x/demopoa/keeper"
	"github.com/RaspiRepo/demopoa/x/demopoa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DemopoaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
