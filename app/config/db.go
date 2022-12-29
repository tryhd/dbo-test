package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load enf file!")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

// type Migrator interface {
// 	// AutoMigrate
// 	AutoMigrate(dst ...interface{}) error

// 	// Database
// 	CurrentDatabase() string
// 	FullDataTypeOf(*schema.Field) clause.Expr

// 	// Tables
// 	CreateTable(dst ...interface{}) error
// 	DropTable(dst ...interface{}) error
// 	HasTable(dst interface{}) bool
// 	RenameTable(oldName, newName interface{}) error

// 	// Columns
// 	AddColumn(dst interface{}, field string) error
// 	DropColumn(dst interface{}, field string) error
// 	AlterColumn(dst interface{}, field string) error
// 	HasColumn(dst interface{}, field string) bool
// 	RenameColumn(dst interface{}, oldName, field string) error
// 	MigrateColumn(dst interface{}, field *schema.Field, columnType *sql.ColumnType) error
// 	ColumnTypes(dst interface{}) ([]*sql.ColumnType, error)

// 	// Constraints
// 	CreateConstraint(dst interface{}, name string) error
// 	DropConstraint(dst interface{}, name string) error
// 	HasConstraint(dst interface{}, name string) bool

// 	// Indexes
// 	CreateIndex(dst interface{}, name string) error
// 	DropIndex(dst interface{}, name string) error
// 	HasIndex(dst interface{}, name string) bool
// 	RenameIndex(dst interface{}, oldName, newName string) error
// }
