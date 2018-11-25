package main

import (
	"hive/configs"
	"holiday/cmd"
	"holiday/pkg/logger"
)

var Env string

func main() {
	var start = logger.Info{
		Package: "main",
		Action:  "start server",
		Message: "start server by gin",
	}
	logger.InfoLog(start)
	if Env != "" {
		configs.ENV = Env
	} else {
		configs.ENV = "dev"
	}
	cmd.Execute()
}
