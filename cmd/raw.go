package cmd

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/spf13/cobra"
)

var rawCmd = &cobra.Command{
	Use:   "raw",
	Short: "Redirect keyboard to roku",
	Run: func(cmd *cobra.Command, args []string) {
		device := deviceFromFlag(cmd)

		s, err := tcell.NewScreen()
		if err != nil {
			panic(err)
		}
		if err := s.Init(); err != nil {
			panic(err)
		}
		defer s.Fini()

		quit := make(chan struct{}, 1)

		go func() {
			for {
				ev := s.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyBackspace2:
						device.Keypress("back")
					case tcell.KeyHome:
						device.Keypress("home")
					case tcell.KeyLeft:
						device.Keypress("left")
					case tcell.KeyRight:
						device.Keypress("right")
					case tcell.KeyUp:
						device.Keypress("up")
					case tcell.KeyDown:
						device.Keypress("down")
					case tcell.KeyEnter:
						device.Keypress("select")
					case tcell.KeyCtrlC:
						close(quit)
						return
					default:
						device.Keypress(fmt.Sprintf("Lit_%c", ev.Rune()))
					}
				}
			}
		}()

		<-quit
	},
}

func init() {
	rootCmd.AddCommand(rawCmd)
}
