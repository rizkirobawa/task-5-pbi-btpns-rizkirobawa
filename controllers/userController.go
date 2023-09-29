package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/app"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/database"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/helpers"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/models"
)

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RegisterUser(c *gin.Context) {
	var user_reg app.AuthorizedRegister
	user_reg.Id = uuid.New().String()
	if err := c.ShouldBindJSON(&user_reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, user_reg); err != nil {
		return
	}

	// Enkripsi kata sandi
	hashedPass := helpers.EncryptPassword(c, user_reg.Password)
	if hashedPass == "" {
		return
	}

	// Membuat objek pengguna baru
	user := models.User{
		Id:        user_reg.Id,
		Username:  user_reg.Username,
		Email:     user_reg.Email,
		Password:  hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Membuat pengguna baru di database
	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user. please change the email or the password",
		})
		return
	}

	// Menyiapkan respons dengan "message" dan "data" fields
	response := ResponseData{
		Message: "Successfully registered",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
		},
	}

	// Mengembalikan respons yang telah disiapkan
	c.JSON(http.StatusOK, response)
}

func LoginUser(c *gin.Context) {
	var user_login app.AuthorizedLogin
	var user models.User

	// Mengaitkan data login dengan konteks
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, user_login); err != nil {
		return
	}

	// Mencari pengguna berdasarkan alamat email
	database.DB.First(&user, "email = ?", user_login.Email)

	// Memeriksa alamat email dan kata sandi yang tidak valid
	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Memeriksa kata sandi
	if err := helpers.CheckPassword(user.Password, user_login.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Membuat token
	tokenStr, err := helpers.InitToken(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Menyiapkan respons dengan token dan data pengguna
	response := ResponseData{
		Message: "Successfully login",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"token":    tokenStr,
		},
	}

	// Set token sebagai cookie
	expTime := 86400 // detik (1 hari)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, expTime, "", "", false, true)

	// Mengembalikan respons yang telah disiapkan
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	type UserInput struct {
		Id       string `valid:"required" json:"id"`
		Username string `json:"username"`
		Email    string `valid:"email" json:"email"`
		Password string `json:"password"`
	}

	var user_input UserInput
	var user models.User
	user_input.Id = c.Param("userId")

	if err := c.ShouldBindJSON(&user_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data yang diterima
	if err := helpers.Validation(c, user_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enkripsi kata sandi jika
	if user_input.Password != "" {
		hashedPassword := helpers.EncryptPassword(c, user_input.Password)
		if hashedPassword == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
			return
		}
		user_input.Password = hashedPassword
	}

	// Memeriksa apakah user dengan ID yang diberikan ada di database
	if result := database.DB.First(&user, "id = ?", user_input.Id); result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Memperbarui data pengguna di database
	if user_input.Username != "" {
		user.Username = user_input.Username
	}

	if user_input.Email != "" {
		user.Email = user_input.Email
	}

	if user_input.Password != "" {
		user.Password = user_input.Password
	}

	if result := database.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal melakukan update. Silakan coba lagi"})
		return
	}

	// Menyiapkan respons dengan "message" dan "data" fields
	response := ResponseData{
		Message: "Berhasil mengubah data.",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
		},
	}

	// Return response
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	var photo models.Photo
	var deletedPhotos bool
	userid := c.Param("userId")

	// Menghapus foto-foto yang terkait dengan pengguna
	if err := database.DB.Where("userid = ?", userid).First(&photo).Error; err == nil {
		database.DB.Delete(&models.Photo{}, "userid = ?", userid)
		deletedPhotos = !deletedPhotos
	}

	// Menghapus pengguna berdasarkan ID
	if err := database.DB.Where("id = ?", userid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	database.DB.Where("id = ?", userid).Delete(&user)

	// Menyiapkan respons
	if deletedPhotos {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted photos and account",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted",
	})
}
