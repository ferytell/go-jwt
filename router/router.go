package router

import (
	"github.com/ferytell/go-jwt/controllers"
	"github.com/ferytell/go-jwt/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title myGram API
// @version 1.0.0
// @description This is just sample
// @termsOfService http://swagger.io/terms
// @contact.name API suppoer
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	// register router

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/users")
	{
		// user Register
		userRouter.POST("/register", controllers.UserRegister)
		// user Login
		userRouter.POST("/login", controllers.UserLogin)

	}

	productRouter := r.Group("/products")

	{

		// swagger
		productRouter.Use(middlewares.Authentication())
		// product Create
		productRouter.POST("/", controllers.CreateProduct)
		// product Edit
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		// product Create
		productRouter.GET("/", controllers.GetProduct)
		// product Edit
		productRouter.GET("/:productId", controllers.GetProductByID)
		// product Delete
		productRouter.DELETE("/:productId", controllers.DeleteProduct)
	}

	photoRouter := r.Group("/photos")

	{

		// swagger
		photoRouter.Use(middlewares.Authentication())
		// product Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// product Edit
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// product Create
		photoRouter.GET("/", controllers.GetPhoto)
		// product Edit
		photoRouter.GET("/:photoId", controllers.GetPhotoByID)
		// product Delete
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r

}
