package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ypp2301/app-say-hello/controllers/productcontroller"
	"github.com/ypp2301/app-say-hello/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product/", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product/", productcontroller.Delete)
	r.Run()
}
