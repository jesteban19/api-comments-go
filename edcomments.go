package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jesteban19/edcomments/commons"
	"github.com/jesteban19/edcomments/migration"
	"github.com/jesteban19/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la BD")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto para el servidor web")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizo la migración.")
	}

	//inicia las rutas
	router := routes.InitRoutes()

	//inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	//Incia el servidor

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado el servidor en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecucion programa")
}
