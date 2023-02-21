package Config

import (
	fmt "fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "Server=Fadhlih\\SQLEXPRESS;Database=golang_mnc;Trusted_Connection=True;TrustServerCertificate=True;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed connect to database: %v", err))
	}
	return db
}
