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

	r.Run()
}
