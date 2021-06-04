package main

import (
	"os"

	"github.com/yusys-cloud/fastdfsbeat/cmd"

	_ "github.com/yusys-cloud/fastdfsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
