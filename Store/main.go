package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store:=cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession",store))
	r.POST("/login",func(c *gin.Context){
		var data map[string]string
		if err:=c.BindJSON(&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"invalid data"})
			return
		}
		username:=data["username"]
		password:=data["password"]
		if username!="shabin"||password!="1234"{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
			return
		}
		session:=sessions.Default(c)
		session.Set("user",username)
		session.Save()

		c.SetCookie("user",username,3600,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{"message":"logged in"})
	})
	r.GET("/dashboard",func(c *gin.Context){
		session:=sessions.Default(c)
		user:=session.Get("user")
		if user==nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":"no data found"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message":"welcome "+user.(string)})
	})
	r.GET("/logout",func(c *gin.Context){
		session:=sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK,gin.H{"message":"logged out"})
	})
	r.Run()
}