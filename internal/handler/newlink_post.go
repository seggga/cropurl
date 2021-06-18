package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/seggga/cropurl/internal/storage"
	"github.com/seggga/cropurl/internal/storage/model"
	"go.uber.org/zap"
)

func NewLink(stor storage.CropURLStorage, slogger *zap.SugaredLogger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Unable to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		linkPair := new(model.LinkData)
		err = json.Unmarshal(body, linkPair)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		if stor.IsSet(linkPair.ShortURL) {
			fmt.Fprintf(rw, "short URL %s is already in use", linkPair.ShortURL)
			return
		}

		err = stor.AddURI(linkPair)
		if err != nil {
			slogger.Errorw("error creating new short-to-long pair", err)
			fmt.Fprintf(rw, "cannot add %s as a short URL for %s", linkPair.ShortURL, linkPair.LongURL)
			return
		}

		fmt.Fprintf(rw, "data has been writen\n%s\n%s", linkPair.ShortURL, linkPair.LongURL)
	}
}
