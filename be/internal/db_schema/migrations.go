package db_schema

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/samber/lo"
	. "github.com/seprich/keycape/internal/config"
	. "github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/util"
	"os"
)

var dbHost = Config.String("db.host")
var dbPort = Config.Int("db.port")
var dbName = Config.String("db.dbName")
var dbUser = Config.String("db.user")
var dbPass = Config.String("db.passwd")

func RunMigrations() (util.Void, error) {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error(fmt.Sprintf("Database migration error: %v", r))
			os.Exit(1)
		}
	}()
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	migrator := lo.Must1(migrate.New("file://migrations", dbUrl))

	err := migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	Logger.Info("Database migrations completed")
	return util.Void{}, nil
}
