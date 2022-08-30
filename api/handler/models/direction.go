package models

type Direction struct {
	Name         string `json:"name"`
	Active       string `json:"active"`
	Duration     string `json:"duration"`
	Month_amount string `json:"month_amount"`
	Full_amount  string `json:"full_amount"`
}

type DirectionResp struct {
	DirectionId  string `json:"direction_id"`
	Name         string `json:"name"`
	Active       string `json:"active"`
	Duration     string `json:"duration"`
	Month_amount string `json:"month_amount"`
	Full_amount  string `json:"full_amount"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type DirectionUpdate struct {
	Name         string `json:"name"`
	Active       string `json:"active"`
	Duration     string `json:"duration"`
	Month_amount string `json:"month_amount"`
	Full_amount  string `json:"full_amount"`
}

type DirectionsList struct {
	Directions []DirectionResp `json:"directions"`
	Count      int64           `json:"count"`
}

type Empty struct{}
