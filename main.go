package main

import (
	"FurbotServer-Go/config"
	"FurbotServer-Go/extends/sql"
	"FurbotServer-Go/models"
	"FurbotServer-Go/router"
	"FurbotServer-Go/router/middleware"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	configFile = flag.String("c", "./config.yaml", "配置文件地址")
)

func main() {
	flag.Parse()
	if err := config.InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化数据库
	if err := sql.DB.InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 自动迁移
	sql.DB.Self.AutoMigrate(
		&models.FursuitTable{},
		&models.AuthTable{},
	)

	g := gin.New()
	router.LoadRouter(
		g,
		// 中间件
		middleware.Options,
		middleware.NoCache,
	)
	http.ListenAndServe(viper.GetString("server.addr")+":"+viper.GetString("server.port"), g)
}
