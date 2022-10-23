package cli

import "gitee.com/up-zero/chan/util"

func Conf() {
	cp, err := util.GetProcess()
	if err != nil {
		return
	}
	util.Info("ConfPath : %v", cp.Conf)
}
