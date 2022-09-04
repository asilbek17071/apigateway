package models

type Finance struct {
	Category     string `json:"category"`
	PersonalId   string `json:"personal_id"`
	SalaryTypeId string `json:"salary_type_id"`
	Amount       string `json:"amount"`
	Comment      string `json:"comment"`
}

type FinResp struct {
	Status bool `json:"status"`
}

type FinanceResp struct {
	Name           string `json:"name"`
	Permission     string `json:"permission"`
	SalaryMonth    string `json:"salary_month"`
	Given          string `json:"given"`
	Debt           string `json:"debt"`
	Surcharge      string `json:"surcharge"`
	PrepaidExpense string `json:"prepaid_expense"`
	Comment        string `json:"comment"`
	Status         string `json:"status"`
}

type FinanceRespList struct {
	Finances []FinanceResp `json:"salary_personal"`
	Count    int64         `json:"count"`
}
