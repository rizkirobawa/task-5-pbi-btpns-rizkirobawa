package main

import (
	"github.com/gin-gonic/gin"
	// database
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/database"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/models"

	// middlewares
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/middlewares"
	// controllers
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/controllers"
)

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// connect to database
	database.ConnectDatabase()

	// AutoMigrate will create tables based on the struct definitions
	database.DB.AutoMigrate(&models.User{}, &models.Photo{})

	// Route untuk endpoint registrasi user
	r.POST("/users/register", controllers.RegisterUser)

	// Route untuk endpoint login user
	r.POST("/users/login", controllers.LoginUser)

	// Route untuk endpoint memperbarui data user
	r.PUT("/users/:userId", middlewares.Req_Auth, controllers.UpdateUser)

	// Route untuk endpoint menghapus data user
	r.DELETE("/users/:userId", middlewares.Req_Auth, controllers.DeleteUser)

	// Route untuk endpoint menambahkan data foto
	r.POST("/photos", middlewares.Req_Auth, controllers.CreatePhoto)

	// Route untuk endpoint menampilkan data foto
	r.GET("/photos", middlewares.Req_Auth, controllers.ShowPhoto)

	// Route untuk endpoint memperbarui data foto
	r.PUT("/photos/:photoId", middlewares.Req_Auth, controllers.UpdatePhoto)

	// Route untuk endpoint menghapus data foto
	r.DELETE("/photos/:photoId", middlewares.Req_Auth, controllers.DeletePhoto)

	// Menjalankan server web
	r.Run()
}
