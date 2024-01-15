package main

import (
	"time"

	"github.com/ponyjackal/go-gin-boilerplate/internal/adapters/database"
	"github.com/ponyjackal/go-gin-boilerplate/internal/adapters/database/migrations"
	"github.com/ponyjackal/go-gin-boilerplate/internal/app/routers"
	"github.com/ponyjackal/go-gin-boilerplate/pkg/config"
	"github.com/ponyjackal/go-gin-boilerplate/pkg/logger"

	"github.com/spf13/viper"
)

func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN := config.DbConfiguration()

	if err := database.DbConnection(masterDSN, replicaDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	//later separate migration
	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
