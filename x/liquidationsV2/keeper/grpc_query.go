package keeper

import (
	"github.com/comdex-official/comdex/x/liquidationsV2/types"
)

var _ types.QueryServer = Keeper{}