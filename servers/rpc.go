package servers

import (
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ELA.Monitor/models"
	"github.com/elastos/Elastos.ELA.Monitor/utility/http"
	"github.com/elastos/Elastos.ELA.Monitor/utility/log"
	"github.com/goinggo/mapstructure"
)

type Rpc struct {
	Host string
	Port int16
	Url string
}

func NewRpc(host string, port int16) *Rpc {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return &Rpc{host, port, url}
}

func (rpc *Rpc) GetListProducers(start, limit uint16) ([]models.Producer, string, float64, error) {
	data := models.ListProducers{start, limit}
	response, err :=rpc.callAndReadRpc("listproducers", data)
	if err != nil {
		return nil, "", 0, err
	}

	tmp := models.ListProducersResponse{}
	mapstructure.Decode(response.Result, &tmp)

	resultRaw, _ := json.Marshal(response.Result)
	resultData := make(map[string]interface{})
	err = json.Unmarshal(resultRaw, &resultData)
	if err != nil {
		log.Error(fmt.Sprintf("unmarshal resultData failed! %+v", err))
	}

	producersRaw, _ := json.Marshal(resultData["producers"])
	var producerData []models.Producer
	err = json.Unmarshal(producersRaw, &producerData)
	if err != nil {
		log.Error(fmt.Sprintf("unmarshal producerData failed! %+v", err))
	}

	totalVotes := resultData["totalvotes"].(string)
	totalCounts := resultData["totalcounts"].(float64)

	return producerData, totalVotes, totalCounts, err
}

func (rpc *Rpc) rpcPost(url, method string, params interface{}) ([]byte, error) {
	httpData := &models.HttpData{method, params}
	data, err := json.Marshal(httpData)
	if err != nil {
		log.Warn(fmt.Sprintf("json marshal failed: %+v", data))
		log.Warn(fmt.Sprintf("Error: %+v", err))
		return nil, err
	}

	sendData := string(data)
	log.Info(fmt.Sprintf("call %s with %s", method, sendData))
	return http.Post(url, sendData)
}

func (rpc *Rpc) callAndReadRpc(method string, params interface{}) (*models.RpcResponse, error) {
	response, err := rpc.rpcPost(rpc.Url, method, params)
	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("response is %+v", string(response)))
	rpcResponse := &models.RpcResponse{}
	err = json.Unmarshal(response, rpcResponse)
	if err != nil {
		log.Error(fmt.Printf("Unmarshal json file err %v", err))
		log.Error(fmt.Printf("respone is %v", string(response)))
	}
	return rpcResponse, err
}