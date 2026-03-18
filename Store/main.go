package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	r := gin.Default()

	// Create cookie store
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// LOGIN
	r.POST("/login", func(c *gin.Context) {
		var data map[string]string

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}

		username := data["username"]
		password := data["password"]

		// simple check (demo)
		if username != "admin" || password != "1234" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Create session
		session := sessions.Default(c)
		session.Set("user", username)
		session.Save()

		// Also set cookie
		c.SetCookie("user", username, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
		})
	})

	// DASHBOARD (protected)
	r.GET("/dashboard", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "welcome " + user.(string),
		})
	})

	// LOGOUT
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"message": "logged out",
		})
	})

	r.Run(":2007")
}