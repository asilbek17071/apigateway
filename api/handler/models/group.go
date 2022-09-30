package models

import "time"

type Group struct {
	Name        string    `json:"name"`
	DateOfBegin string    `json:"date_of_begin"`
	DateOfClose string    `json:"date_of_close"`
	WithDays    []int     `json:"with_days"`
	BeginTime   time.Time `json:"begin_time"`
	CloseTime   time.Time `json:"close_time"`
	Lesson      int       `json:"lesson"`
	Comment     string    `json:"comment"`
	Active      bool      `json:"active"`
	DirectionId string    `json:"direction_id"`
	TeacherId   string    `json:"teacher_id"`
	BranchId    string    `json:"branch_id"`
	RoomId      string    `json:"room_id"`
}

type GroupResp struct {
	GroupId     string    `json:"group_id"`
	Name        string    `json:"name"`
	DateOfBegin string    `json:"date_of_begin"`
	DateOfClose string    `json:"date_of_close"`
	WithDays    []int     `json:"with_days"`
	BeginTime   time.Time `json:"begin_time"`
	CloseTime   time.Time `json:"close_time"`
	Lesson      int       `json:"lesson"`
	Comment     string    `json:"comment"`
	Active      bool      `json:"active"`
	DirectionId string    `json:"direction_id"`
	TeacherId   string    `json:"teacher_id"`
	BranchId    string    `json:"branch_id"`
	RoomId      string    `json:"room_id"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type GroupUpdate struct {
	Name        string    `json:"name"`
	DateOfBegin string    `json:"date_of_begin"`
	DateOfClose string    `json:"date_of_close"`
	WithDays    []int     `json:"with_days"`
	BeginTime   time.Time `json:"begin_time"`
	CloseTime   time.Time `json:"close_time"`
	Lesson      int       `json:"lesson"`
	Comment     string    `json:"comment"`
	Active      bool      `json:"active"`
	DirectionId string    `json:"direction_id"`
	TeacherId   string    `json:"teacher_id"`
	BranchId    string    `json:"branch_id"`
	RoomId      string    `json:"room_id"`
}

type GroupsList struct {
	Groups []GroupResp `json:"groups"`
	Count  int64       `json:"count"`
}

type At struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ExA struct {
	WithDate    string `json:"with_date"`
	Comment     string `json:"comment"`
	Participate string `json:"participate"`
	StudentId   string `json:"student_id"`
	GroupId     string `json:"group_id"`
}

type Attendancelist struct {
	Name    string `json:"name"`
	Student []At   `json:"student"`
	Example []ExA  `json:"example"`
}

type ExB struct {
	WithDate  string `json:"with_date"`
	Ball      string `json:"ball"`
	StudentId string `json:"student_id"`
	GroupId   string `json:"group_id"`
}

type Ballist struct {
	Name    string `json:"name"`
	Student []At   `json:"student"`
	Example []ExB  `json:"example"`
}
