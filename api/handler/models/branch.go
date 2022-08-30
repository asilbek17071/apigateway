package models

type Branch struct {
	Name string `json:"name"`
}

type BranchResp struct {
	BranchId  string `json:"branch_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BranchUpdate struct {
	Name string `json:"name"`
}

type BranchsList struct {
	Branchs []BranchResp `json:"branchs"`
	Count   int64        `json:"count"`
}

type WithRoom struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
}

type WithRooms struct {
	BranchWithRoom []WithRoom `json:"branch_with_room"`
	Count          int64      `json:"count"`
}
