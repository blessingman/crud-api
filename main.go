package main

import (
	"net/http"

	"github.com/blessingman/crud-api/config"
	"github.com/blessingman/crud-api/controller"
	"github.com/blessingman/crud-api/helper"
	"github.com/blessingman/crud-api/model"
	"github.com/blessingman/crud-api/repository"
	"github.com/blessingman/crud-api/router"
	"github.com/blessingman/crud-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

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
