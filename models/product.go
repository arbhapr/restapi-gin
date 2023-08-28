package models

type Product struct {
	Id          int64 `gorm:"primaryKey" json:"id"`
	Name        int64 `gorm:"type:varchar(255)" json:"name"`
	Description int64 `gorm:"type:text" json:"description"`
}
