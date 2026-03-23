// // package main
// // import (
// //
// //	"encoding/json"
// //	"net/http"
// //
// // )
// //
// //	type User struct {
// //		Id   int    `json:"id"`
// //		Name string `json:"name"`
// //	}
// //
// // var users []User
// //
// //	func userhandle(w http.ResponseWriter,r *http.Request){
// //		switch r.Method{
// //		case "GET":
// //			json.NewEncoder(w).Encode(users)
// //		case "POST":
// //			var newUser User
// //			json.NewDecoder(r.Body).Decode(&newUser)
// //			users=append(users,newUser)
// //		case "PUT":
// //			var updateuser User
// //			json.NewDecoder(r.Body).Decode(&updateuser)
// //			users[0]=updateuser
// //		case "DELETE":
// //			users=users[1:]
// //		}
// //	}
// //
// //	func main(){
// //		http.HandleFunc("/user",userhandle)
// //		http.ListenAndServe(":2000",nil)
// //	}

// // package main
// // import (
// // 	"fmt"
// // 	"golang.org/x/crypto/bcrypt"
// // )
// // func hashpass(password string)(string,error){
// // 	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
// // 	return string(hash),err
// // }
// // func compare(hash,paassword string)bool{
// // 	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(paassword))
// // 	return err==nil
// // }
// // func main(){
// // 	hash,_:=hashpass("1234")
// // 	fmt.Println(hash)
// // 	if compare(hash,"1234"){
// // 		fmt.Println(true)
// // 	}else{
// // 		fmt.Println(false)
// // 	}
// // }

// // get with gin

// package main
// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )
// type User struct{
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }
// var users []User
// func main() {
// 	r:=gin.Default()
// 	r.POST("/users",func(c *gin.Context){
// 		var newUser []User
// 		if err:=c.BindJSON(&newUser);err!=nil{
// 			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid"})
// 			return
// 		}
// 		users = append(users, newUser...)
// 		c.JSON(http.StatusOK,gin.H{
// 			"message":newUser,
// 		})
// 	})	
// 	r.GET("/",func(c *gin.Context){
// 		c.String(200,"hello world")		
// 	})
// 	r.Run()
// }
 

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// First, register a new user
	fmt.Println("1. Registering user...")
	registerData := map[string]string{
		"name":     "Test User",
		"email":    "test123@example.com",
		"password": "password123",
	}
	jsonData, _ := json.Marshal(registerData)
	
	resp, err := http.Post("http://localhost:8080/api/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Register error: %v\n", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Register Status: %d, Response: %s\n\n", resp.StatusCode, string(body))
	resp.Body.Close()
	
	// Then login with the same credentials
	fmt.Println("2. Logging in...")
	loginData := map[string]string{
		"email":    "test123@example.com",
		"password": "password123",
	}
	jsonData, _ = json.Marshal(loginData)
	
	resp, err = http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Login error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Login Status: %d\n", resp.StatusCode)
	fmt.Printf("Login Response: %s\n", string(body))
	resp.Body.Close()
}