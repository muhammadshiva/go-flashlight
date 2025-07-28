package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Username  string    `gorm:"type:varchar(20);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(100);not null"`
	Phone     string    `gorm:"type:varchar(15);not null"`
	RoleID    uint      `gorm:"type:uint;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Role      Role      `gorm:"foreignKey:role_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
