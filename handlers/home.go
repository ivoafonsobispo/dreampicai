package handler

import (
	"imaginai/views/home"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
