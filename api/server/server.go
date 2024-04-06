package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	h "alan/blog/health"
	m "alan/blog/middlewares"
	p "alan/blog/posts"
	u "alan/blog/users"
)

var (
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(m.JsonLoggerMiddleware())

	u.InitRouters(engine)
	p.InitRouters(engine)
	h.InitRouters(engine)

	if err := engine.Run(port); err != nil {
		panic(err)
	}

	fmt.Println("Listening on port: " + port)
}
