package models

type TeacherSalary struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	TeacherId   string `json:"teacher_id"`
}

type TeacherSalaryResp struct {
	SalaryId    string `json:"salary_id"`
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	TeacherId   string `json:"teacher_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TeacherSalaryUpdate struct {
	Amount      string `json:"amount"`
	Comment     string `json:"comment"`
	DateOfBegin string `json:"date_of_begin"`
	TeacherId   string `json:"teacher_id"`
}

type TeachersSalaryList struct {
	TeachersSalary []TeacherSalaryResp `json:"teachers"`
	Count          int64               `json:"count"`
}
