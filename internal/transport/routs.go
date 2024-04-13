package transport

import (
	"OZONTestCaseLinks/internal/services"
	"net/http"
)

func Routs(mux *http.ServeMux) {
	mux.HandleFunc("/", services.LinkReducer)
}
