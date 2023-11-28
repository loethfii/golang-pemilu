package domain

import (
	"context"
	"time"
)

type Partai struct {
	ID            int64     `gorm:"type:bigint; primaryKey; autoIncrement" json:"id"`
	Name          string    `gorm:"type:varchar(100)" json:"name"`
	Chairman      string    `gorm:"type:varchar(100)" json:"chairman"`
	VisionMission string    `gorm:"type:varchar(100)" json:"vision_mission"`
	Address       string    `gorm:"type:text" json:"address"`
	Image         string    `gorm:"type:varchar(300)" json:"image"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PartaiResponses struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Chairman      string `json:"chairman"`
	VisionMission string `json:"vision_mission"`
	Address       string `json:"address"`
	Image         string `json:"image"`
}

type PartaiResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Chairman      string `json:"chairman"`
	VisionMission string `json:"vision_mission"`
	Address       string `json:"address"`
	Image         string `json:"image"`
}

type PartaiRepository interface {
	Fetch(ctx context.Context) ([]Partai, error)
	GetByID(ctx context.Context, id int64) (Partai, error)
	Store(ctx context.Context, partai Partai) (Partai, error)
	Update(ctx context.Context, id int64, partai Partai) (Partai, error)
	Delete(ctx context.Context, id int64) error
}

type PartaiUseCase interface {
	Fetch() ([]PartaiResponses, error)
	GetByID(id int64) (PartaiResponse, error)
	Store(partai Partai) (Partai, error)
	Update(id int64, partai Partai) (Partai, error)
	Delete(id int64) error
}
