package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var sizeCmd = &cobra.Command{
	Use:     "size",
	Aliases: []string{"sizes"},
	Short:   "Details of Civo instance sizes",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	rootCmd.AddCommand(sizeCmd)
	sizeCmd.AddCommand(sizeListCmd)

	sizeListCmd.Flags().StringVarP(&filterSize, "filter", "s", "", "filter the result by the type (kubernest, database, inatance)")
}
