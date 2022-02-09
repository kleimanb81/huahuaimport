package keeper

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/ChihuahuaChain/chihuahua/x/nft/types"
    "github.com/ChihuahuaChain/chihuahua/x/nft/keeper"
    keepertest "github.com/ChihuahuaChain/chihuahua/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NftKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
