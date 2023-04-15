package middlewares

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ferytell/go-jwt/database"
	"github.com/ferytell/go-jwt/models"

	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":    "Bad Request",
				"messages": "invailid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorization",
				"message": "you are not allowed to acces this data",
			})

			return
		}

		c.Next()
	}
}
