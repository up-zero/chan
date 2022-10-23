package main

import (
	"gitee.com/up-zero/chan/service/cli"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		cli.CommandUsage()
		return
	}

	switch strings.ToUpper(args[1]) {
	case "VERSION":
		cli.Version()
	case "RUN":
		cli.Run(args)
	case "START":
		cli.Start(args)
	}
}
