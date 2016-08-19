package main

import (
	"fmt"
	"github.com/tozhengxq/gossh/core/commands"
	"github.com/tozhengxq/gossh/core/conf"
	"github.com/tozhengxq/gossh/core/glog"
	"strconv"
)

func main() {
	//日志和配置模块可以选用
	//
	//初始化log模块
	loger, ok := glog.Setloger("/tmp/", "test.log")
	if ok != nil {
		fmt.Println("init logfile fail: ", ok)
	}

	//
	//初始化配置模块
	sshConfig := new(conf.Config)
	sshConfig.InitConfig("./gossh.cfg")

	//
	loger.Println("hello")

	//
	var (
		user       = sshConfig.Read("general", "user")
		host       = sshConfig.Read("general", "host")
		port, _    = strconv.Atoi(sshConfig.Read("general", "port"))
		connstring = sshConfig.Read("general", "testconn")
	)
	//初始化conn结构体
	cmdhandle := &commands.Conn{true, user, host, port, connstring}
	//普通命令执行
	cmdhandle.Runcmd("ls /tmp/;")
	cmdhandle.Runcmd("ls /tmp/;")
	//交互式命令执行
	cmdhandle.RunTerminal("top")

}
