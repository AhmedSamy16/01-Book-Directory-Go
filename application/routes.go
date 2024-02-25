package application

import (
	"github.com/AhmedSamy16/01-Book-Directory-Go/handlers"
	"github.com/go-chi/chi/v5"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Route("/books", a.loadBookRoutes)

	a.router = router
}

func (a *App) loadBookRoutes(router chi.Router) {
	bookHandler := &handlers.BookHandler{
		DB: a.DB,
	}

	router.Get("/", bookHandler.GetAllBooks)
	router.Post("/", bookHandler.CreateBook)
	router.Get("/{id}", bookHandler.GetBookById)
	router.Put("/{id}", bookHandler.UpdateBookById)
	router.Delete("/{id}", bookHandler.DeleteBookById)
}
