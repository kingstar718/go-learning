package main

import (
	log "github.com/sirupsen/logrus"
)

/**
初始化项目：go mod init 项目名
添加代理：go env -w GOPROXY=https://goproxy.cn
运行项目：go run 文件名
*/

func LogrusDemo(s string) {
	log.WithFields(log.Fields{
		"animals": "walrus",
	}).Info("A walrus appears")
	log.Info("log info s: {}", s)
	log.Debug("log debug s: {}")
}
