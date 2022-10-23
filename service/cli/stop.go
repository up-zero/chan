package cli

import (
	"gitee.com/up-zero/chan/util"
	"os"
)

func Stop() {
	cp, err := util.GetProcess()
	if err != nil {
		return
	}
	p, err := os.FindProcess(cp.Pid)
	if err != nil {
		util.Info("[GET PROCESS ERROR] : " + err.Error())
		return
	}
	if err = p.Kill(); err != nil {
		util.Info("[KILL PROCESS ERROR] : " + err.Error())
		return
	}
	util.Info("STOP SUCCESS")
}
