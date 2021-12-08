package main

import (
	"fmt"
	"log"
	router "mainRoot/infraestructure/router"
	controller "mainRoot/interface/controller"
	http "net/http"
)

func main() {
	controller.LlenarPokedex()
	router.Listen() //Rutas manejadas desde router.go

	fmt.Println("Servidor listo escuchando en " + router.Port)
	log.Fatal(http.ListenAndServe(router.Port, nil))
}
