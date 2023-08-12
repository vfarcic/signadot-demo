package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetOutput(os.Stderr)
	// Server
	log.Println("Starting server...")
	router := gin.New()
	router.GET("/", rootHandler)
	router.GET("/ping", pingHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}

func rootHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "This is a silly demo xxx")
}

func httpErrorBadRequest(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusBadRequest)
}

func httpError(err error, ctx *gin.Context, status int) {
	log.Println(err.Error())
	ctx.String(status, err.Error())
}
