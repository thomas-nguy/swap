package keeper

import (
	"github.com/thomas-nguy/swap/x/swap/types"
)

var _ types.QueryServer = Keeper{}
