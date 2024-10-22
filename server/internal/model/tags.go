package model

type Tag struct {
	ID   uint   `json:"id", gorm:"primaryKey;autoIncrement"`
	name string `json:"name" gorm:"unique;not null"`
}
