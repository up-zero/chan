package service

import (
	"errors"
	"gitee.com/up-zero/chan/util"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	fileDir := "."
	filePath := fileDir + r.RequestURI
	index := strings.Index(filePath, "?")
	if index > 0 {
		filePath = filePath[:index]
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
			return
		}
		util.Error("ERROR : %v", err)
	}

	w.Write(data)
}
