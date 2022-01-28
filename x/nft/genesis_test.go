package nft_test

import (
	"testing"

	keepertest "github.com/ChihuahuaChain/chihuahua/testutil/keeper"
	"github.com/ChihuahuaChain/chihuahua/x/nft"
	"github.com/ChihuahuaChain/chihuahua/x/nft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftKeeper(t)
	nft.InitGenesis(ctx, *k, genesisState)
	got := nft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
