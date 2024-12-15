package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name" binding:"required,min=3"`
	Email    string `gorm:"type:varchar(50)" json:"email" binding:"required,min=3"`
	Password string `gorm:"type:varchar(100)" json:"password" binding:"required,min=3"`
	Role     string `gorm:"type:varchar(50)" json:"role" binding:"required,min=3"`
}
