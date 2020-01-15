package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/liamcottam/roku/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string

var (
	ipFlag   *net.IP
	portFlag *string
)

var rootCmd = &cobra.Command{
	Use: "roku",
}

func deviceFromFlag(cmd *cobra.Command) *pkg.Device {
	var device *pkg.Device
	cmdIPFlag := cmd.Flag("device-ip")
	if cmdIPFlag.Changed {
		device = pkg.DeviceFromIP(*ipFlag, *portFlag)
	} else {
		device = pkg.SearchForDevice()
	}
	return device
}

func Execute() {
	ipFlag = rootCmd.PersistentFlags().IPP("device-ip", "i", nil, "device IP")
	portFlag = rootCmd.PersistentFlags().StringP("device-port", "p", ":8060", "device port")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
