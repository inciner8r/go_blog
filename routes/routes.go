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
	router.GET("/blogs", controllers.GetAllBlogs)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/user", controllers.User)
}
