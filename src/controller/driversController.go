package controller

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImportCSV(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Lê o cabeçalho (primeira linha)
	headers, err := reader.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read CSV headers"})
		return
	}

	var data []map[string]string

	// Lê todas as linhas restantes
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse CSV"})
			return
		}

		// Cria um mapa para cada linha usando os headers como chaves
		row := make(map[string]string)
		for i, value := range record {
			if i < len(headers) {
				row[headers[i]] = value
			}
		}
		data = append(data, row)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Archive processed successfully",
		"total":   len(data),
		"headers": headers,
		"data":    data,
	})
}
