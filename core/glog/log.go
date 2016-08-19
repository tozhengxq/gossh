package glog

import (
	"fmt"
	"log"
	"os"
)

func Setloger(logpath string, logname string) (*log.Logger, error) {
	logfile, err := os.OpenFile(logpath+logname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
		//系统退出，返回code，0表示正常退出，非0表示异常退出
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
	return logger, nil
}
