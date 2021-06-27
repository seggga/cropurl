package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/seggga/cropurl/internal/storage"
	"github.com/seggga/cropurl/internal/storage/model"
	"go.uber.org/zap"
)

func NewLink(stor storage.CropURLStorage, slogger *zap.SugaredLogger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		// obtain request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Unable to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// compose dataset
		linkPair := new(model.LinkData)
		err = json.Unmarshal(body, linkPair)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		longURL, err := url.Parse(linkPair.LongURL)
		// check user's input: incorrect URL format
		if err != nil {
			fmt.Fprintf(rw, "entered long URL cannot be recognized (%s)", linkPair.LongURL)
			return
		}
		// check user's input: scheme is empty
		if longURL.Scheme == "" {
			fmt.Fprintf(rw, "protocol should be set (http:// or https:// or ...) (%s)", linkPair.LongURL)
			return
		}
		// check user's input: host not set
		if longURL.Host == "" {
			fmt.Fprintf(rw, "host address was not set (%s)", linkPair.LongURL)
			return
		}

		// check if shortID is free
		if stor.IsSet(linkPair.ShortID) {
			fmt.Fprintf(rw, "short URL %s is already in use", linkPair.ShortID)
			return
		}

		// add dataset to the database
		err = stor.AddURL(linkPair)
		if err != nil {
			slogger.Errorw("error creating new short-to-long pair", err)
			fmt.Fprintf(rw, "cannot add %s as a short URL for %s", linkPair.ShortID, linkPair.LongURL)
			return
		}
		// output message
		fmt.Fprintf(rw, "data has been writen\n%s\n%s", linkPair.ShortID, linkPair.LongURL)
	}
}
