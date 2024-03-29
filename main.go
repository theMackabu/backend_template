package main

import (
	"flag"
	"fmt"
	"log"

	"backend_template/config"
	"backend_template/routes"
   "backend_template/database"

	"github.com/labstack/echo/v4/middleware"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.yml", "path to the config file")
	flag.Parse()
}

func main() {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
      
	router := routes.NewApi(cfg)
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[1;34m[${time_rfc3339}]\033[0m \033[1;36m${method}\033[0m ${uri} ${status} \033[1;33m(${latency_human})\033[0m \033[1;31m${error}\033[0m\n",
	}))

   database.CreateDatabase(cfg.Database)
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", cfg.Server.Port)))
}
