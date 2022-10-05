package v1

import "github.com/gin-gonic/gin"

func New(handler *gin.Engine, l logger.Interface) {
	//Регистрируем middle ware
	//handler.Use(gin.Logger())
	//handler.Use(gin.Recovery())

	//h := handler.Group("/v1")

}
