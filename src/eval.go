package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func eval(ctx *gin.Context) {
	message, err := execute(processRequest(ctx))
	if err != nil {
		message = "Failed to execute"
	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
