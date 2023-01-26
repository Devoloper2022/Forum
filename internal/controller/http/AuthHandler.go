package http

import (
	"fmt"
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

	posts, err := h.services.GetAllPosts()
	fmt.Printf("Level handler result: %s", posts)
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, posts)
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
