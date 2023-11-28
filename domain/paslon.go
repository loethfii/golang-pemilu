package domain

import (
	"context"
	"time"
)

type Paslon struct {
	ID            int64     `gorm:"type:bigint; primaryKey; autoIncrement" json:"id"`
	Name          string    `gorm:"type:varchar(100)" json:"name"`
	SerialNumber  string    `gorm:"type:varchar(100)" json:"serial_number"`
	VisionMission string    `gorm:"type:varchar(100)" json:"vision_mission"`
	Image         string    `gorm:"type:varchar(300)" json:"image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaslonResponses struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	SerialNumber  string `json:"serial_number"`
	VisionMission string `json:"vision_mission"`
	Image         string `json:"image"`
}

type PaslonResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	SerialNumber  string `json:"serial_number"`
	VisionMission string `json:"vision_mission"`
	Image         string `json:"image"`
}

type PaslonRepository interface {
	Fetch(ctx context.Context) ([]Paslon, error)
	GetByID(ctx context.Context, id int64) (Paslon, error)
	Store(ctx context.Context, paslon Paslon) (Paslon, error)
	Update(ctx context.Context, id int64, paslon Paslon) (Paslon, error)
	Delete(ctx context.Context, id int64) error
}

type PaslonUseCase interface {
	Fetch() ([]PaslonResponses, error)
	GetByID(id int64) (PaslonResponse, error)
	Store(paslon Paslon) (Paslon, error)
	Update(id int64, paslon Paslon) (Paslon, error)
	Delete(id int64) error
}
