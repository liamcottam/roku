package cmd

import (
	"os"

	"github.com/liamcottam/roku/pkg"
	"github.com/spf13/cobra"
)

var twitchCmd = &cobra.Command{
	Use:   "twitch channel",
	Short: "Launch a twitch channel",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		deviceFromFlag(cmd).LaunchApp("50539", args[0], pkg.MediaTypeLive)
	},
}

func init() {
	rootCmd.AddCommand(twitchCmd)
}
