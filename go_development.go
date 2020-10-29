package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFileError
)

var Version = "0.0.0"

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	os.Args = args // 簡易テスト用
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Fprintf(c.outStream, "version: %s \n", Version)
		return 0
	} else {
		fmt.Fprintln(c.errStream, "バージョンオプションがありません")
		return 1
	}
}
