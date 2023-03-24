package httphandler

import (
	"github.com/gin-gonic/gin"
	"app/internal/handlers/user"
)

func NewHTTPServer(userHdl userHandler.UserHdl) *gin.Engine {
	g := gin.Default()
	g.GET("/health", func(ctx *gin.Context) { ctx.String(200, "server is runing...") })

	v1 := g.Group("v1")
	v1.POST("/user", userHdl.Register)
	v1.POST("/user/login", userHdl.Login)
	v1.Use(userHdl.AuthMiddleware)
	v1.PUT("/user", userHdl.UpdateUser)
	v1.DELETE("/user", userHdl.DeleteUser)
	v1.GET("/user", userHdl.GetUser)
	v1.GET("/user/logout", userHdl.LogOut)
	v1.GET("/user/check", func(ctx *gin.Context) { ctx.String(200, "pass") })
	
	return g
}