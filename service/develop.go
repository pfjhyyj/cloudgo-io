package service

import (
	"net/http"
)

func developHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Under developing", http.StatusNotImplemented)
	}
}
