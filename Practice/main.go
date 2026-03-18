// package main
// import (
//
//	"encoding/json"
//	"net/http"
//
// )
//
//	type User struct {
//		Id   int    `json:"id"`
//		Name string `json:"name"`
//	}
//
// var users []User
//
//	func userhandle(w http.ResponseWriter,r *http.Request){
//		switch r.Method{
//		case "GET":
//			json.NewEncoder(w).Encode(users)
//		case "POST":
//			var newUser User
//			json.NewDecoder(r.Body).Decode(&newUser)
//			users=append(users,newUser)
//		case "PUT":
//			var updateuser User
//			json.NewDecoder(r.Body).Decode(&updateuser)
//			users[0]=updateuser
//		case "DELETE":
//			users=users[1:]
//		}
//	}
//
//	func main(){
//		http.HandleFunc("/user",userhandle)
//		http.ListenAndServe(":2000",nil)
//	}

// package main
// import (
// 	"fmt"
// 	"golang.org/x/crypto/bcrypt"
// )
// func hashpass(password string)(string,error){
// 	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
// 	return string(hash),err
// }
// func compare(hash,paassword string)bool{
// 	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(paassword))
// 	return err==nil
// }
// func main(){
// 	hash,_:=hashpass("1234")
// 	fmt.Println(hash)
// 	if compare(hash,"1234"){
// 		fmt.Println(true)
// 	}else{
// 		fmt.Println(false)
// 	}
// }

// get with gin

package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}
var users []User
func main() {
	r:=gin.Default()
	r.POST("/users",func(c *gin.Context){
		var newUser []User
		if err:=c.BindJSON(&newUser);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid"})
			return
		}
		users = append(users, newUser...)
		c.JSON(http.StatusOK,gin.H{
			"message":newUser,
		})
	})	
	r.GET("/",func(c *gin.Context){
		c.String(200,"hello world")		
	})
	r.Run()
}