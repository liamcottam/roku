package cmd

import (
	"encoding/json"
	"os"

	"github.com/koron/go-ssdp"
	"github.com/spf13/cobra"
)

var searchDevicesCmd = &cobra.Command{
	Use:   "search-devices",
	Short: "Search for devices",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := ssdp.Search("roku:ecp", 1, "")
		if err != nil {
			panic(err)
		}

		je := json.NewEncoder(os.Stdout)
		je.SetIndent("", "  ")
		je.Encode(res)
	},
}

func init() {
	rootCmd.AddCommand(searchDevicesCmd)
}
