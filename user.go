package gym

type User struct {
	Id               int    `json:"-"` //db:"id"
	First_Name       string `json:"first_name" binding:"required"`
	Last_Name        string `json:"last_name" binding:"required"`
	Middle_Name      string `json:"middle_name"`
	Relationship     string `json:"relationship" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	Date_of_birth    string `json:"date_of_birth" binding:"required"`
	Date_of_registry string `json:"date_of_registry"` //db:"date_of_registry"
}

type SysUser struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type Member struct {
	UserId       int    `json:"-"`
	MembershipId int    `json:"-"`
	Expires_at   string `json:"expires_at" binding:"required"`
}

type Instructor struct {
	UserId    int    `json:"-"`
	Hire_date string `json:"hire_date" db:"hire_date"`
	Salary    int    `json:"salary" binding:"required"`
}
