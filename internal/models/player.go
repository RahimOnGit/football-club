package models

type Player struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"unique,size:255" binding:"required,alpha"`
	Position string `json:"position"`
	Club     string `json:"club" gorm:"default:18"  binding:"required"`
	Goals    int    `json:"goals" gorm:"default:0" binding:"required,numeric"`
}
