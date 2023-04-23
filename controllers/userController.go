package controllers

import (
	"net/http"

	"github.com/ferytell/go-jwt/database"
	"github.com/ferytell/go-jwt/helpers"
	"github.com/ferytell/go-jwt/models"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// CreateTags		godoc
// @Summary			Create new User
// @Description		Register new user data in Db.
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Security 		Bearer
// @Produce			application/json
// @Tags			User
// @Success			200 {object} []models.User{}
// @Router			/users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"user_name": User.UserName,
		"age":       User.Age,
	})
}

// CreateTags		godoc
// @Summary			User Login
// @Description		New user Login and verived based on data in Db.
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce			application/json
// @Tags			User
// @Success			200 {object} models.User{}
// @Router			/users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invailid email/Password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invailid email/Password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
