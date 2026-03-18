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

//		r.Run(":8080")
//	}

package main
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "shabin", Age: 18},
	{Id: 2, Name: "john", Age: 20},
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{
		"message": newUser,
	})
}
func updateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var Update User
	if err := c.BindJSON(&Update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	for i, user := range users {
		if user.Id == id {
			Update.Id = id
			users[i]=Update
			c.JSON(http.StatusOK, gin.H{"message": Update})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
}
func deleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/user", getUser)
	api.POST("/user", createUser)
	api.PUT("/user/:id", updateUser)
	api.DELETE("/user/:id", deleteUser)
	r.Run(":2007")
}