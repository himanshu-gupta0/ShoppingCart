package routes

import (
	"net/http"

	"backend/controllers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all routes and middleware.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Simple CORS for frontend on different port
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Public endpoints
	r.POST("/users", controllers.RegisterUser)
	r.GET("/users", controllers.ListUsers)
	r.POST("/users/login", controllers.LoginUser)

	r.POST("/items", controllers.CreateItem) // could be protected in real app
	r.GET("/items", controllers.ListItems)

	// Protected endpoints (require token)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/carts", controllers.AddItemToCart)
		auth.GET("/carts", controllers.ListCarts)

		auth.POST("/orders", controllers.CreateOrder)
		auth.GET("/orders", controllers.ListOrders)
	}

	return r
}
