package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "docker-loginex",
	Short: "A Docker CLI plugins for slightly extending `docker login` command.",
	Long:  "Log in to a Docker registry or cloud backend.\nSee also help for `docker login`.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, docker-loginex.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
