package models

type Ball struct {
	Ball      string `json:"ball"`
	WithDate  string `json:"with_date"`
	StudentId string `json:"student_id"`
	GroupId   string `json:"group_id"`
}

type BallResp struct {
	BallId    string `json:"ball_id"`
	Ball      string `json:"ball"`
	WithDate  string `json:"with_date"`
	StudentId string `json:"student_id"`
	GroupId   string `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BallUpdate struct {
	Ball      string `json:"ball"`
	WithDate  string `json:"with_date"`
	StudentId string `json:"student_id"`
	GroupId   string `json:"group_id"`
}

type BallsList struct {
	Balls []BallResp `json:"balls"`
	Count int64      `json:"count"`
}
