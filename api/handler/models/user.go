package models

type User struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Permission     string `json:"permission"`
	Password       string `json:"password"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
}

type UserResp struct {
	UserId         string `json:"user_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Permission     string `json:"permission"`
	Password       string `json:"password"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Gender         string `json:"gender"`
	Phone1         string `json:"phone1"`
	Phone2         string `json:"phone2"`
	Password       string `json:"password"`
	Permission     string `json:"permission"`
	WorkingDay     string `json:"working_day"`
	Address        string `json:"address"`
	PassportSeries string `json:"passport_series"`
	PassportImage1 string `json:"passport_image1"`
	PassportImage2 string `json:"passport_image2"`
	Image          string `json:"image"`
	BranchId       string `json:"branch_id"`
}

type UsersList struct {
	Users []UserResp `json:"users"`
	Count int64      `json:"count"`
}
