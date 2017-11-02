package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matejl/cron-manager/controller"
	"github.com/spf13/viper"
	"flag"
	"fmt"
	"github.com/matejl/cron-manager/model"
)

type params struct {
	configPath *string
}

func main() {

	p := params{}
	p.configPath = flag.String("configPath", "config", "Config path relative to current directory")

	loadConfig(&p)
	err := checkConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")
	router.GET("/index", controller.Index)
	router.Run(viper.GetString("app.listen.addr") + ":" + viper.GetString("app.listen.port"))
}

func loadConfig(p *params) {

	viper.SetConfigName("cron-manager")
	viper.AddConfigPath(*p.configPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func checkConfig() error {
	db := model.GetConnection()
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}