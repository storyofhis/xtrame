package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/storyofhis/xtrame/config"
	"github.com/storyofhis/xtrame/httpserver"
	"github.com/storyofhis/xtrame/httpserver/controllers"
	"github.com/storyofhis/xtrame/httpserver/repository/gorm"
	"github.com/storyofhis/xtrame/httpserver/services"
	// "gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return
	}
}
func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
		log.Println(err)
		return
	}

	router := gin.Default()
	config.GenerateJwtSignature()

	// users
	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	app := httpserver.NewRouter(router, userHandler)
	PORT := os.Getenv("PORT")
	app.Start(":" + PORT)
}
