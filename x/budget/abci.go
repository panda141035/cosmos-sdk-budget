package budget

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/keeper"
	"github.com/tendermint/budget/x/budget/types"
)

// BeginBlocker collects budgets for the current block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	err := k.CollectBudgets(ctx)
	if err != nil {
		panic(err)
	}
}
