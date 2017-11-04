package main

import (
	"os"

	"github.com/go-martini/martini"
	flag "github.com/spf13/pflag"
)

const (
	//PORT 默认的端口号
	PORT string = "8080"
)

func main() {
	// 从系统环境变量中获得端口号
	port := os.Getenv("PORT")
	if len(port) == 0 { //如果没有PORT系统变量，则为默认端口号
		port = PORT
	}

	// 从命令行输入中获取端口号
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	/*
		启动一个服务器，并监听端口port，对路径为/+id，执行m.Get的回调函数
		martini使用详情请看https://github.com/go-martini/martini/blob/master/translations/README_zh_cn.md
	*/
	m := martini.Classic()
	m.Get("/:id", func(params martini.Params) string {
		return "Hello " + params["id"] + "\n"
	})
	m.RunOnAddr(":" + port)
}
