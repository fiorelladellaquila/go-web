package routes

import (
	"database/sql"

	"github.com/aldogayaladh/go-web/pkg/middleware"

	"github.com/aldogayaladh/go-web/cmd/server/handler/ping"
	handlerProducto "github.com/aldogayaladh/go-web/cmd/server/handler/producto"

	"github.com/aldogayaladh/go-web/internal/domain/producto"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildProductRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildProductRoutes maps all routes for the product domain.
func (r *router) buildProductRoutes() {
	// Create a new product controller.
	repository := producto.NewRepositoryMySql(r.db)
	service := producto.NewService(repository)
	controlador := handlerProducto.NewControladorProducto(service)

	r.routerGroup.POST("/productos", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/productos", middleware.Authenticate(), controlador.GetAll())
	r.routerGroup.GET("/productos/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.DELETE("/productos/:id", middleware.Authenticate(), controlador.Delete())

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControladorPing()
	r.routerGroup.GET("/ping", pingController.Ping())

}
