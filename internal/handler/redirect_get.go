package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/seggga/cropurl/internal/storage"
)

/*
308 redirect to long URL
400 Invalid short URL supplied
*/

func Redirect(stor storage.CropURLStorage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// define shortID from users query
		shortID := chi.URLParam(r, "shortID")
		// defint corresponding long URL from database
		longURL, err := stor.Resolve(shortID)
		if err != nil {
			fmt.Fprintf(rw, "there is no URL linked to %s", shortID)
			return
		}
		// implement redirect
		http.Redirect(rw, r, longURL, http.StatusPermanentRedirect)
	}
}
