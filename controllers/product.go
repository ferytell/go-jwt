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

// CreateTags		godoc
// @Summary			Create Product
// @Description		Save product data in Db.
// @Produce			application/json
// @Tags			Products
// @Success			200 {object} models.Product{}
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router			/products [post]
func CreateProduct(c *gin.Context) {
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

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body models.Product{} true  "Update tags"
// @Tags			Products
// @Produce			application/json
// @Success			200 {object} models.Product{}
// @Router			/products/{productId} [patch]
func UpdateProduct(c *gin.Context) {
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

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			Products
// @Security 		Bearer
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success			200 {obejct} []models.Product{}
// @Router			/products [get]
func GetProduct(c *gin.Context) {
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

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param 				Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				Products
// @Success				200 {object} []models.Product{productId}
// @Router				/products/{productId} [get]
func GetProductByID(c *gin.Context) {
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

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Param 			Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags			Products
// @Success			200 {object} []models.Product{}
// @Router			/products/{productId} [delete]
func DeleteProduct(c *gin.Context) {
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
