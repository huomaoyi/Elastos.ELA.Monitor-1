package nodes

import (
	"github.com/elastos/Elastos.ELA.Monitor/models"
	"github.com/elastos/Elastos.ELA.Monitor/servers"
)

type Ela struct {
	Name 		string
	Host 		string
	Restful 	*servers.Restful
	Rpc 		*servers.Rpc
	JarServer 	*servers.JarServer
}

func NewEla(mainChain *models.MainChain) *Ela {
	return &Ela{
		Name: "Ela node",
		Host: (*mainChain).Host,
		Restful: servers.NewRestful((*mainChain).Host, (*mainChain).RestfulPort),
		Rpc: servers.NewRpc((*mainChain).Host, (*mainChain).RpcPort),
		JarServer: servers.NewJarServer((*mainChain).JarServer.Binary, (*mainChain).JarServer.Host, (*mainChain).JarServer.Port),
	}
}