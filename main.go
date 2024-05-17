package main

import (
	"fmt"
	"github.com/liberopassadorneto/quake/cmd"
	"github.com/liberopassadorneto/quake/logger"
	"os"
)

func main() {
	logger.Log.Info("Starting application")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
