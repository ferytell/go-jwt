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

// CreateComment		godoc
// @Summary			Create Comment
// @Description		Create Comment on photo it take userId who post it.
// @Produce			application/json
// @Tags			Comments
// @Success			200 {object} models.Comments{}
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router			/photos/{photoId}/comments [post]
func CreateComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comments{}
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64) // get photo ID from URL parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// Retrieve photo and user from database
	photo := models.Photo{}
	err = db.First(&photo, photoID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Photo not found",
		})
		return
	}
	user := models.User{}
	err = db.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "User not found",
		})
		return
	}

	Comment.PhotoID = photo.ID
	Comment.UserID = user.ID

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment		godoc
// @Summary			Update Comments
// @Description		Update tags data.
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags			Comments
// @Produce			application/json
// @Success			200 {object} models.Comments{}
// @Router			/photos/{photoId}/comments/{productId} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64) // get comment ID from URL parameter
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Retrieve comment from database
	comment := models.Comments{}
	err = db.First(&comment, commentID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Comment not found",
		})
		return
	}

	// Check if the user ID matches the creator of the comment
	userID := uint(userData["id"].(float64))
	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to edit this comment",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err = db.Debug().Save(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// FindAllComments 		godoc
// @Summary				Get All Comments.
// @Description			Return list of Comments.
// @Tags				Comments
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success				200 {obejct} []models.Comments{}
// @Router				/photos/{photoId}/comments [get]
func GetComments(c *gin.Context) {
	db := database.GetDB()

	comments := []models.Comments{}
	err := db.Where("user_id IS NOT NULL AND photo_id IS NOT NULL").Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, comments)
}

// FindByIdComments 	godoc
// @Summary				Get Single Comment by id.
// @Security 			Bearer
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				Comments
// @Success				200 {object} models.Comments{}
// @Router				/photos/{photoId}/comments/{productId} [get]
func GetCommentByID(c *gin.Context) {
	db := database.GetDB()

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	comments := models.Comments{}
	err := db.Where("id = ?", photoId).First(&comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// DeleteComments	godoc
// @Summary			Delete Comments
// @Description		Remove Comments data on Photo by id.
// @Produce			application/json
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags			Comments
// @Success 		200 {message} string "Comment deleted successfully"
// @Failure 		401 {error} string "Unauthorized"
// @Failure 		404 {error} string "Comment not found"
// @Failure 		500 {string} string "Internal Server Error"
// @Router			/photos/{photoId}/comments/{commentId} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(c.Param("commentId"))

	// Retrieve comment from database
	var comment models.Comments
	if err := db.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Comment not found",
		})
		return
	}

	// Check if the user ID matches the creator of the comment
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	if comment.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to delete this comment",
		})
		return
	}

	// Delete the comment
	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete comment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}
