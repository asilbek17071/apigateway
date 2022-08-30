package models

type Resp struct {
	Count string `json:"count"`
}

type RespList struct {
	Resp []Resp `json:"resp"`
}

type ProfitResp struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type ProfitRespList struct {
	Count      string       `json:"count"`
	ProfitResp []ProfitResp `json:"profit"`
}
