package servers

import (
	"fmt"
	"github.com/elastos/Elastos.ELA.Monitor/config"
)

var (
	JarServerApi *JarServer
)

type JarServer struct {
	Binary string
	Host string
	Port int16
	Url string
}

func init() {
	jarServerEndpoint := config.ConfigManager.MonitorConfig.Nodes.MainChain.JarServer
	url := fmt.Sprintf("http://%s:%d", (*jarServerEndpoint).Host, (*jarServerEndpoint).Port)
	JarServerApi = &JarServer{"",(*jarServerEndpoint).Host, (*jarServerEndpoint).Port, url}
}

func NewJarServer(binary, host string, port int16) *JarServer {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return &JarServer{binary, host, port, url}
}