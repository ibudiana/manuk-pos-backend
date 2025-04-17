package sync

import (
	"time"

	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
)

// SyncLog represents the log of a sync operation performed on the system
type SyncLog struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceID string `gorm:"type:varchar(100);not null" json:"device_id" binding:"required"`
	UserID   int    `gorm:"not null" json:"user_id" binding:"required"`
	BranchID int    `gorm:"not null" json:"branch_id" binding:"required"`

	SyncType     string    `gorm:"type:varchar(100);not null" json:"sync_type" binding:"required"`
	SyncStatus   string    `gorm:"type:varchar(100);default:'in_progress'" json:"sync_status"`
	StartTime    time.Time `gorm:"autoCreateTime" json:"start_time"`
	EndTime      string    `gorm:"type:varchar(100)" json:"end_time"`
	DataCount    int       `gorm:"default:0" json:"data_count"`
	ErrorMessage string    `gorm:"type:varchar(100)" json:"error_message"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi dengan tabel lain
	Branch *store.Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User   *user.User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// SyncDetail represents details about individual records synced during a sync operation
type SyncDetail struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	SyncLogID int `gorm:"not null" json:"sync_id" binding:"required"`
	// RecordID  int `gorm:"not null" json:"record_id"`

	TableName    string `gorm:"type:varchar(100);not null" json:"table_name" binding:"required"`
	SyncAction   string `gorm:"type:varchar(100);not null" json:"sync_action" binding:"required"`
	SyncStatus   string `gorm:"type:varchar(50);default:'pending'" json:"sync_status"`
	ErrorMessage string `gorm:"type:varchar(100)" json:"error_message"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relasi dengan tabel lain
	SyncLog SyncLog `gorm:"foreignKey:SyncLogID" json:"sync_log"`
}
