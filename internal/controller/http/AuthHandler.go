package http

import (
	"html/template"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notFound(w)
		return
	}

	files := []string{
		"./ui/templates/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		h.serverError(w, err)
	}

	w.Write([]byte("home page"))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignIn page"))
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SignUp page"))
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Refresh page"))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Refresh page"))
}
