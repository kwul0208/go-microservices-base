package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(50)" json:"name" binding:"required,min=3"`
	Description string `gorm:"type:text" json:"description" binding:"required,min=3"`
}
