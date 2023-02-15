package productcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/ypp2301/app-say-hello/models"
	"net/http"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}
