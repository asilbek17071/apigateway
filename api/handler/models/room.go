package models

type Room struct {
	Name     string `json:"name"`
	BranchId string `json:"branch_id"`
}

type RoomResp struct {
	RoomId    string `json:"room_id"`
	Name      string `json:"name"`
	BranchId  string `json:"branch_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RoomUpdate struct {
	Name     string `json:"name"`
	BranchId string `json:"branch_id"`
}

type RoomsList struct {
	Rooms []RoomResp `json:"rooms"`
	Count int64      `json:"count"`
}
