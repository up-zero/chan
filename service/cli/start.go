package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func Start(args []string) {
	go func(args []string) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)
		select {
		case <-ch:
			fmt.Println("Backend Running")
			var cmd *exec.Cmd
			if len(args) >= 4 && strings.ToUpper(args[2]) == "-T" {
				cmd = exec.Command("chan", "run", "-t", args[3])
			} else {
				cmd = exec.Command("chan", "run")
			}
			if err := cmd.Start(); err != nil {
				panic("[START ERROR] : " + err.Error())
			}
			os.Exit(1)
		}
	}(args)
	Run(args)
}
