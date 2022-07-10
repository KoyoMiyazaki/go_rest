package models

type Book struct {
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `gorm:"not null" json:"-"` // 外部キー
	User        User   `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
