package cmd

import (
    "github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/cosmos/cosmos-sdk/client/tx"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/TorramLabs-Team/TorramChain/x/newtx/types"
)

func CmdCreateNewTx() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "create-new-tx [amount]",
        Short: "Create a new transaction with a specified amount",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            amount, err := sdk.ParseCoinNormalized(args[0])
            if err != nil {
                return err
            }

            msg := types.NewMsgCreateNewTx(clientCtx.GetFromAddress(), amount)

            if err := msg.ValidateBasic(); err != nil {
                return err
            }

            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)

    return cmd
}