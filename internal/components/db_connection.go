package components

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"luthfi/pemilu/internal/config"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"user=%s password=%s "+
			"dbname=%s port=%s "+
			"sslmode=disable "+
			"timezone=Asia/Jakarta",
		cnf.Database.Host,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
		cnf.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = migrate(db)
	if err != nil {
		panic(err)
	}

	return db
}
