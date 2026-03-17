package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

var users []User

func handle(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "GET":
		json.NewEncoder(w).Encode(users)
	case "POST":
		var newUser User
		json.NewDecoder(r.Body).Decode(&newUser)
		users = append(users, newUser)
		fmt.Println("User created")
	case "PUT":
		var updateUser User
		json.NewDecoder(r.Body).Decode(&updateUser)
		users[0]=updateUser
		fmt.Println("updated user")
	case "DELETE":
		users=users[1:]
		fmt.Println("usesr deleted")
	default:
		http.Error(w,"Not allowed",http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/user",handle)
	fmt.Println("server runing")
	http.ListenAndServe(":2006",nil)
}