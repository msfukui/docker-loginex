package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

type dockerCliPluginMetadata struct {
	SchemaVersion    string
	Vendor           string
	Version          string
	Revision         string
	ShortDescription string
	URL              string
}

var metadata dockerCliPluginMetadata = dockerCliPluginMetadata{
	SchemaVersion:    "0.1.0",
	Vendor:           "msfukui",
	Version:          version,
	Revision:         revision,
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
	Use:     "docker-loginex [SERVER]",
	Version: version,
	Short:   "A Docker CLI plugins for slightly extending `docker login` command.",
	Long: "A Docker CLI plugins for slightly extending `docker login` command.\n" +
		"Log in to a Docker registry or cloud backend.\n" +
		"If no server is specified, the default is defined by the daemon.\n" +
		"See also help for `docker login`.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			if args[0] == "docker-cli-plugin-metadata" {
				// Return metadata as a requirement of Docker CLI pligins.
				j, _ := json.Marshal(getDockerCliPluginMetadata())
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
	v := "docker-loginex version: " + version + " (rev: " + revision + ")\n"
	rootCmd.SetVersionTemplate(v)
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

	out, err := runDockerLogin(login)
	if err != nil {
		return err
	}

	fmt.Print(out)

	return nil
}

func verifyloginOptions(opts loginOptions) error {
	if opts.serverAddress == "" {
		return fmt.Errorf("no server is specified in the argument")
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

	return fmt.Errorf("no etnry in .netrc for the specified server %v", opts.serverAddress)
}

func runDockerLogin(login loginInfo) (string, error) {
	var buf bytes.Buffer
	var cmd *exec.Cmd

	cmd = exec.Command("docker", "login", "--password-stdin", "--username", login.username, login.server)

	cmd.Stderr = os.Stderr
	cmd.Stdout = &buf
	cmd.Stdin = strings.NewReader(login.password)

	if err := cmd.Run(); err != nil {
		return buf.String(), err
	}
	return buf.String(), nil
}
