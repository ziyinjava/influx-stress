package cmd

import (
	stress2 "github.com/bemyth/influx-stress/stress"
	"github.com/spf13/cobra"
)

/*
速率测试
*/

var (
	ip, port, username, password string
)
var speedCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress Test",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
		ExecuteStress()
	},
}

func init() {
	rootCmd.AddCommand(speedCmd)
	speedCmd.Flags().StringVar(&ip, "ip", "localhost", "ip of influxdb server")
	speedCmd.Flags().StringVar(&port, "port", "8086", "port of influxdb server")
	speedCmd.Flags().StringVarP(&username, "username", "u", "", "username of influxdb server")
	speedCmd.Flags().StringVarP(&password, "password", "p", "", "password of influxdb server")
}
func ExecuteStress() {
	server := stress2.NewServer(ip, port, username, password)
	server.Run()
}
