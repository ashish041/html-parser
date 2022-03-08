package httphdl

import (
	"github.com/ashish041/html-parser/internal/core/domain"
	"github.com/ashish041/html-parser/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Domain domain.DomainLogic
}

func NewHTTPHandler(l ports.DomainService) *HTTPHandler {
	v := l.(domain.DomainLogic)
	return &HTTPHandler{Domain: v}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	response, err := hdl.Domain.New(c.Query("url"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	defer response.ResponseBody.Close()
	info, err := hdl.Domain.Get(response)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, info)
}
