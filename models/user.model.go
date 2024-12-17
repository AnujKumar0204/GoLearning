package models

type User struct {
	ID uint `gorm:"primaryKey"`
	// Name  string `gorm:"size:100"`
	// Email string `gorm:"unique;size:100"`
	Firstname string `gorm:"size:255"`
	Lastname  string `gorm:"size:255"`
	Username  string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	Posts     []Post `gorm:"foreignKey:UserID"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "account_useraccount"
}
