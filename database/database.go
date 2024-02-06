package database

import (
	"fmt"
	"os"

	"github.com/tcampbppu/server/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	// Database connection
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
}

var DB *gorm.DB

func NewDatabase() *Database {
	return &Database{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		Database:   os.Getenv("DB_DATABASE"),
	}
}

// DSN Database Source Name
func (d *Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}

// Connect Connect to the database
func (d *Database) Connect() error {
	var err error

	DB, err = gorm.Open(mysql.Open(d.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	d.Migrate()

	return nil
}

// Close Close the database connection
func (d *Database) Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (d *Database) Migrate() {
	DB.AutoMigrate(&models.Product{})
}
