package models

type Status struct {
	Id          string `json:"id"`
	Personal_id string `json:"personal_id"`
	Status      string `json:"status"`
}

type StatusReq struct {
	Status Status `json:"status"`
}

type StatusSpending struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type StatusSpendingReq struct {
	Status StatusSpending `json:"status"`
}
