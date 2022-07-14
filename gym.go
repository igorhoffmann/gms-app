package gym

// import "errors"

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

// type UpdateListInput struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// }

// func (i UpdateListInput) Validate() error {
// 	if i.Title == nil && i.Description == nil {
// 		return errors.New("update structure has no values")
// 	}
// 	return nil
// }

// type UpdateItemInput struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// 	Done        *bool   `json:"done"`
// }

// func (i UpdateItemInput) Validate() error {
// 	if i.Title == nil && i.Description == nil && i.Done == nil {
// 		return errors.New("update structure has no values")
// 	}
// 	return nil
// }
