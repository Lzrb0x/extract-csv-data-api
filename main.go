package main

import (
	"log"

	"github.com/Lzrb0x/extract-csv-data-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
