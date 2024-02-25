package application

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/AhmedSamy16/01-Book-Directory-Go/internal/database"
	_ "github.com/lib/pq"
)

type App struct {
	router http.Handler
	DB     *database.Queries
}

func New() *App {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB URL is not found")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		DB: database.New(conn),
	}

	app.loadRoutes()

	return app
}

func (app *App) Start() {
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found")
	}

	log.Println("Starting the server on port:", portString)
	server := &http.Server{
		Addr:    ":" + portString,
		Handler: app.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
