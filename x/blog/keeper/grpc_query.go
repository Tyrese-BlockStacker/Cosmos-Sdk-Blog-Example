package keeper

import (
	"github.com/bm-commit/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
