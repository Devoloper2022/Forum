package http

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlUser {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != "GET" {
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.errorHandler(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	ts, err := template.ParseFiles("./ui/templates/user/user.html")
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	user, err := h.services.User.Get(int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = ts.Execute(w, user)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}
