package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"
	"github.com/gitopia/gitopia/app/params"
	"github.com/gitopia/gitopia/testutil/sample"
	"github.com/gitopia/gitopia/testutil/simapp"
	"github.com/gitopia/gitopia/x/gitopia/keeper"
	gitopiatypes "github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestTokenDistributionSucessNoDistribution(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	gParams := gitopiatypes.Params{}
	minter := gitopiatypes.MinterAccountName
	feeCollector := authtypes.FeeCollectorName
	gitopiaKeeper := app.GitopiaKeeper
	bankKeeper := app.BankKeeper
	accountKeeper := app.AccountKeeper

	// setup
	err := testutil.FundModuleAccount(bankKeeper, ctx, minter, sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50))))
	assert.NoError(t, err)
	gitopiaKeeper.SetParams(ctx, gParams)

	gitopiaKeeper.TokenDistribution(ctx)

	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(feeCollector), params.BaseCoinUnit))
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(0)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(minter), params.BaseCoinUnit))
}

func TestTokenDistributionSucessWithNonEmptyEcosystemAddress(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	gParams := gitopiatypes.Params{
		PoolProportions: gitopiatypes.PoolProportions{
			Ecosystem: &gitopiatypes.DistributionProportion{
				Address: "addr",
			},
		},
	}
	gitopiaKeeper := app.GitopiaKeeper
	gitopiaKeeper.SetParams(ctx, gParams)

	require.Panics(t, func() {
		gitopiaKeeper.TokenDistribution(ctx)
	})
}

func TestTokenDistributionSucessWithNonEmptyTeamAddress(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	gParams := gitopiatypes.Params{
		PoolProportions: gitopiatypes.PoolProportions{
			Team: &gitopiatypes.DistributionProportion{
				Address: "addr",
			},
		},
	}
	gitopiaKeeper := app.GitopiaKeeper
	gitopiaKeeper.SetParams(ctx, gParams)

	require.Panics(t, func() {
		gitopiaKeeper.TokenDistribution(ctx)
	})
}

func TestTokenDistributionSucess(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	ecosystemProportion, _ := sdk.NewDecFromStr("30.0")
	teamProportion, _ := sdk.NewDecFromStr("28.0")
	teamProportion1, _ := sdk.NewDecFromStr("7.5")
	teamProportion2, _ := sdk.NewDecFromStr("92.5")
	teamAddress1 := sample.AccAddress()
	teamAddress2 := sample.AccAddress()
	gParams := gitopiatypes.Params{
		PoolProportions: gitopiatypes.PoolProportions{
			Ecosystem: &gitopiatypes.DistributionProportion{
				Proportion: ecosystemProportion,
			},
			Team: &gitopiatypes.DistributionProportion{
				Proportion: teamProportion,
			},
		},
		TeamProportions: []gitopiatypes.DistributionProportion{
			{Proportion: teamProportion1, Address: teamAddress1},
			{Proportion: teamProportion2, Address: teamAddress2},
		},
	}
	minter := gitopiatypes.MinterAccountName
	feeCollector := authtypes.FeeCollectorName
	gitopiaKeeper := app.GitopiaKeeper
	bankKeeper := app.BankKeeper
	accountKeeper := app.AccountKeeper

	// setup
	err := testutil.FundModuleAccount(bankKeeper, ctx, minter, sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50))))
	assert.NoError(t, err)
	gitopiaKeeper.SetParams(ctx, gParams)

	gitopiaKeeper.TokenDistribution(ctx)

	// NOTE: 58% of 50 is 29. 50-29 = 21
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(21)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(feeCollector), params.BaseCoinUnit))
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(0)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(minter), params.BaseCoinUnit))

	accAdrr := accountKeeper.GetModuleAddress(gitopiatypes.EcosystemIncentivesAccountName)
	// NOTE: 30% of 50 is 15
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(keeper.ECOSYSTEM_INCENTIVES_AMOUNT+15)), bankKeeper.GetBalance(ctx, accAdrr, params.BaseCoinUnit))
	accAdrr = accountKeeper.GetModuleAddress(gitopiatypes.TeamAccountName)
	// NOTE: 28% of 50 is 12.5
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(14)), bankKeeper.GetBalance(ctx, accAdrr, params.BaseCoinUnit))
}

func TestTokenDistributionSucessWhenNoTokenMinted(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "gitopia-1", Time: time.Now().UTC()})
	ecosystemProportion, _ := sdk.NewDecFromStr("30.0")
	teamProportion, _ := sdk.NewDecFromStr("28.0")
	gParams := gitopiatypes.Params{
		PoolProportions: gitopiatypes.PoolProportions{
			Ecosystem: &gitopiatypes.DistributionProportion{
				Proportion: ecosystemProportion,
			},
			Team: &gitopiatypes.DistributionProportion{
				Proportion: teamProportion,
			},
		},
	}
	minter := gitopiatypes.MinterAccountName
	feeCollector := authtypes.FeeCollectorName
	gitopiaKeeper := app.GitopiaKeeper
	bankKeeper := app.BankKeeper
	accountKeeper := app.AccountKeeper

	gitopiaKeeper.SetParams(ctx, gParams)

	gitopiaKeeper.TokenDistribution(ctx)

	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(0)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(feeCollector), params.BaseCoinUnit))
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(0)), bankKeeper.GetBalance(ctx, accountKeeper.GetModuleAddress(minter), params.BaseCoinUnit))

	accAdrr := accountKeeper.GetModuleAddress(gitopiatypes.EcosystemIncentivesAccountName)
	assert.Equal(t, sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(keeper.ECOSYSTEM_INCENTIVES_AMOUNT)), bankKeeper.GetBalance(ctx, accAdrr, params.BaseCoinUnit))
}
