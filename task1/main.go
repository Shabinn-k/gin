// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Server is running",
// 		})
// 	})

// 	r.Run(":8080")
// }


// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	r := gin.Default()

// 	api := r.Group("/api")

// 	api.GET("/test", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "API working",
// 		})
// 	})

// 	r.Run(":8080")
// }

package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {

	r := gin.Default()

	api := r.Group("/api")

	api.GET("/users", getUsers)
	api.POST("/users", createUser)

	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(200, users)
}

func createUser(c *gin.Context) {

	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)

	c.JSON(201, gin.H{
		"message": "User created",
		"user":    newUser,
	})
}