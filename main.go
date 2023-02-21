package main

import (
	config "go-test/Config"
	controller "go-test/Controllers"
	models "go-test/Models"
	repository "go-test/Repository"
	services "go-test/Services"
	middleware "go-test/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabase()
	db.AutoMigrate(&models.Token{}, &models.Customer{}, &models.Merchant{}, &models.UserLog{})

	repository := repository.NewRepository(db)
	logService := services.NewLogService(repository)
	tokenService := services.NewTokenService(repository)

	authService := services.NewAuthService(repository, logService, tokenService)
	authController := controller.NewAuthController(authService)

	r := gin.Default()
	public := r.Group("/api")

	public.POST("/register", authController.Register)
	public.POST("/login", authController.Login)
	protected := r.Group("/api/index")
	protected.Use(middleware.JwtAuthMiddleware())
	// protected.POST("/logout", authController.Logout)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
