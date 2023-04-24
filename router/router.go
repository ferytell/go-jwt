package router

import (
	"github.com/ferytell/go-jwt/controllers"
	"github.com/ferytell/go-jwt/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	// add swagger
	// open here http://localhost:8000/swagger/index.html
	// @Security BearerAuth
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello bro, tester aja gw mah",
		})
	})

	userRouter := r.Group("/users")
	{
		// user Register
		userRouter.POST("/register", controllers.UserRegister)
		// user Login
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		// swagger
		photoRouter.Use(middlewares.Authentication())
		// photo Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// photo Edit
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// photo Create
		photoRouter.GET("/", controllers.GetPhoto)
		// photo Edit
		photoRouter.GET("/:photoId", controllers.GetPhotoByID)
		// photo Delete
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		// swagger
		socialMediaRouter.Use(middlewares.Authentication())
		// socmed Create
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		// socmed Edit
		socialMediaRouter.PUT("/:socmedId", middlewares.SocMedAuthorization(), controllers.UpdateSocialMedia)
		// socmed Create
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		// socmed Edit
		socialMediaRouter.GET("/:socmedId", controllers.GetSocialMediaById)
		// socmed Delete
		socialMediaRouter.DELETE("/:socmedId", controllers.DeleteSocialMedia)
	}

	commentsRouter := r.Group("/photos")
	{
		// swagger
		commentsRouter.Use(middlewares.Authentication())
		// socmed Create
		commentsRouter.POST("/:photoId/comments", controllers.CreateComments)
		// socmed Edit
		commentsRouter.PUT("/:photoId/comments/:commentId", middlewares.CommentsAuthorization(), controllers.UpdateComment)
		// socmed Create
		commentsRouter.GET("/comments", controllers.GetComments)
		// socmed Edit
		commentsRouter.GET("/:photoId/comments/:commentId", controllers.GetCommentByID)
		// socmed Delete
		commentsRouter.DELETE("/:photoId/comments/:commentId", controllers.DeleteComment)
	}

	return r
}
