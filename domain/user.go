package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `gorm:"type:bigint; primaryKey; autoIncrement" json:"id"`
	FUllName  string    `gorm:"type:varchar(100)" json:"full_name"`
	Address   string    `gorm:"type:text" json:"address"`
	Gender    string    `gorm:"type:varchar(100)" json:"gender"`
	Username  string    `gorm:"type:varchar(100)" json:"username"`
	Password  string    `gorm:"type:varchar(100)" json:"password"`
	Role      string    `gorm:"type:varchar(100)" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	RegisterUser(ctx context.Context, user User) (User, error)
	LoginUser(ctx context.Context, username, password string) (User, error)
}

type UserUseCase interface {
	RegisterUser(user User) (User, error)
	LoginUser(username, password string) (User, error)
}
