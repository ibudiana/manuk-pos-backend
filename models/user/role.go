package user

type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(50);unique;not null" json:"role_name" binding:"required"`
	Description string `gorm:"type:text" json:"description"`
}
