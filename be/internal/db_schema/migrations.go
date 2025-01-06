package db_schema

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	. "github.com/seprich/keycape/internal/config"
	. "github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/util"
	"os"
	"time"
)

var dbHost = Config.String("db.host")
var dbPort = Config.Int("db.port")
var dbName = Config.String("db.dbName")
var dbUser = Config.String("db.user")
var dbPass = Config.String("db.passwd")

func RunMigrations() (util.Void, error) {
	err := tryMigration()
	if err != nil {
		for n := range 3 {
			Logger.Info(fmt.Sprintf("Migration attempt failed: %v :=> Sleep %d seconds and try again", err, n+1))
			time.Sleep(time.Duration(n+1) * time.Second)
			err = tryMigration()
			if err == nil {
				break
			}
		}
	}
	if err != nil {
		Logger.Error(fmt.Sprintf("Database migration error: %v", err))
		os.Exit(1)
	}

	Logger.Info("Database migrations completed")
	return util.Void{}, nil
}

func tryMigration() error {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	migrator, err1 := migrate.New("file://migrations", dbUrl)
	if err1 != nil {
		return err1
	}
	err2 := migrator.Up()
	if err2 != nil && !errors.Is(err2, migrate.ErrNoChange) {
		return err2
	}
	return nil
}
