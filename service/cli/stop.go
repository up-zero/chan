package cli

import (
	"gitee.com/up-zero/chan/define"
	"gitee.com/up-zero/chan/util"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

func Stop() {
	u, err := user.Current()
	if err != nil {
		util.Info("[GET CURRENT USER ERROR] : " + err.Error())
		return
	}
	processFilePath := u.HomeDir + string(os.PathSeparator) + define.ProcessName
	pid, err := ioutil.ReadFile(processFilePath)
	if err != nil {
		util.Info("[GET PROCESS FILE ERROR] : " + err.Error())
		return
	}
	pidInt, _ := strconv.Atoi(string(pid))
	p, err := os.FindProcess(pidInt)
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
