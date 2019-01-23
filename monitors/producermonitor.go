package monitors

import (
	"github.com/elastos/Elastos.ELA.Monitor/implements"
	"github.com/elastos/Elastos.ELA.Monitor/nodes"
)

var ProducerMonitor *producerMonitor

type producerMonitor struct {
	Name string
}

func (producerMonitor *producerMonitor) Test(node *nodes.Ela) {
	implements.ProducerMonitorImp.Test(node.Rpc)
}