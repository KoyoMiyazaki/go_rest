package models

type User struct {
	ID       uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"` // 読み取り、書き込みOK
	Token    string `gorm:"-" json:"token,omitempty"`
}
