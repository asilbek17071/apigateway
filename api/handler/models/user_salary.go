package models

type UserSalary struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	UserId      string `json:"user_id"`
}

type UserSalaryResp struct {
	SalaryId    string `json:"salary_id"`
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	UserId      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserSalaryUpdate struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	UserId      string `json:"user_id"`
}

type UsersSalaryList struct {
	UsersSalary []UserSalaryResp `json:"users"`
	Count       int64            `json:"count"`
}
