# docker-loginex

`docker-loginex` is a Docker CLI plugins for slightly extending `docker login` command.

Read the .netrc and uses that information to log in, if there is an entry corresponding to the specified server.

If the entry for the specified server does not exist, it returns an error and exits.

See https://github.com/docker/cli/issues/1534 for the Docker CLI plugins.

## How-to-use

```
$ mkdir -p ~/.docker/cli-plugins
$ curl --output ~/.docker/cli-plugins/docker-loginex [Release URLs]
$ chmod +x ~/.docker/cli-plugins/docker-loginex
$ docker loginex --help
A Docker CLI plugins for slightly extending `docker login` command.
Log in to a Docker registry or cloud backend.
If no server is specified, the default is defined by the daemon.
See also help for `docker login`.

Usage:
  docker loginex [SERVER] [flags]

Flags:
  -h, --help   help for docker
```

## Reference

* CLI Plugins Design - docker/cli

    https://github.com/docker/cli/issues/1534

* cli/login.go at master - docker/cli

    https://github.com/docker/cli/blob/master/cli/command/registry/login.go

* src/cmd/go/internal/auth/netrc.go - The Go Programming Language

    https://golang.org/src/cmd/go/internal/auth/netrc.go

* jdxcode/netrc : Golang netrc parser

    https://github.com/jdxcode/netrc

* spf13/cobra : A Commander for modern Go CLI interactions

    https://github.com/spf13/cobra

* go-homedir はもう要らない | text.Baldanders.info

    https://text.baldanders.info/golang/no-need-go-homedir/

* Goツールのリリースにおけるバージョニングについて

    https://songmu.jp/riji/entry/2017-10-10-go-tool-version.html

## License

[MIT License](https://opensource.org/licenses/MIT).

However The process of retrieving data from .netrc was borrowed from `src/cmd/go/internal/auth/netrc.go`.

The license for this code part belongs to the BSD license of golang itself.
