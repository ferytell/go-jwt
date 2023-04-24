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

// CreateSocialMedia	godoc
// @Summary				Create Social Media
// @Description			Save SocialMedia data in database it take userId who post it.
// @Produce				application/json
// @Tags				SocialMedia
// @Success				200 {object} models.SocialMedia{}
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router				/socialmedia [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, SocialMedia)
}

// UpdateSocialMedia	godoc
// @Summary				Update Social Media
// @Description			Update SocialMedia data in database it take userId who post it.
// @Produce				application/json
// @Tags				SocialMedia
// @Success				200 {object} models.SocialMedia{}
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router				/socialmedia{socmedId} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	socmed := models.SocialMedia{}

	socmedId, _ := strconv.Atoi(c.Param("socmedId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socmed)
	} else {
		c.ShouldBind(&socmed)
	}

	socmed.UserID = userID
	socmed.ID = uint(socmedId)

	// check if the social media exists and belongs to the user
	var existingSocmed models.SocialMedia
	err := db.Where("id = ? AND user_id = ?", socmedId, userID).First(&existingSocmed).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to update this social media",
		})
		return
	}

	// update the social media
	err = db.Model(&existingSocmed).Updates(models.SocialMedia{Name: socmed.Name, SocialMediaURL: socmed.SocialMediaURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, existingSocmed)
}

// GetSocialMedia		godoc
// @Summary				get Social Media
// @Description			get SocialMedia data in database it take userId who post it.
// @Produce				application/json
// @Tags				SocialMedia
// @Success				200 {object} models.SocialMedia{}
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router				/socialmedia [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()

	SocialMedia := []models.SocialMedia{}
	err := db.Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// GetSocialMediaById	godoc
// @Summary				Get Social Media
// @Description			Get SocialMedia data in database it take userId who post it.
// @Produce				application/json
// @Tags				SocialMedia
// @Success				200 {object} models.SocialMedia{}
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router				/socialmedia{socmedId} [get]
func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()

	socmedId, _ := strconv.Atoi(c.Param("socmedId"))

	SocialMedia := models.SocialMedia{}
	err := db.Where("id = ?", socmedId).First(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia	godoc
// @Summary				Delete Social Media
// @Description			Delete SocialMedia data in database it take userId who post it.
// @Produce				application/json
// @Tags				SocialMedia
// @Success				200 {string} hell yeah "Social Media deleted successfully"
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router				/socialmedia{socmedId} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	socmedID, _ := strconv.Atoi(c.Param("socmedId"))
	userID := uint(userData["id"].(float64))

	// Check if the social media record exists and is owned by the user
	SocialMedia := models.SocialMedia{}
	err := db.Where("id = ? AND user_id = ?", socmedID, userID).First(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete Social Media",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social Media deleted successfully",
	})
}
