package main 

import "github.com/gin-gonic/gin"

func main (){
	server := gin.Default()

	server.GET("/home", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{"message":"HELLO"})
	})
	server.Run(":8000")
}