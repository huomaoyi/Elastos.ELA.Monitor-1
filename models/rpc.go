package models

type ListProducers struct {
	Start uint16 `json:"start"`
	Limit uint16 `json:"limit"`
}

type ListProducersResponse struct {
	Producers	[]Producer	`json:"producers"`
	TotalVotes	string 	`json:"totalvotes"`
	TotalCounts int64 	`json:"totalcounts"`
}

type Height struct {
	Height int16	`json:"height"`
}

type RpcResponse struct {
	Error string 		`json:"error"`
	Id int16 			`json:"id"`
	JsonRpc string 		`json:"jsonrpc"`
	Result interface{} 	`json:"result"`
}
