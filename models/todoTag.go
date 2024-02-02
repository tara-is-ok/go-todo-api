package models

type TodoTag struct {
	TodoId int `json:"todo_id" gorm:"primaryKey"`
	TagId  int `json:"tag_id" gorm:"primaryKey"`
}
