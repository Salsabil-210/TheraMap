package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"theramap/models"
)

// NewConnection establishes a connection to PostgreSQL using GORM
func NewConnection(env *Env, logger *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Europe/Istanbul",
		env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort,
	)

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statements
		}),
		&gorm.Config{
			Logger: gormLogger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gormLogger.Config{
					SlowThreshold: time.Second,
					LogLevel:      gormLogger.Info,
					Colorful:      true,
				},
			),
		},
	)
	if err != nil {
		logger.Panic("failed connecting to database.", zap.Error(err))
	}

	// Configure connection pool
	dbConn, err := db.DB()
	if err != nil {
		logger.Panic("failed getting sql.DB from gorm", zap.Error(err))
	}

	dbConn.SetMaxOpenConns(150)
	dbConn.SetMaxIdleConns(100)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)
	dbConn.SetConnMaxLifetime(30 * time.Minute)

	logger.Info("database connection established.")
	return db
}

// Migrate executes SQL migrations and auto-migrates GORM models
func Migrate(db *gorm.DB, env *Env, logger *zap.Logger) {
	migrations := &migrate.FileMigrationSource{Dir: "migrations/"}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Panic("error getting db connection for migration", zap.Error(err))
	}

	logger.Info("running migrations...")

	_, err = migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		logger.Panic("error executing migrations", zap.Error(err))
	}

	// AutoMigrate GORM models
	if err := db.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.Emotion{},
		&models.Thought{},
		&models.MindMap{},
		&models.TherapistNote{},
		&models.AISuggestion{},
	); err != nil {
		logger.Panic("error auto-migrating models", zap.Error(err))
	}

	logger.Info("migration completed successfully.")
}
