package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "dst")

		c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
	})

	router.Run(":8080")
}
