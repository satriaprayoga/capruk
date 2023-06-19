package capruk

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/capruk/config"
	"github.com/satriaprayoga/capruk/utils"
	"gorm.io/gorm"
)

var (
	appRoute *echo.Echo
	Config   *config.AppConfig
	DB       *gorm.DB
	rdb      *redis.Client

	RootPath string
)

func New(rootPath string) error {
	err := utils.InitPath(rootPath, "controllers", "dbdata", "redisdata", "usecases", "tmp", "models", "logs", "middlewares")
	if err != nil {
		return err
	}
	RootPath = rootPath
	Config = config.LoadConfig()
	Setup()
	DB = setupDB()
	appRoute = setupRoute()
	rdb = setupRedis()
	return nil
}

func Start() {
	startServer(Config.Server.HTTPPort)
}

func AutoMigrate(models ...interface{}) {
	log.Println("START AUTO MIGRATE")
	DB.AutoMigrate(models...)
	log.Println("FINISHING AUTO MIGRATE ")
}
