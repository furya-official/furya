package rewards

import (
	"time"

	"github.com/comdex-official/comdex/x/rewards/keeper"
	"github.com/comdex-official/comdex/x/rewards/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	k.TriggerAndUpdateEpochInfos(ctx)

	rewards := k.GetRewards(ctx)
	for _, v := range rewards {
		appId := v.App_mapping_ID
		assetIds := v.Asset_ID
		err := k.IterateLocker(ctx, appId, assetIds)
		if err != nil {
			return
		}
	}
	AppIdsVault := k.GetAppIds(ctx).WhitelistedAppMappingIdsVaults
	for i := range AppIdsVault {
		err := k.IterateVaults(ctx, AppIdsVault[i])
		if err != nil {
			return
		}
	}
	err := k.DistributeExtRewardLocker(ctx)
	if err != nil {
		return
	}
	err = k.DistributeExtRewardVault(ctx)
	if err != nil {
		return
	}
}

// EndBlocker for incentives module.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {}