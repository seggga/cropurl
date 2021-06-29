package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/seggga/cropurl/internal/storage"
)

/*
200 successful operation
400 Invalid short URL
*/

func ViewStatistics(stor storage.CropURLStorage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// define shortID from users query
		shortID := chi.URLParam(r, "shortID")
		// define data about the specified short ID
		urlData, err := stor.ViewStat(shortID)
		if err != nil {
			fmt.Fprintf(rw, "there is no URL linked to %s", shortID)
			return
		}
		// output data
		fmt.Fprintf(
			rw,
			"short ID: %s\nlong URL: %s\ndescriptoin: %s\ncounter: %d\n",
			urlData.ShortID,
			urlData.LongURL,
			urlData.Description,
			urlData.Statistics,
		)
	}
}
