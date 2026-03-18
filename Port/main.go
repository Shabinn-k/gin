package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(users)
	case "POST":
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}
		users = append(users, newUser)
		json.NewEncoder(w ).Encode(newUser)
		fmt.Println("User created")
	case "PUT":
		var updateUser User
		if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
			http.Error(w, "invalid", http.StatusBadRequest)
			return
		}
		for i, user := range users {
			if user.Id == updateUser.Id {
				users[i] = updateUser
				json.NewEncoder(w).Encode(updateUser)
				return
			}
		}
		http.Error(w, "user not found", http.StatusNotFound)
	case "DELETE":
		var req User
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}
		for i, user := range users {
			if user.Id == req.Id {
				users = append(users[:i], users[i+1:]...)

				json.NewEncoder(w).Encode(map[string]string{
					"message": "user deleted",
				})
				return
			}
		}

		http.Error(w, "user not found", http.StatusNotFound)
	default:
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/user", handle)
	fmt.Println("server runing")
	http.ListenAndServe(":2006", nil)
}
