# nba_statistics_cli
## Install GO
Make sure you have a working Go environment. Go version 1.10+ is supported. [See the install instructions for Go.](https://golang.org/doc/install)

Make sure your PATH includes GOPATH
```
export PATH=$PATH:$GOPATH/bin
```

## Build CLI
A simple, fast GOLANG package is being used:
https://github.com/urfave/cli

Complete the installation of all packages, run:
```
go install
```
nba_statistics_cli will be available on your machine.

## Run test
Under tests folder, run
```
go test
```
