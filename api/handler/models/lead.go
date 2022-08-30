package models

type Lead struct {
	FullName    string `json:"full_name"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Source      string `json:"source"`
	Comment     string `json:"comment"`
	UserId      string `json:"user_id"`
	DirectionId string `json:"direction_id"`
}

type LPhon struct {
	Whose  string `json:"whose"`
	Phone  string `json:"phone"`
	LeadAd string `json:"lead_id"`
}

type LPhone struct {
	Id        string `json:"id"`
	Whose     string `json:"whose"`
	Phone     string `json:"phone"`
	LeadId    string `json:"lead_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LeadResp struct {
	LeadId      string `json:"lead_id"`
	FullName    string `json:"full_name"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Source      string `json:"source"`
	Comment     string `json:"comment"`
	UserId      string `json:"user_id"`
	DirectionId string `json:"direction_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type LeadUpdate struct {
	FullName    string `json:"full_name"`
	Birthday    string `json:"birthday"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Source      string `json:"source"`
	Comment     string `json:"comment"`
	UserId      string `json:"user_id"`
	DirectionId string `json:"direction_id"`
}

type LeadsList struct {
	Leads []LeadResp `json:"leads"`
	Count int64      `json:"count"`
}
