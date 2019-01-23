package models

type HttpData struct {
	Method string `json:"method"`
	Params interface{} `json:"params"`
}
