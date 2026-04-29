package executors

import (
	"fmt"
	"net/http"
	"urlShortner/db"
	"urlShortner/models"
	"urlShortner/storage"

	"github.com/jmoiron/sqlx"
)

type Executor struct {
	store *storage.Store
	db    *sqlx.DB
}

func NewHandler(store *storage.Store) *Executor {
	db := db.SetDB()
	return &Executor{store: store, db: db}
}

func (h *Executor) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "LongURL is empty", http.StatusBadRequest)
		return
	}

	var url models.URLs

	err = h.db.Get(&url,
		`SELECT shorturl, longurl FROM urls WHERE longurl = $1`,
		longURL,
	)

	// ✅ CASE 1: URL already exists
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("http://localhost:8080/" + url.ShortURL))
		return
	}

	// ❌ Real DB error
	if err != nil && err.Error() != "sql: no rows in result set" {
		http.Error(w, "DB error", http.StatusInternalServerError)
		fmt.Println("DB ERROR:", err)
		return
	}

	// ✅ CASE 2: URL not found → create new
	shortURL := models.GenerateUrl()

	_, err = h.db.Exec(
		`INSERT INTO urls (shorturl, longurl) VALUES ($1, $2)`,
		shortURL, longURL,
	)
	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		fmt.Println("INSERT ERROR:", err)
		return
	}

	h.store.Save(shortURL, longURL)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + shortURL))
}

func (h *Executor) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	shortURL := r.URL.Path[1:]
	longURl, exists := h.store.Get(shortURL)
	if !exists {
		http.Error(w, "Url doesn't exist", http.StatusNotFound)
	}
	http.Redirect(w, r, longURl, http.StatusFound)
}
