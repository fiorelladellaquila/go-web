package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	puerto = ":8080"
)

// Persona es una estructura que describre el modelo de una persona.
type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {

	// Punto 1. Creamos una persona e imprimimos en consola en formato JSON
	persona := Persona{
		Nombre:    "Juan",
		Apellido:  "Perez",
		Edad:      20,
		Direccion: "Av. Siempre Viva",
		Telefono:  "01112343456345",
		Activo:    true,
	}

	json, err := json.Marshal(persona)
	if err != nil {
		log.Fatal(err)
	}

	data := string(json)
	fmt.Println(data)

	// Punto 2. Creamos otra persona y lo enviamos al cliente en formato JSON.
	// Definimos el router/engine de GIN.
	router := gin.Default()

	// Creamos un enpoint de control.
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	})

	// Creamos endpoint de consulta de una persona.
	router.GET("/persona", handlePersona)

	// Ejecutamos el servidor en el puerto :8080
	router.Run(puerto)

}

// Controller/handler que devuelve una persona.
func handlePersona(ctx *gin.Context) {
	personEndpoint := Persona{
		Nombre:    "Marcos",
		Apellido:  "Silva",
		Edad:      25,
		Direccion: "Av Mariano Moreno",
		Telefono:  "01123423452",
		Activo:    true,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": personEndpoint,
	})
}
