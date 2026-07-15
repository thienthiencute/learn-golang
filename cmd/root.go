package cmd

import (
	"fmt"
	"gocas/cmd/pubsub"
	"gocas/cmd/schedule"
	"gocas/cmd/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gocas",
	Version: "1.0.0",
	Short:   "go clean architecture service context",
	Long:    `Go clean architecture service context tranning youtube`,
	Run: func(cmd *cobra.Command, arg []string) {
		fmt.Println("----------- Server -------------")
	},
}

func init() {
	rootCmd.AddCommand(server.ServerCmd)
	rootCmd.AddCommand(pubsub.PubSubCmd)
	rootCmd.AddCommand(schedule.ScheduleCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
