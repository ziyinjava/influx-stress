package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	path string
)
var rootCmd = &cobra.Command{
	Use: "influx-stress",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	time.Sleep(1 * time.Minute)
}

func init() {
	cobra.OnInitialize(initConfig)

}
func initConfig() {
	viper.AutomaticEnv()
}
