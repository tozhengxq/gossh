gossh
====
不用登录远程服务器，在本地运行远程命令，轻松定制自己的部署和发布功能

### 安装/更新

```
go get -u github.com/tozhengxq/gossh
```

### 功能
简单几行代码就可以远程执行脚本或者系统命令，也可以进行批量操作

* 支持密码/公钥登录

* 支持非交互命令执行，如 sed，awk等

* 支持交互命令执行，并显示在本地终端， 如vim，top
* 支持配置文件模块和日志模块<可选>

### 使用
示例代码在example/gossh.go ，最复杂的代码，也不过如下：

```
go
package main

import (
	"fmt"
	"strconv"

	"github.com/tozhengxq/gossh/core/commands"
	"github.com/tozhengxq/gossh/core/conf"
	"github.com/tozhengxq/gossh/core/glog"
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
	sshConfig.InitConfig("./src/github.com/tozhengxq/gossh/gossh.cfg")

	//
	loger.Println("hello")

	//
	var (
		isUseKey   = true
		user       = sshConfig.Read("general", "user")
		host       = sshConfig.Read("general", "host")
		port, _    = strconv.Atoi(sshConfig.Read("general", "port"))
		connstring = sshConfig.Read("general", "testconn")
	)
	//初始化conn结构体
	cmdhandle := &commands.Conn{isUseKey, user, host, port, connstring}
	//普通命令执行
	cmdhandle.Runcmd("ls /tmp/;")
	cmdhandle.Runcmd("ls /tmp/;")
	//交互式命令执行
	cmdhandle.RunTerminal("top")

}

```
配置文件示例：

```
[general]
#登录用账户，地址和端口
user=root
host= 192.168.56.101
port = 22
#根据自己选择的登录方式来配置
#如果isUseKey为true，这里就配置公钥文件，为false这里就是密码
testconn = /Users/zhengxq/.ssh/id_rsa


[test1]
host= 192.168.56.101
port = 22


[test2]
num =  666
something  = wrong  #注释1
#fdfdfd = fdfdfd    注释整行
refer= refer       //注释3
```

