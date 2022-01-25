package keeper_test

import (
	"testing"

	testkeeper "github.com/RaspiRepo/demopoa/testutil/keeper"
	"github.com/RaspiRepo/demopoa/x/demopoa/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DemopoaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
