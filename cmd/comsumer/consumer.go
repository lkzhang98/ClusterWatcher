package main

import (
	"ClusterWatcher/internal/probe"
	"os"
)

func main() {
	command := probe.NewConsumerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
