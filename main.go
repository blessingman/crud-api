package main

import (
	"net/http"

	"github.com/blessingman/crud-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	routes := gin.Default()

	log.Info().Msg("Started Server!")
	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
