package util

import (
	"gitee.com/up-zero/chan/define"
	"os"
	"os/user"
	"strconv"
)

func SaveProcessId() {
	pid := strconv.Itoa(os.Getpid())
	u, err := user.Current()
	if err != nil {
		panic("[GET CURRENT USER ERROR] : " + err.Error())
	}
	processFilePath := u.HomeDir + string(os.PathSeparator) + define.ProcessName
	if err = os.RemoveAll(processFilePath); err != nil {
		panic("[REMOVE PROCESS ERROR] : " + err.Error())
	}
	if err = os.WriteFile(processFilePath, []byte(pid), 066); err != nil {
		panic("[WRITE PROCESS ERROR] : " + err.Error())
	}
}
