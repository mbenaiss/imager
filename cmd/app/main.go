package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/mbenaiss/imager/image"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}

	svc := image.NewProcessor()

	server := gin.Default()
	server.GET("/*o_image", processHandler(svc))

	err = server.Run(":" + cfg.HTTPPort)
	if err != nil {
		panic(fmt.Errorf("unable to start server: %w", err))
	}
}
