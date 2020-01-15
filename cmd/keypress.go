package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var keypressCmd = &cobra.Command{
	Use:   "keypress",
	Short: "Press button",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			os.Exit(1)
		}
		for i := 0; i < len(args); i++ {
			deviceFromFlag(cmd).Keypress(args[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(keypressCmd)
}
