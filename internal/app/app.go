package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"shorter/config"
	"shorter/pkg/httpserver"
	"shorter/pkg/postgres"
)

type App struct {
	Name         string
	Version      string
	Http         any //todo
	FileLogger   *log.Logger
	RabbitLogger *log.Logger
	Db           *postgres.Postgres
}

var appInstance *App

func GetAppInstance() *App {
	if appInstance != nil {
		return appInstance
	}
	appInstance = &App{}
	return appInstance
}

func (app *App) Init(cfg *config.Config) error {
	app.Name = cfg.Name
	app.Version = cfg.Version
	err := app.initFileLogger(cfg.Log.Path, cfg.Log.Level)
	if err != nil {
		return err
	}
	err = app.initDataBase(cfg.PG.DSN)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) initFileLogger(path string, lvl string) error {
	fileStream, err := os.OpenFile("../../log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer fileStream.Close()

	app.FileLogger = log.New(fileStream, lvl, log.LstdFlags)

	return nil
}

func (app *App) initDataBase(dsn string) error {
	var err error
	app.Db, err = postgres.New(dsn)
	if err != nil {
		return err
	}
	defer app.Db.Close()
	return nil
}

//todo here
func (app *App) initHttpServer() error {
	handler := gin.New()
	app.Http = httpserver.New()
	return nil
}

func (app *App) Run() {

}
