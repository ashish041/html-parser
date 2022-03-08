package main

import (
	"github.com/ashish041/html-parser/internal/core/domain"
	"github.com/ashish041/html-parser/internal/handlers/httphdl"
	"github.com/gin-gonic/gin"
)

func main() {
	l := domain.DomainLogic{}
	httpHandler := httphdl.NewHTTPHandler(l)

	router := gin.New()
	router.GET("/", httpHandler.Get)

	router.Run(":8080")
}
