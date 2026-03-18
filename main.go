// http Method
// package main
// import (
// 	"encoding/json"
// 	"net/http"
// )
// type User struct{
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// }
// var users []User
// func userHandle(w http.ResponseWriter,r *http.Request){
// 	switch r.Method{
// 	case "GET":
// 		json.NewEncoder(w).Encode(users)
// 	case "POST":
// 		var newUser User
// 		if err:=json.NewDecoder(r.Body).Decode(&newUser);err!=nil{
// 			http.Error(w,"invalid data",http.StatusBadRequest)
// 			return
// 		}
// 		users=append(users, newUser)
// 		json.NewEncoder(w).Encode(newUser)
// 	case "PUT":
// 		var updateUser User
// 		if err:=json.NewDecoder(r.Body).Decode(&updateUser);err!=nil{
// 			http.Error(w,"invalid data",http.StatusBadRequest)
// 			return
// 		}
// 		for i,user:=range users{
// 			if user.Id==updateUser.Id{
// 				users[i]=updateUser
// 				json.NewEncoder(w).Encode(updateUser)
// 				return
// 			}
// 		}
// 		http.Error(w,"user not ffound",http.StatusNotFound)
// 	case "DELETE":
// 		var del User
// 		if err:=json.NewDecoder(r.Body).Decode(&del);err!=nil{
// 			http.Error(w,"invalid data",http.StatusBadRequest)
// 			return
// 		}
// 		for i,user:=range users{
// 			if user.Id==del.Id{
// 				users = append(users[:i],users[i+1:]... )
// 				json.NewEncoder(w).Encode(map[string]string{"message":"user deleted"})
// 				return
// 			}
// 		}
// 		http.Error(w, "user not found", http.StatusNotFound)
// 	default:
// 		http.Error(w,"not allowed",http.StatusMethodNotAllowed)
// 	}

// }
// func main(){
// 	http.HandleFunc("/user",userHandle)
// 	http.ListenAndServe(":2007",nil)
// }

// hash password and compare
// package main
// import (
// 	"fmt"
// 	"golang.org/x/crypto/bcrypt"
// )
// func hashPass(password string)(string,error){
// 	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
// 	return string(hash),err
// }
// func compare(hash,pasword string)bool{
// 	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pasword))
// 	return err==nil
// }
// func main() {
// 	hash,_:=hashPass("1234")
// 	fmt.Println(hash)
// 	if compare(hash,"1234"){
// 		fmt.Println(true)
// 	}else{
// 		fmt.Println(false)
// 	}
// }

// basic gin
// package main
// import 	"github.com/gin-gonic/gin"
// func main() {
// 	r:=gin.Default()
// 	r.GET("/",func(c *gin.Context){
// 		c.String(200,"hello world")
// 	})
// 	r.Run(":2000")
// }

// Gin Methods
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
	{Id: 1, Name: "shabin", Age: 19},
}
func getUser(c *gin.Context){
	c.JSON(http.StatusOK,users)
}
func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{"message": newUser})
}
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var update User
	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}
	for i, user := range users {
		if user.Id == id {
			update.Id = id
			users[i] = update
			c.JSON(http.StatusOK, gin.H{"message": update})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
}
func Deleteuser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i,user:=range users{
		if user.Id==id{
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK,gin.H{"messgae":"user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound,gin.H{"error":"no found user"})
}
func main() {
	r := gin.Default()
	api:=r.Group("/api")
	api.GET("/user",getUser)
	api.POST("/user",createUser)
	api.PUT("/user",UpdateUser)
	api.DELETE("/user",Deleteuser)
	r.Run(":2000")	

}
