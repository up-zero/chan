package cli

func Restart(args []string) {
	Stop()
	Start(args)
}
