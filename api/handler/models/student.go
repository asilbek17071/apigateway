package models

type Student struct {
	FullName       string `json:"full_name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	GroupId        string `json:"group_id"`
}

type SPhon struct {
	Whose     string `json:"whose"`
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
}

type SPhone struct {
	Id        string `json:"id"`
	Whose     string `json:"whose"`
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type LoginStudentReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type StudentResp struct {
	StudentId      string `json:"student_id"`
	FullName       string `json:"full_name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	GroupId        string `json:"group_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type StudentUpdate struct {
	FullName       string `json:"full_name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	GroupId        string `json:"group_id"`
}

type StudentsList struct {
	Students []StudentResp `json:"students"`
	Count    int64         `json:"count"`
}
