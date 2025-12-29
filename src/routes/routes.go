package routes

import (
	"github.com/Lzrb0x/extract-csv-data-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("import/csv", controller.ImportCSV)
}
