package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ferytell/go-jwt/database"
	"github.com/ferytell/go-jwt/helpers"
	"github.com/ferytell/go-jwt/models"
	"github.com/gin-gonic/gin"
)

// CreatePhoto		godoc
// @Summary			Create Photo
// @Description		Save Photo data in database it take userId who post it.
// @Produce			application/json
// @Tags			Photo
// @Success			200 {object} models.Photo{}
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router			/photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto			godoc
// @Summary				Update Photo
// @Description			Update Photo data on database.
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Photo
// @Produce				application/json
// @Success				200 {object} models.Photo{}
// @Router				/photo/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	// check if the photo exists and belongs to the user
	var existingPhoto models.Photo
	err := db.Where("id = ? AND user_id = ?", photoId, userID).First(&existingPhoto).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to update this photo",
		})
		return
	}

	// update the photo
	err = db.Model(&Photo).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetPhoto				godoc
// @Summary				get Photo
// @Description			Get All Photo data on database.
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Photo
// @Produce				application/json
// @Success				200 {object} models.Photo{}
// @Router				/photo [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()

	Photo := []models.Photo{}
	err := db.Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetPhoto				godoc
// @Summary				Get Photo by Id
// @Description			Get specific Photo data on database by Id.
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Photo
// @Produce				application/json
// @Success				200 {object} models.Photo{}
// @Router				/photo/{photoId} [get]
func GetPhotoByID(c *gin.Context) {
	db := database.GetDB()

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo := models.Photo{}
	err := db.Where("id = ?", photoId).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto			godoc
// @Summary				Delete Photo
// @Description			Delete Photo data on database based on inputed Id.
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Photo
// @Produce				application/json
// @Success				200 {string} string "hellyeah"
// @Router				/photo/{photoId} [delete]
func DeletePhoto(c *gin.Context) {

	tx := database.GetDB().Begin()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// check if the photo exists and belongs to the user
	var existingPhoto models.Photo
	err := tx.Where("id = ? AND user_id = ?", photoId, userID).First(&existingPhoto).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to delete this photo",
		})
		return
	}

	// delete the photo
	if err := tx.Delete(&existingPhoto).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete photo",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted successfully",
	})

	// db := database.GetDB()
	// photoId, _ := strconv.Atoi(c.Param("photoId"))
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	// // check if the photo exists and belongs to the user
	// var existingPhoto models.Photo
	// err := db.Where("id = ? AND user_id = ?", photoId, userID).First(&existingPhoto).Error
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":   "Unauthorized",
	// 		"message": "You are not authorized to delete this photo",
	// 	})
	// 	return
	// }

	// // hard delete = db.Unscoped().Where("id = ?", photoId).Delete(&models.Photo{})
	// // delete the photo
	// if err := db.Delete(&existingPhoto).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error":   "Internal Server Error",
	// 		"message": "Failed to delete photo",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Photo deleted successfully",
	// })
}
