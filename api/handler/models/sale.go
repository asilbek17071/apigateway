package models

type Sale struct {
	Name    string `json:"name"`
	Amount  string `json:"amount"`
	Comment string `json:"comment"`
}

type SaleResp struct {
	SaleId    string `json:"sale_id"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SaleUpdate struct {
	Name    string `json:"name"`
	Amount  string `json:"amount"`
	Comment string `json:"comment"`
}

type SalesList struct {
	Sales []SaleResp `json:"sales"`
	Count int64      `json:"count"`
}
