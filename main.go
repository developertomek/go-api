package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/developertomek/go-api/internal/app"
	"github.com/developertomek/go-api/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go abckend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	app.Logger.Println("we are running app")

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
