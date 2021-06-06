package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

	in := "Hello, docker-loginex"
	var buf bytes.Buffer

	if err := run("docker version", strings.NewReader(in), &buf); err != nil {
		return err
	}

	fmt.Println("Hello, docker-loginex")
	fmt.Printf("  docker version: [%v]", buf.String())
	fmt.Printf("  loginOptions: %v\n", opts)
	fmt.Printf("  loginInfo: %v\n", login)

	return nil
}

func verifyloginOptions(opts loginOptions) error {
	if opts.password != "" {
		fmt.Println("WARNING! Using --password via the CLI is insecure. Use --password-stdin.")
		if opts.passwordStdin {
			return fmt.Errorf("--password and --password-stdin are mutually exclusive")
		}
	}

	if opts.passwordStdin {
		if opts.username == "" {
			return fmt.Errorf("Must provide --username with --password-stdin")
		}
	}

	return nil
}

func setloginInfo(opts loginOptions, info *loginInfo) error {
	if opts.passwordStdin {
		contents, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		info.password = strings.TrimSuffix(string(contents), "\n")
		info.password = strings.TrimSuffix(info.password, "\r")
	}

	return nil
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
