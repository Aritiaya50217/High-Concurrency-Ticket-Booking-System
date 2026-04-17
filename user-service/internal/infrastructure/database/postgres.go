package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/config"
	"github.com/Aritiaya50217/High-Concurrency-Ticket-Booking-System/user-service/internal/infrastructure/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg.Database.Host == "" ||
		cfg.Database.Port == 0 ||
		cfg.Database.User == "" ||
		cfg.Database.Name == "" {

		return nil, fmt.Errorf("invalid database config: %+v", cfg.Database)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	log.Println("Connecting DB:", dsn) // debug

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %w", err)
	}

	// connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// ping check
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// migrate
	if err := db.AutoMigrate(
		&model.UsersModel{},
	); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	log.Println("database connected and migrated successfully")

	return db, nil
}
