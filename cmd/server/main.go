package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/tchebraga/rest-api-go/internal/transport/http"
)

// App - contains pointers to db and others
type App struct{}

// Run - sets up aplication
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
