package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/new", controllers.CreateBlog)
}
