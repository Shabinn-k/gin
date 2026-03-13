package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// login page
	r.GET("/login", showLogin)

	// login action
	r.POST("/login", login)

	// dashboard (protected)
	r.GET("/dashboard", dashboard)

	// logout
	r.GET("/logout", logout)

	r.Run(":8080")
}

func showLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send POST request with username and password",
	})
}

func login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	// hardcoded credentials
	if username == "admin" && password == "1234" {

		// create cookie
		c.SetCookie("session", "user_logged_in", 3600, "/", "localhost", false, true)

		c.JSON(200, gin.H{
			"message": "Login successful",
		})

	} else {

		c.JSON(401, gin.H{
			"message": "Invalid credentials",
		})

	}
}

func dashboard(c *gin.Context) {

	session, err := c.Cookie("session")

	if err != nil || session != "user_logged_in" {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Please login first",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Welcome to dashboard",
	})
}

func logout(c *gin.Context) {

	// delete cookie
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message": "Logged out successfully",
	})
}