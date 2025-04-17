package inventory

import "time"

type Category struct {
	ID       int  `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID *int `json:"parent_id"` // Nullable jika tidak punya parent

	Name        string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Code        string `gorm:"type:varchar(50);unique" json:"code"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Level       int    `gorm:"default:1" json:"level"`
	Path        string `gorm:"type:varchar(100)" json:"path"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Parent   *Category  `gorm:"foreignKey:ParentID" json:"parent"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children"`
}
