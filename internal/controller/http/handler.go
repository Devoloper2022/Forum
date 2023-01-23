package http

import (
	"fmt"
	"forum/internal/service"
	"log"
	"net/http"
	"runtime/debug"
)

type Handler struct {
	services *service.Service
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewHandler(services *service.Service, errorLog *log.Logger, infoLog *log.Logger) *Handler {
	return &Handler{
		services: services,
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func (h *Handler) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handler) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	h.clientError(w, http.StatusNotFound)
}
