package models

type Blog struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Topic  string `json:"topic"`
	UserID uint   `json:"-"`
}
