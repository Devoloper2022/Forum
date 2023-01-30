package http

import (
	dto "forum/internal/DTO"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) errorHandler(w http.ResponseWriter, code int, errorText string) {
	w.WriteHeader(code)
	sysErr := dto.SystemErr{
		Status: code,
		Name:   http.StatusText(code),
		Msg:    errorText,
	}

	ts, err := template.ParseFiles("./ui/templates/error.html")
	if err != nil {
		log.Printf("Create Post: Execute:%v", err)
		return
	}

	err = ts.Execute(w, sysErr)
	if err != nil {
		h.serverError(w, err)
		return
	}
}
