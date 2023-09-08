package main

import (
	"log"
	"net/http"

	handlerProducto "github.com/aldogayaladh/go-web/cmd/server/handler/producto"
	"github.com/aldogayaladh/go-web/internal/domain/producto"
	"github.com/aldogayaladh/go-web/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	puerto = ":8080"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())

	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controlador := handlerProducto.NewControladorProducto(service)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	})

	router.GET("/productos", middleware.Authenticate(), controlador.GetAll())
	router.DELETE("/productos/:id", controlador.Delete())

	if err := router.Run(puerto); err != nil {
		panic(err)
	}

}
