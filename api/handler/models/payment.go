package models

type Payment struct {
	Amount      string `json:"amount"`
	Type        string `json:"type"`
	PysysId     string `json:"pysys_id"`
	StudentId   string `json:"student_id"`
	DirectionId string `json:"direction_id"`
	GroupId     string `json:"group_id"`
}

type PaymentResp struct {
	PaymentId   string `json:"payment_id"`
	Amount      string `json:"amount"`
	Type        string `json:"type"`
	PysysId     string `json:"pysys_id"`
	StudentId   string `json:"student_id"`
	DirectionId string `json:"direction_id"`
	GroupId     string `json:"group_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PaymentUpdate struct {
	Amount      string `json:"amount"`
	Type        string `json:"type"`
	PysysId     string `json:"pysys_id"`
	StudentId   string `json:"student_id"`
	DirectionId string `json:"direction_id"`
	GroupId     string `json:"group_id"`
}

type PaymentsList struct {
	Payments []PaymentResp `json:"payments"`
	Count    int64         `json:"count"`
}
