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
		panic("[GET CURRENT USER ERROR] : " + err.Error())
	}
	processFilePath := u.HomeDir + string(os.PathSeparator) + define.ProcessName
	pid, err := ioutil.ReadFile(processFilePath)
	if err != nil {
		panic("[GET PROCESS FILE ERROR] : " + err.Error())
	}
	pidInt, _ := strconv.Atoi(string(pid))
	p, err := os.FindProcess(pidInt)
	if err != nil {
		panic("[GET PROCESS ERROR] : " + err.Error())
	}
	if err = p.Kill(); err != nil {
		panic("[KILL PROCESS ERROR] : " + err.Error())
	}
	util.Info("STOP SUCCESS")
}
