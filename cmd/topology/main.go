package main

import (
	"ClusterWatcher/internal/topology"
	"os"
)

func main() {
	command := topology.NewTopologyCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
