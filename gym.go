package gym

type Membership struct {
	Id           int    `json:"-" db:"id"`
	Title        string `json:"title" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Duration     int    `json:"duration" binding:"required"`
	InstructorId int
}

type Visit struct {
	Id      int `json:"-" db:"id"`
	InfoId  int
	Came_at string `json:"came_at"` //db:"came_at"
	Left_at string `json:"left_at"`
}
