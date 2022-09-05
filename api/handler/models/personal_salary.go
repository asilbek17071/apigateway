package models

type PersonalSalary struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	PersonalId  string `json:"personal_id"`
}

type PersonalSalaryResp struct {
	SalaryId    string `json:"salary_id"`
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	PersonalId  string `json:"personal_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PersonalSalaryUpdate struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	PersonalId  string `json:"personal_id"`
}

type PersonalsSalaryList struct {
	PersonalsSalary []PersonalSalaryResp `json:"personals"`
	Count           int64                `json:"count"`
}
