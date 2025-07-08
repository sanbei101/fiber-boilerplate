package bootstrap

import (
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
		db.Log.Error().Err(err).Msg("An unknown error occurred when connecting to the database!")
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
