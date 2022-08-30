package models

type Personal struct {
	Amount    string `json:"full_name"`
	Type      string `json:"position"`
	PysysId   string `json:"working_day"`
	StudentId string `json:"phone"`
}

type PersonalResp struct {
	PersonalId string `json:"personal_id"`
	Amount     string `json:"full_name"`
	Type       string `json:"position"`
	PysysId    string `json:"working_day"`
	StudentId  string `json:"phone"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PersonalUpdate struct {
	Amount    string `json:"full_name"`
	Type      string `json:"position"`
	PysysId   string `json:"working_day"`
	StudentId string `json:"phone"`
}

type By struct {
	Name string `json:"name"`
}

type Personals struct {
	Personals []By `json:"personals"`
}

type PersonalsList struct {
	Personals []PersonalResp `json:"personals"`
	Count     int64          `json:"count"`
}
