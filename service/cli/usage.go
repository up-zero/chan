package cli

import (
	"gitee.com/up-zero/chan/util"
	"github.com/olekukonko/tablewriter"
	"os"
)

// CommandUsage 命令行使用帮助
func CommandUsage() {
	Version()
	util.Info("Usage : ")
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Name", "Usage", "Param"})
	t.Append([]string{"version", "chan version", "no"})
	t.Append([]string{"run", "chan run", "-t config file path"})
	t.Append([]string{"start", "chan start , can backend run", "-t config file path"})
	t.Render()
}
