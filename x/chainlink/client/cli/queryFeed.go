package cli

import (
	"context"
	"errors"
	"github.com/ChainSafe/chainlink-cosmos/x/chainlink/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdGetRoundFeedData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getRoundFeedData [roundId] [feedId]",
		Short: "List round feed data",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "" {
				return errors.New("roundId is required")
			}
			roundId := args[0]
			roundIdInt, err := strconv.ParseInt(roundId, 10, 64)
			if err != nil {
				return errors.New("roundId is invalid")
			}
			feedId := args[1]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.GetRoundDataRequest{
				RoundId:    uint64(roundIdInt),
				FeedId:     feedId,
				Pagination: pageReq,
			}

			res, err := queryClient.GetRoundData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdGetLatestFeedData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getLatestFeedData [feedId]",
		Short: "List the latest round feed data",
		RunE: func(cmd *cobra.Command, args []string) error {
			feedId := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.GetLatestRoundDataRequest{
				FeedId: feedId,
			}

			res, err := queryClient.LatestRoundData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}