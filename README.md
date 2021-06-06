# docker-loginex

`docker-loginex` is a Docker CLI plugins for slightly extending `docker login` command.

See https://github.com/docker/cli/issues/1534 for the Docker CLI plugins.

## How-to-use

```
$ mkdir -p ~/.docker/cli-plugins
$ curl --output ~/.docker/cli-plugins/docker-loginex [Release URLs]
$ chmod +x ~/.docker/cli-plugins/docker-loginex
$ docker loginex --help
Log in to a Docker registry or cloud backend.
See also help for `docker login`.

Usage:
  docker loginex [OPTIONS] [SERVER] [flags]

Flags:
  -h, --help              Help for login
  -p, --password string   password
      --password-stdin    Take the password from stdin
  -u, --username string   username
```

## Feature

* Add the ability to use `.netrc` to `docker login`.

## ToDo

* [ ] Implementation of `ghcr.io` subcommand.

## Reference

* CLI Plugins Design - docker/cli

    https://github.com/docker/cli/issues/1534

* src/cmd/go/internal/auth/netrc.go - The Go Programming Language

    https://golang.org/src/cmd/go/internal/auth/netrc.go

* jdxcode/netrc : Golang netrc parser

    https://github.com/jdxcode/netrc

* spf13/cobra : A Commander for modern Go CLI interactions

    https://github.com/spf13/cobra

* go-homedir はもう要らない | text.Baldanders.info

    https://text.baldanders.info/golang/no-need-go-homedir/
