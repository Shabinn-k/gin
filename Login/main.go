package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
var users =map[string]string{
	"shabin":"1234",
}
func Auth()gin.HandlerFunc{
	return func(c *gin.Context){
		user,err:=c.Cookie("session")
		if err != nil||user==""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"no users found"})
			c.Abort()
			return
		}
		c.Set("user",user)
		c.Next()
	}
}

func login(c *gin.Context){
	var Login struct{
		Username string `json:"username" binding:"required min=3"`
		Password string `json:"password" binding:"required min=4"`
	}
	if err:=c.ShouldBindJSON(&Login);err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"name and pass no meet required format"})
		return
	}
	if pass,ok:=users[Login.Username];ok||pass==Login.Password{
		c.SetCookie("session",Login.Username,3600,"/","localhost",false,true)
		c.JSON(http.StatusOK,gin.H{"message":"logged in"})
		return
	}
	c.JSON(http.StatusUnauthorized,gin.H{"error":"invalid cred"})
}
func logout(c *gin.Context){
	c.SetCookie("session","",-1,"/","localhost",false,true)
	c.JSON(http.StatusOK,gin.H{"meddage":"logged out"})
}
func dashboard(c *gin.Context){
	user,_:=c.Get("user")
	c.JSON(http.StatusOK,gin.H{
		"message":"welcome",
		"name":user,
	})
}
func home(c *gin.Context){
	c.String(http.StatusOK,"Login to dashboard")
}

func main() {
	r:=gin.Default()
	api:=r.Group("/")
	api.GET("/",home)
	api.POST("/login",login)
	api.GET("/logout",logout)
	protect:=r.Group("/dashboard")
	protect.Use(Auth())
	protect.GET("/",dashboard)
	r.Run(":2000")
}