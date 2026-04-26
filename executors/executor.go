package executors

import (
	"fmt"
	"net/http"
	"urlShortner/models"
	"urlShortner/storage"
)

type Executor struct {
	store *storage.Store
}

func NewHandler(store *storage.Store) *Executor {
	return &Executor{store: store}
}

func (h *Executor) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	r.ParseForm()
	fmt.Println(r.Form)
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "LongURL is empty", http.StatusBadRequest)
		return
	}
	shortUrl := models.GenerateUrl()
	h.store.Save(shortUrl, longURL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://shorty.url/" + shortUrl))

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
