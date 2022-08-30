package models

type Spending struct {
	Amount       string `json:"amount"`
	Comment      string `json:"comment"`
	Status       string `json:"status"`
	BranchId     string `json:"branch_id"`
	SpendingType string `json:"s_type_id"`
	PysysId      string `json:"pysys_id"`
}

type SpendingResp struct {
	SpendingId   string `json:"spending_id"`
	Amount       string `json:"amount"`
	Comment      string `json:"comment"`
	Status       string `json:"status"`
	BranchId     string `json:"branch_id"`
	SpendingType string `json:"s_type_id"`
	PysysId      string `json:"pysys_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type SpendingUpdate struct {
	Amount       string `json:"amount"`
	Comment      string `json:"comment"`
	Status       string `json:"status"`
	BranchId     string `json:"branch_id"`
	SpendingType string `json:"s_type_id"`
	PysysId      string `json:"pysys_id"`
}

type SpendingsList struct {
	Spendings []SpendingResp `json:"spendings"`
	Count     int64          `json:"count"`
}
