package models

type SalaryType struct {
	Name string `json:"name"`
}

type SalaryTypeResp struct {
	SalaryTypeId string `json:"salary_type_id"`
	Name         string `json:"name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type SalaryTypeUpdate struct {
	Name string `json:"name"`
}

type SalaryTypesList struct {
	SalaryTypes []SalaryTypeResp `json:"salary_tayps"`
	Count       int64            `json:"count"`
}
