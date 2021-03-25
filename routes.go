package main

import (
	"ACAT/controller"
	_ "ACAT/docs"
	"ACAT/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(router *gin.Engine) *gin.Engine{
   router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Use(middleware.CORSmw())
   usergroup := router.Group("/api/user")
   {
   	usergroup.POST("/sendcode",middleware.Sendmail(),controller.Sendcode)
   	usergroup.POST("/verifycode",controller.Verifycode)
   	usergroup.POST("/register",controller.Register)
   	usergroup.POST("/login",controller.UserLogin)
   	usergroup.POST("/sighup",middleware.Verifymw(),controller.Sighup)
   	usergroup.POST("/changegroup",middleware.Verifymw(),controller.Changegroup)
   }
   return router
}