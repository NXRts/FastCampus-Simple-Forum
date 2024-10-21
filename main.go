package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
