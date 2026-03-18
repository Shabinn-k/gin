package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"admin": "1234",
}

func main() {
	r := gin.Default()

	// Global Logging Middleware
	r.Use(LoggerMiddleware())

	// Public routes
	public := r.Group("/")
	{
		public.GET("/", homeHandler)
		public.POST("/login", loginHandler)
		public.GET("/logout", logoutHandler)
	}

	// Protected routes
	protected := r.Group("/dashboard")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/", dashboardHandler)
	}

	r.Run(":8080")
}
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		fmt.Println("➡️", c.Request.Method, c.Request.URL.Path)

		c.Next()

		fmt.Println("✅ Completed in:", time.Since(start))
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := c.Cookie("session")

		if err != nil || user == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Please login first",
			})
			c.Abort()
			return
		}

		// store user in context (optional)
		c.Set("user", user)

		c.Next()
	}
}
func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome! Use POST /login",
	})
}
func loginHandler(c *gin.Context) {
	var login struct {
		Username string `json:"username" binding:"required,min=3"`
		Password string `json:"password" binding:"required,min=4"`
	}

	// Validate input
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username must be >=3 chars & password >=4 chars",
		})
		return
	}

	// Check credentials
	if pass, ok := users[login.Username]; ok && pass == login.Password {

		// Set cookie (session)
		c.SetCookie(
			"session",      // name
			login.Username, // value
			3600,           // expires (1 hr)
			"/",
			"localhost",
			false, // secure (true in HTTPS)
			true,  // httpOnly
		)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid username or password",
	})
}
func logoutHandler(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
func dashboardHandler(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to your dashboard 🚀",
		"user":    user,
	})
}