package capruk

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func setupDB() *gorm.DB {
	now := time.Now()
	var (
		Conn *gorm.DB
		err  error
	)
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Config.Database.Host,
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Name,
		Config.Database.Port)
	fmt.Printf("%s", connectionString)

	Conn, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Config.Database.TablePrefix,
			SingularTable: true,
		},
		Logger: dbLogger,
	})
	if err != nil {
		log.Fatalf("connection.setup err : %v", err)

	}
	sqlDB, err := Conn.DB()
	if err != nil {
		log.Fatalf("connection.setup DB err : %v", err)

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
	return Conn

}
