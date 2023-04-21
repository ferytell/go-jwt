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

// CreatePhoto godoc
// @Summary Create Photo
// @Description Post new Photo based on userId
// @Tags photo
// @Accept json
// @Produce json
// @Param models.photo body models.photo true "create photo"
// @Succes 200 {object} models.photo
// @Router /products [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product
// @Tags product
// @Accept json
// @Produce json
// @Param models.product body models.product true "update product"
// @Succes 200 {object} models.product
// @Router /products/{productId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	// err := db.Table("Product").Where("id = ?", c.param("productId")).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error
	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Product)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()

	products := []models.Product{}
	err := db.Find(&products).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetPhotoByID(c *gin.Context) {
	db := database.GetDB()

	productId, _ := strconv.Atoi(c.Param("productId"))

	product := models.Product{}
	err := db.Where("id = ?", productId).First(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	productId, _ := strconv.Atoi(c.Param("productId"))

	// Check if the product exists
	var product models.Product
	if err := db.First(&product, productId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Product not found",
		})
		return
	}

	// Delete the product
	if err := db.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
