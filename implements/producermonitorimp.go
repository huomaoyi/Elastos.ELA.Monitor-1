package implements

import (
	"fmt"
	"github.com/elastos/Elastos.ELA.Monitor/servers"
	"github.com/elastos/Elastos.ELA.Monitor/utility/log"
)

var ProducerMonitorImp *producerMonitorImplements

type producerMonitorImplements struct {
}

func (pmi *producerMonitorImplements) Test(rpc *servers.Rpc) {
	producers, totalVotes, totalCounts, err := rpc.GetListProducers(0, 100)
	 if err != nil{
	 	log.Error(fmt.Sprintf("call failed! %+v", err))
	 }

	log.Debug(fmt.Sprintf("producers is %+v", producers[0].NickName))
	log.Debug(fmt.Sprintf("totalVotes is %+v", totalVotes))
	log.Debug(fmt.Sprintf("totalCounts is %+v", totalCounts))
}

func (pmi *producerMonitorImplements) ReadBlock(rpc *servers.Rpc) {
	producers, totalVotes, totalCounts, err := rpc.GetListProducers(0, 100)
	if err != nil{
		log.Error(fmt.Sprintf("call failed! %+v", err))
	}

	log.Debug(fmt.Sprintf("producers is %+v", producers[0].NickName))
	log.Debug(fmt.Sprintf("totalVotes is %+v", totalVotes))
	log.Debug(fmt.Sprintf("totalCounts is %+v", totalCounts))
}