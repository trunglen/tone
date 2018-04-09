package main

import (
	"net/http"
	"tone/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.NewApiServer(r.Group("api"))
	r.StaticFS("/static", http.Dir("./upload"))
	http.ListenAndServe(":3006", r)
	// r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}
