// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

//	func main() {
//		password:="shabin"
//		hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
//		if err!=nil{
//			panic(err)
//		}
//		fmt.Println(string(hash))
//	}

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
func HashPass(password string)(string,error){
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(hash),err
}
func compare(hash,password string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err==nil
}
func main() {
	hash,_:=HashPass("shabi")
	fmt.Println(hash)
	if compare(hash,"shabin"){
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}