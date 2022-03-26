package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/new", controllers.CreateBlog)
	router.GET("/blog/:blogId", controllers.GetABlog)
	router.PUT("/blog/:blogId", controllers.EditABlog)
	router.DELETE("/blog/:blogId", controllers.DeleteABlog)
}
