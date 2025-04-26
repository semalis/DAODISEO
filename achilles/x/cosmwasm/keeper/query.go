package keeper

import (
	"github.com/olimdzhon/achilles/x/cosmwasm/types"
)

var _ types.QueryServer = Keeper{}
