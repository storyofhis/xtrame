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
	}

	router := gin.Default()
	config.GenerateJwtSignature()

	// users
	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	// category
	categoryRepo := gorm.NewCategoryRepo(db)
	categorySvc := services.NewCategorySvc(categoryRepo, userRepo)
	categoryHandler := controllers.NewCategoryController(categorySvc)

	// ticket
	ticketRepo := gorm.NewTicketRepo(db)
	tickerSvc := services.NewTicketSvc(ticketRepo, userRepo, categoryRepo)
	ticketHandler := controllers.NewTicketController(tickerSvc)

	// mail
	// config.ConnectMail()
	// to := []string{"azizi.maula@gmail.com", "maulaizzaazizi@gmail.com"}
	// cc := []string{"maulaizzazizi@gmail.com"}
	// subject := "Test Email"
	// message := "Hello"
	// email := services.NewEmailSvc(to, cc, subject, message)
	// email.ConnectEmail()

	app := httpserver.NewRouter(router, userHandler, categoryHandler, ticketHandler)
	PORT := os.Getenv("PORT")
	app.Start(":" + PORT)
}
