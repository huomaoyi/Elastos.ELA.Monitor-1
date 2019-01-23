package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/elastos/Elastos.ELA.Monitor/utility/log"
)

func Post(url, data string) ([]byte, error){
	response, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		log.Error(fmt.Sprintf("http post failed: %+v", err))
		log.Error(fmt.Sprintf("http url is %s", url))
		log.Error(fmt.Sprintf("data is %+v", data))
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("read response failed! %+v", err))
		log.Error(fmt.Sprintf("response is %+v", response))
	}

	return body, err
}
