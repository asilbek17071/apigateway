package models

type Accountant struct {
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
}

type AccountantResp struct {
	AccountantId         string `json:"acountant_id"`
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
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type AccountantUpdate struct {
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
}

type AccountantsList struct {
	Accountants []AccountantResp `json:"acountants"`
	Count int64      `json:"count"`
}
