package storage

import (
	"fmt"
	"log"
	"main/pkg/globals"
	"os"
	"time"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Default() *DBConfig {
	return &DBConfig{
		Host:     globals.Env.DB_HOST,
		Port:     globals.Env.DB_PORT,
		Password: globals.Env.DB_PASS,
		User:     globals.Env.DB_USER,
		SSLMode:  globals.Env.DB_SSLMODE,
		DBName:   globals.Env.DB_NAME,
	}
}

func Connect(config *DBConfig) {
	// connStr := "postgresql://alfashop_owner:Ov9nAT1lNftb@ep-shy-paper-a2e1ldcp.eu-central-1.aws.neon.tech/alfashop?sslmode=require"
	// dbs, err := sql.Open("postgres", connStr)

	// if err != nil {
	// 	// handle errors
	// 	fmt.Println("Database connection error", err)
	// }

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
		),
		// Conn: connStr,
		PreferSimpleProtocol: true,
		// DisableForeignKeyConstraintWhenMigrating: true,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,         // Don't include params in the SQL log
				Colorful:                  true,          // Disable color
			},
		),
	})

	if err != nil {
		// handle errors
		fmt.Println("Database connection error", err)
	}

	DB = db
}
