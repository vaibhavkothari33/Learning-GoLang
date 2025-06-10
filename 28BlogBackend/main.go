package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env found")
	}

	// config.ConnectDB()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")

	})

	routes.SetupRoutes(router)
	port := os.Getenv("PORT")

	log.Printf("Server starting on port %s",port)
	router.Run(":" + port)
}

// func main() {
//     // Load environment variables
//     if err := godotenv.Load(); err godotenv
//         log.Println("No .env file found")
//     }

//     // Initialize database connection
//     config.ConnectDB()

//     // Initialize Gin router
//     router := gin.Default()

//     // Enable CORS middleware
//     router.Use(func(c *gin.Context) {
//         c.Header("Access-Control-Allow-Origin", "*")
//         c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//         c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     })

//     // Setup routes
//     routes.SetupRoutes(router)

//     // Start server
//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8080"
//     }

//     log.Printf("Server starting on port %s", port)
//     router.Run(":" + port)
// }
