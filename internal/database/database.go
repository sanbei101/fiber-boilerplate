package database

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/model"
	"github.com/efectn/fiber-boilerplate/utils/config"
	"github.com/rs/zerolog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
	Log  zerolog.Logger
	Cfg  *config.Config
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}
	return db
}

func (db *Database) ConnectDatabase() {
	gormDB, err := gorm.Open(postgres.Open(db.Cfg.DB.Postgres.DSN), &gorm.Config{})
	if err != nil {
		db.Log.Error().Err(err)
	} else {
		db.Log.Info().Msg("Connected to the database successfully!")
	}
	db.Gorm = gormDB
}

func (db *Database) ShutdownDatabase() {
	sqlDB, err := db.Gorm.DB()
	if err != nil {
		db.Log.Error().Err(err).Msg("Failed to get sql.DB from gorm.DB!")
		return
	}
	if err := sqlDB.Close(); err != nil {
		db.Log.Error().Err(err).Msg("An unknown error occurred when shutting down the database!")
	}
}

func (db *Database) Migrate() {
	if err := db.Gorm.AutoMigrate(model.Article{}); err != nil {
		db.Log.Error().Err(err)
	} else {
		db.Log.Info().Msg("Database migrated successfully!")
	}
}
