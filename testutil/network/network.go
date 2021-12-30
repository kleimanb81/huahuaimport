package network

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/tendermint/spm/cosmoscmd"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmdb "github.com/tendermint/tm-db"

	"github.com/ChihuahuaChain/chihuahua/app"
	"github.com/ChihuahuaChain/chihuahua/testutil"
)

type (
	Network = network.Network
	Config  = network.Config
)

// New creates instance with fully configured Chihuahua network, along with
// creating a funded account for each name provided
func New(t *testing.T, config network.Config, accounts ...string) (*network.Network, keyring.Keyring) {
	kr := testutil.GenerateKeyring(t)

	// add genesis accounts
	genAuthAccs := make([]authtypes.GenesisAccount, len(accounts))
	genBalances := make([]banktypes.Balance, len(accounts))
	for i, name := range accounts {
		a, b := newGenAccout(kr, name, 1000000000000)
		genAuthAccs[i] = a
		genBalances[i] = b
	}

	config, err := addGenAccounts(config, genAuthAccs, genBalances)
	if err != nil {
		panic(err)
	}

	net := network.New(t, config)
	t.Cleanup(net.Cleanup)
	return net, kr
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
func DefaultConfig() network.Config {
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	return network.Config{
		Codec:             encoding.Marshaler,
		TxConfig:          encoding.TxConfig,
		LegacyAmino:       encoding.Amino,
		InterfaceRegistry: encoding.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			return app.New(
				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
				encoding,
				simapp.EmptyAppOptions{},
				baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Marshaler),
		TimeoutCommit:   2 * time.Second,
		ChainID:         "chain-" + tmrand.NewRand().Str(6),
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
		PruningStrategy: storetypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}
}

// addGenAccounts adds the provided accounts to the genesis state in the config
func addGenAccounts(cfg network.Config, genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance) (network.Config, error) {
	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return cfg, err
	}

	authGenState.Accounts = append(authGenState.Accounts, accounts...)
	cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = append(bankGenState.Balances, genBalances...)
	cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankGenState)

	return cfg, nil
}

// newGenAccount creates a genesis account with some default staking asset and
// some unique token.
func newGenAccout(kr keyring.Keyring, name string, amount int64) (authtypes.GenesisAccount, banktypes.Balance) {
	info, _, err := kr.NewMnemonic(name, keyring.English, "", "", hd.Secp256k1)
	if err != nil {
		panic(err)
	}

	// create coin
	balances := sdk.NewCoins(
		sdk.NewCoin(fmt.Sprintf("%stoken", name), sdk.NewInt(amount)),
		sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(amount)),
	)

	bal := banktypes.Balance{
		Address: info.GetAddress().String(),
		Coins:   balances.Sort(),
	}

	return authtypes.NewBaseAccount(info.GetAddress(), info.GetPubKey(), 0, 0), bal
}
