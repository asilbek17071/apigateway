package models

type Teacher struct {
	FullName       string `json:"full_name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Permission     string `json:"permission"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
	DirectionId    string `json:"direction_id"`
}

type TeacherResp struct {
	TeacherId      string `json:"teacher_id"`
	FullName       string `json:"full_name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Permission     string `json:"permission"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
	DirectionId    string `json:"direction_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type TeacherUpdate struct {
	FullName       string `json:"full_name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Permission     string `json:"permission"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
	DirectionId    string `json:"direction_id"`
}

type TeachersList struct {
	Teachers []TeacherResp `json:"teachers"`
	Count    int64         `json:"count"`
}
