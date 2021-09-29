package cmd

import (
	"fmt"
	"os"

	songwhip "github.com/dspacejs/songwhip-cli/src"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "songwhip",
	Short: "Simple CLI to get a song's url on Songwhip.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		songwhip.Get(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
