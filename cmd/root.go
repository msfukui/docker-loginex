package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type loginOptions struct {
	passwordStdin bool
	password      string
	username      string
	serverAddress string
}

var options loginOptions

type loginInfo struct {
	server   string
	username string
	password string
}

var rootCmd = &cobra.Command{
	Use:   "docker loginex [OPTIONS] [SERVER]",
	Short: "A Docker CLI plugins for slightly extending `docker login` command.",
	Long: "A Docker CLI plugins for slightly extending `docker login` command.\n" +
		"Log in to a Docker registry or cloud backend.\n" +
		"If no server is specified, the default is defined by the daemon.\n" +
		"See also help for `docker login`.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			options.serverAddress = args[0]
		}
		return runLoginex(options)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&options.passwordStdin, "password-stdin", "", false, "Take the password from stdin")
	rootCmd.PersistentFlags().StringVarP(&options.password, "password", "p", "", "password")
	rootCmd.PersistentFlags().StringVarP(&options.username, "username", "u", "", "username")
}

func runLoginex(opts loginOptions) error {
	var login loginInfo

	if err := verifyloginOptions(opts); err != nil {
		return err
	}

	if err := setloginInfo(opts, &login); err != nil {
		return err
	}

	fmt.Println("Hello, docker-loginex")
	fmt.Printf("  loginOptions: %v\n", opts)
	fmt.Printf("  loginInfo: %v\n", login)

	return nil
}

func verifyloginOptions(opt loginOptions) error {
	return nil
}

func setloginInfo(opt loginOptions, info *loginInfo) error {
	return nil
}
