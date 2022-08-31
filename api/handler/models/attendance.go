package models

type Attendance struct {
	Participate bool   `json:"participate"`
	Comment     string `json:"comment"`
	WithDate    string `json:"with_date"`
	StudentId   string `json:"student_id"`
	GroupId     string `json:"group_id"`
}

type AttendanceResp struct {
	AttendanceId string `json:"attendance_id"`
	Participate  bool   `json:"participate"`
	Comment      string `json:"comment"`
	WithDate     string `json:"with_date"`
	StudentId    string `json:"student_id"`
	GroupId      string `json:"group_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type AttendanceUpdate struct {
	Participate bool   `json:"participate"`
	Comment     string `json:"comment"`
	WithDate    string `json:"with_date"`
	StudentId   string `json:"student_id"`
	GroupId     string `json:"group_id"`
}

type AttendancesList struct {
	Attendances []AttendanceResp `json:"attendances"`
	Count       int64            `json:"count"`
}

type ParticipateCount struct {
	Count string `json:"count"`
}

type Participate struct {
	GroupId          string             `json:"group_id"`
	GroupName        string             `json:"group_name"`
	TeacherId        string             `json:"teacher_id"`
	TeacherName      string             `json:"teacher_name"`
	WithDays         string             `json:"with_days"`
	BeginTime        string             `json:"begin_time"`
	CloseTime        string             `json:"close_time"`
	ParticipateCount []ParticipateCount `json:"count"`
}
