package models

type SpendingType struct {
	Name string `json:"name"`
}

type SpendingTypeResp struct {
	SpendingTypeId string `json:"s_type_id"`
	Name           string `json:"name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type SpendingTypeUpdate struct {
	Name string `json:"name"`
}

type SpendingTypesList struct {
	Spendings []SpendingTypeResp `json:"spending_tayps"`
	Count     int64              `json:"count"`
}
