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

/*
return
	200 successful operation
	400 Invalid short URL supplied
*/

func Delete(stor storage.CropURLStorage, slogger *zap.SugaredLogger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// obtain request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			slogger.Errorw("Unable to parse request body", err)
			http.Error(rw, "Unable to parse request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// decode entered data in a structure
		linkPair := new(model.LinkData)
		err = json.Unmarshal(body, linkPair)
		if err != nil {
			slogger.Errorw("Unable to unmarshal JSON", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = stor.Delete(linkPair.ShortID)
		if err != nil {
			slogger.Errorw("error deleting short-to-long pair", err)
			fmt.Fprintf(rw, "cannot delete %s", linkPair.ShortID)
			return
		}

		fmt.Fprintf(rw, "data has been deleted\n%s\n%s", linkPair.ShortID, linkPair.LongURL)
	}
}
