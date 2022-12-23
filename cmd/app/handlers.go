package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/mbenaiss/imager/image"
)

func processHandler(svc *image.Processor) gin.HandlerFunc {
	return func(c *gin.Context) {
		op := image.Operation{
			OperationType: c.Param("o"),
			Width:         strToInt(c.Query("w")),
			Height:        strToInt(c.Query("h")),
			Quality:       strToInt(c.Query("q")),
			Format:        c.Query("f"),
		}

		url := strings.TrimLeft(c.Request.URL.Path, "/")

		newImage, contentType, err := svc.ProcessFromURL(c.Request.Context(), url, op)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		c.Header("Cache-Control", "public, max-age=86400")
		c.Header("Content-Length", strconv.Itoa(len(newImage)))
		c.Data(http.StatusOK, contentType, newImage)
	}
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}
