package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/elastos/Elastos.ELA.Monitor/models"
)

const (
	DefaultConfigFilename = "./config.json"
)

var (
	ConfigManager *models.Config
)

//type Config struct {
//	MonitorConfig Configuration `json:"Configuration"`
//}
//
//type LogConfig struct {
//	Path string 			`json:"Path"`
//	Level uint8    			`json:"Level"`
//	MaxPerLogSizeMb int64	`json:"MaxPerLogSizeMb"`
//	MaxLogsSizeMb int64		`json:"MaxLogsSizeMb"`
//}
//
//type Configuration struct {
//	Version  int	`json:"Version"`
//	Log *LogConfig	`json:"Log"`
//	Nodes *Nodes	`json:"Nodes"`
//}
//
//type Nodes struct {
//	MainChain *MainChain	`json:"MainChain"`
//}
//
//type MainChain struct {
//	Host string				`json:"Host"`
//	RpcPort int16			`json:"RpcPort"`
//	RestfulPort int16		`json:"RestfulPort"`
//	JarServer *JarServer	`json:"JarServer"`
//}
//
//type JarServer struct {
//	Url
//	Binary string `json:"Binary"`
//}
//
//type Url struct {
//	Host  string	`json:"Host"`
//	Port int16		`json:"Port"`
//}

func init() {
	file, err := ioutil.ReadFile(DefaultConfigFilename)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	configJson := &models.Config{
		MonitorConfig: models.Configuration{
			Version: 0,
		},
	}
	err = json.Unmarshal(file, configJson)
	if err != nil {
		fmt.Printf("Unmarshal json file err %v", err)
		os.Exit(1)
	}

	ConfigManager = configJson
}

