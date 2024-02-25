package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AhmedSamy16/01-Book-Directory-Go/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type BookHandler struct {
	DB *database.Queries
}

func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid Input")
	}

	book, err := b.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:     uuid.New(),
		Title:  params.Title,
		Author: params.Author,
	})
	if err != nil {
		respondWithError(w, 500, "Something wen wrong, Please try again")
		return
	}
	respondWithJson(w, 201, book)
}

func (b *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := b.DB.GetAllBooks(r.Context())
	if err != nil {
		respondWithError(w, 500, "Something went wrong, Please try again")
		return
	}
	respondWithJson(w, 200, books)
}

func (b *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	bookIdParam := chi.URLParam(r, "id")
	bookId, err := uuid.Parse(bookIdParam)
	if err != nil {
		respondWithError(w, 400, "Invalid book id")
		return
	}
	book, err := b.DB.GetBookById(r.Context(), bookId)
	if err != nil {
		respondWithError(w, 404, "Book Not found")
		return
	}
	respondWithJson(w, 200, book)
}

func (b *BookHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	bookIdParam := chi.URLParam(r, "id")
	bookId, err := uuid.Parse(bookIdParam)
	if err != nil {
		respondWithError(w, 400, "Invalid Id")
		return
	}
	_, err = b.DB.DeleteBookById(r.Context(), bookId)
	if err != nil {
		respondWithError(w, 404, "Book Not found")
		return
	}
	respondWithJson(w, 204, struct{}{})
}

func (b *BookHandler) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	bookIdParam := chi.URLParam(r, "id")
	bookId, err := uuid.Parse(bookIdParam)
	if err != nil {
		respondWithError(w, 400, "Invalid id")
		return
	}

	type Parameter struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameter{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid Book body")
		return
	}

	book, err := b.DB.UpdateBookById(r.Context(), database.UpdateBookByIdParams{
		Title:  params.Title,
		Author: params.Author,
		ID:     bookId,
	})
	if err != nil {
		respondWithError(w, 404, "Book not found")
		return
	}
	respondWithJson(w, 200, book)
}
