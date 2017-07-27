package main

import (

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"posts/core"

)




func main() {

	router := gin.Default()

    //Routes
	router.GET(core.POST,  func(c *gin.Context){core.GetPost(c)})
	router.GET(core.POSTS, func (c *gin.Context){core.GetPosts(c)})
	router.POST(core.POST, func(c *gin.Context) { core.SetPost(c)})
	router.DELETE(core.POST, func(c *gin.Context) { core.DelPost(c)})




	router.Run(":8080")



}
