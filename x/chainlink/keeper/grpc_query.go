package keeper

import (
	"context"

	"github.com/ChainSafe/chainlink-cosmos/x/chainlink/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

// GetRoundData implements the Query/GetRoundData gRPC method
func (k Keeper) GetRoundData(c context.Context, req *types.GetRoundDataRequest) (*types.GetRoundDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.GetRoundFeedDataByFilter(ctx, req)
}

// LatestRoundData implements the Query/LatestRoundData gRPC method
func (k Keeper) LatestRoundData(c context.Context, req *types.GetLatestRoundDataRequest) (*types.GetLatestRoundDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.GetLatestRoundFeedDataByFilter(ctx, req)
}