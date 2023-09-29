package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/database"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/helpers"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/models"
)

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	photo.Id = uuid.New().String()
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Create(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to add item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Photo has been created successfully", "data": photo})
}

func ShowPhoto(c *gin.Context) {
	var photos []models.Photo
	userid, _ := c.Get("userid")

	database.DB.Where("userid = ?", userid).Find(&photos)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": photos})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	photo.Id = c.Param("photoId")
	userid, _ := c.Get("userid")
	if userid == nil {
		// Lakukan sesuatu, misalnya kembalikan respons error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is nil"})
		return
	}
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Melakukan operasi pembaruan
	result := database.DB.Model(&photo).Where("userid = ?", photo.Userid).Updates(&photo)

	// Memeriksa apakah pembaruan berhasil
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}

	// Mengambil data foto yang diperbarui dari database
	database.DB.First(&photo, "id = ?", photo.Id)

	c.JSON(http.StatusOK, gin.H{"message": "data has been updated", "data": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	photo.Id = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := database.DB.Where("userid = ?", photo.Userid).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Where("userid = ?", photo.Userid).Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": "Deleted successfully"})
}
