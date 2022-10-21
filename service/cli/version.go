package cli

import (
	"gitee.com/up-zero/chan/define"
	"gitee.com/up-zero/chan/util"
)

// Version 打印系统版本
func Version() {
	util.Info("Version %v", define.Version)
}
