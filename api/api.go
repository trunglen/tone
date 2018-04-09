package api

import (
	"tone/api/tone"

	"github.com/gin-gonic/gin"
)

func NewApiServer(root *gin.RouterGroup) {
	tone.NewToneServer(root, "tone")
}
