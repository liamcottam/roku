package cmd

import (
	"os"

	"github.com/liamcottam/roku/pkg"
	"github.com/spf13/cobra"
)

var youtubeCmd = &cobra.Command{
	Use:   "youtube video-id",
	Short: "Play a youtube video",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		deviceFromFlag(cmd).LaunchApp("837", args[0], pkg.MediaTypeNone)
	},
}

func init() {
	rootCmd.AddCommand(youtubeCmd)
}
