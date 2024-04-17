package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	a "alan/restaurant/appetizers"
	h "alan/restaurant/health"
	m "alan/restaurant/middlewares"
	u "alan/restaurant/users"
)

var (
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(m.JsonLoggerMiddleware())

	u.InitRouters(engine)
	a.InitRouters(engine)
	h.InitRouters(engine)

	if err := engine.Run(port); err != nil {
		panic(err)
	}

	fmt.Println("Listening on port: " + port)
}
