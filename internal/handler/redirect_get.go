package handler

import (
	"fmt"
	"net/http"

	"github.com/seggga/cropurl/internal/storage"
)

func Redirect(stor storage.CropURLStorage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		URI := r.URL.Path
		longURL, err := stor.Resolve(URI)
		if err != nil {
			return
		}

		fmt.Fprint(rw, "long URL is ", longURL)
	}
}
