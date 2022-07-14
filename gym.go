package gym

type Membership struct {
	Id           int    `json:"id" db:"id"`
	Title        string `json:"title" db:"title" binding:"required"`
	Price        int    `json:"price" db:"price" binding:"required"`
	Duration     int    `json:"duration" db:"duration" binding:"required"`
	InstructorId int    `json:"instructor_id" db:"instructor_id"`
}

type Visit struct {
	Id      int    `json:"id" db:"id"`
	InfoId  int    `json:"info_id" binding:"required"`
	Came_at string `json:"came_at" db:"came_at"`
	Left_at string `json:"left_at" db:"left_at"`
}
