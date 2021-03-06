package main

import (
	"fmt"
	"net/http"

	"github.com/tchebraga/rest-api-go/internal/comment"
	"github.com/tchebraga/rest-api-go/internal/database"
	transportHTTP "github.com/tchebraga/rest-api-go/internal/transport/http"
)

// App - contains pointers to db and others
type App struct{}

// Run - sets up aplication
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
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
