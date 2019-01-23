package main

import (
	"github.com/elastos/Elastos.ELA.Monitor/config"
	"github.com/elastos/Elastos.ELA.Monitor/monitors"
	"github.com/elastos/Elastos.ELA.Monitor/nodes"
	"github.com/elastos/Elastos.ELA.Monitor/utility/log"
)

func init() {
	log.Init(
		config.ConfigManager.MonitorConfig.Log.Path,
		config.ConfigManager.MonitorConfig.Log.Level,
		config.ConfigManager.MonitorConfig.Log.MaxPerLogSizeMb,
		config.ConfigManager.MonitorConfig.Log.MaxLogsSizeMb,
	)
}

func main() {
	log.Info("Elastos Monitor")
	mainChain := config.ConfigManager.MonitorConfig.Nodes.MainChain

	elaNode := nodes.NewEla(mainChain)
	monitors.ProducerMonitor.Test(elaNode)
}
