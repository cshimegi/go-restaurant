package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

const (
	addr = ":8000"
)

func main() {
	engine := gin.Default()
	initRouters(engine)
	engine.Run(addr)
}
