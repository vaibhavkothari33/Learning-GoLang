package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavkothari33/backendapi/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/verify-email", controllers.VerifyEmail)
		auth.POST("/login", controllers.Login)
		auth.POST("/forgot-password", controllers.ForgotPassword)
		auth.POST("/reset-password", controllers.ResetPassword)
		// auth.POST("/logout", controllers.Register)

		// public := api.Group("/blogs")
		// {
		// 	// public.GET("/",controllers.GetBlogs)
		// }

		router.GET("/heath", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "OK",
				"message": "API is running",
			})
		})
	}
}

// ```go
// package routes

// import (
//     "blog-backend/controllers"
//     "blog-backend/middleware"

//     "github.com/gin-gonic/gin"
// )

// func SetupRoutes(router *gin.Engine) {
//     // API version 1
//     api := router.Group("/api/v1")

//     // Auth routes (public)
//     auth := api.Group("/auth")
//     {
//         auth.POST("/register", controllers.Register)
//         auth.POST("/verify-email", controllers.VerifyEmail)
//         auth.POST("/login", controllers.Login)
//         auth.POST("/forgot-password", controllers.ForgotPassword)
//         auth.POST("/reset-password", controllers.ResetPassword)
//         auth.POST("/logout", controllers.Logout)
//     }

//     // Public blog routes
//     public := api.Group("/blogs")
//     {
//         public.GET("/", controllers.GetBlogs) // Get all published blogs
//         public.GET("/:id", middleware.AuthMiddleware(), controllers.GetBlog) // Get single blog (auth required to access unpublished)
//     }

//     // Protected routes (require authentication)
//     protected := api.Group("/")
//     protected.Use(middleware.AuthMiddleware())
//     {
//         // Blog management routes
//         blogs := protected.Group("/blogs")
//         {
//             blogs.POST("/", controllers.CreateBlog)           // Create new blog
//             blogs.PUT("/:id", controllers.UpdateBlog)         // Update blog
//             blogs.DELETE("/:id", controllers.DeleteBlog)      // Delete blog
//         }

//         // User's own blogs
//         protected.GET("/my-blogs", controllers.GetBlogs) // Will use user_only=true query param
//     }

//     // Health check route
//     router.GET("/health", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "status":  "OK",
//             "message": "Blog API is running",
//         })
//     })
// }
// ```
