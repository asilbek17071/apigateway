package models

type Target struct {
	Name string `json:"name"`
}

type TargetResp struct {
	TargetId  string `json:"target_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TargetUpdate struct {
	Name string `json:"name"`
}

type TargetsList struct {
	Targets []TargetResp `json:"target"`
	Count   int64        `json:"count"`
}
