package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type dockerCliPluginMetadata struct {
	SchemaVersion    string
	Vendor           string
	Version          string
	ShortDescription string
	URL              string
}

var metadata dockerCliPluginMetadata = dockerCliPluginMetadata{
	SchemaVersion:    "0.1.0",
	Vendor:           "msfukui",
	Version:          "unreleased",
	ShortDescription: "Slightly extending `docker login` command",
	URL:              "https://github.com/msfukui/docker-loginex",
}

type loginOptions struct {
	serverAddress string
}

var options loginOptions

type loginInfo struct {
	server   string
	username string
	password string
}

var rootCmd = &cobra.Command{
	Use:   "docker loginex [SERVER]",
	Short: "A Docker CLI plugins for slightly extending `docker login` command.",
	Long: "A Docker CLI plugins for slightly extending `docker login` command.\n" +
		"Log in to a Docker registry or cloud backend.\n" +
		"If no server is specified, the default is defined by the daemon.\n" +
		"See also help for `docker login`.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			if args[0] == "docker-cli-plugin-metadata" {
				// Return metadata as a requirement of Docker CLI pligins.
				v := getDockerCliPluginMetadata()
				j, _ := json.Marshal(v)
				fmt.Println(string(j))
				return nil
			} else if args[0] == "loginex" {
				// Judged to be called as a subcommand of docker/cli.
				if len(args) > 1 {
					options.serverAddress = args[1]
				}
			} else {
				options.serverAddress = args[0]
			}
		}
		return runLoginex(options)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}

func getDockerCliPluginMetadata() dockerCliPluginMetadata {
	return metadata
}

func runLoginex(opts loginOptions) error {
	var login loginInfo

	if err := verifyloginOptions(opts); err != nil {
		return err
	}

	if err := setloginInfo(opts, &login); err != nil {
		return err
	}

	var cmd string
	cmd = fmt.Sprintf("docker login --password-stdin --username %v %v", login.username, login.server)
	var in = login.password
	var buf bytes.Buffer

	if err := run(cmd, strings.NewReader(in), &buf); err != nil {
		return err
	}

	fmt.Printf(buf.String())

	return nil
}

func verifyloginOptions(opts loginOptions) error {
	if opts.serverAddress == "" {
		return fmt.Errorf("No server is specified in the argument.")
	}

	return nil
}

func setloginInfo(opts loginOptions, info *loginInfo) error {
	info.server = opts.serverAddress
	readNetrc()
	for _, v := range netrc {
		if v.machine == info.server {
			info.username = v.login
			info.password = v.password
			return nil
		}
	}

	return fmt.Errorf("No etnry in .netrc for the specified server %v.", opts.serverAddress)
}

func run(command string, r io.Reader, w io.Writer) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}
