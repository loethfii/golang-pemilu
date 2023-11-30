package components

import (
	"gorm.io/gorm"
	"luthfi/pemilu/domain"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Article{},
		&domain.User{},
		&domain.Partai{},
		&domain.Paslon{},
		&domain.PaslonPartai{},
	)

}
