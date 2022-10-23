package util

import (
	"encoding/json"
	"gitee.com/up-zero/chan/define"
	"io/ioutil"
	"os"
	"os/user"
)

func SaveProcess(confPath string) {
	cp := &define.ChanProcess{
		Pid:  os.Getpid(),
		Conf: confPath,
	}
	data, err := json.Marshal(cp)
	if err != nil {
		panic("[MARSHAL ERROR] : " + err.Error())
	}
	u, err := user.Current()
	if err != nil {
		panic("[GET CURRENT USER ERROR] : " + err.Error())
	}
	processFilePath := u.HomeDir + string(os.PathSeparator) + define.ProcessName
	if err = os.RemoveAll(processFilePath); err != nil {
		panic("[REMOVE PROCESS ERROR] : " + err.Error())
	}
	if err = os.WriteFile(processFilePath, data, 066); err != nil {
		panic("[WRITE PROCESS ERROR] : " + err.Error())
	}
}

func GetProcess() (*define.ChanProcess, error) {
	u, err := user.Current()
	if err != nil {
		Info("[GET CURRENT USER ERROR] : " + err.Error())
		return nil, err
	}
	processFilePath := u.HomeDir + string(os.PathSeparator) + define.ProcessName
	p, err := ioutil.ReadFile(processFilePath)
	if err != nil {
		Info("[GET PROCESS FILE ERROR] : " + err.Error())
		return nil, err
	}
	cp := new(define.ChanProcess)
	err = json.Unmarshal(p, cp)
	if err != nil {
		Info("[UNMARSHAL PROCESS FILE ERROR] : " + err.Error())
		return nil, err
	}
	return cp, nil
}
