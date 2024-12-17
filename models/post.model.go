package models

// Post model
type Post struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint // Foreign key
}

// TableName specifies the table name for User
func (Post) TableName() string {
	return "account_post"
}
