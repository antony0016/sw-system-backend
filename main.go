package main

import (
	"fmt"
	"github.com/antony0016/sw-system-backend/middlewares"
	"github.com/antony0016/sw-system-backend/routers"
	"github.com/antony0016/sw-system-backend/services/pgdb"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Initial database instance and auto migrate
	pgdb.Init()
	router := gin.Default()
	router.Use(middlewares.Cors)
	routers.InitRouter(router.Group("/"))
	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}
}
