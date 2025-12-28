package controller

import (
	"net/http"

	"github.com/Lzrb0x/extract-csv-data-api/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ImportDrivers(c *gin.Context) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	file, _ := fileHeader.Open()
	defer file.Close()

	var drivers []model.Driver
	if err := gocsv.Unmarshal(file, &drivers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse CSV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Archive processed successfully",
		"total":   len(drivers),
		"data":    drivers,
	})
}
