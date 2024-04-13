package services

import (
	"OZONTestCaseLinks/database"
	"encoding/json"
	"net/http"
	"strings"
)

func LinkReducer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/")
	var link string
	var err error

	if path != "" {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		link, err = database.OriginalLink(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var url string
		err = json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error decoding json"))
			return
		}

		link, err = database.ReduceLink(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		link = r.Host + "/" + link
	}

	jsonResp, err := json.Marshal(link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while marhal result"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
