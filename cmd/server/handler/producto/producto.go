package producto

import (
	"net/http"
	"os"
	"strconv"

	"github.com/aldogayaladh/go-web/internal/domain/producto"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service producto.Service
}

func NewControladorProducto(service producto.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := os.Getenv("TOKEN")
		tokenHeader := ctx.GetHeader("tokenPostman")

		if tokenHeader != token {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "token invalido",
			})
			return
		}

		productos, err := c.service.GetAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": productos,
		})
	}
}

func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"mensaje": "id invalido",
			})
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "no se pudo eliminar el producto",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "producto eliminado",
		})
	}
}
