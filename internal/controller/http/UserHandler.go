package http

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlUser {
		h.notFound(w)
		return
	}
	if r.Method != "GET" {
		h.clientError(w, 400)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	ts, err := template.ParseFiles("./ui/templates/user/user.html")
	if err != nil {
		log.Printf("Get Post: Execute:%v", err)
		return
	}

	user, err := h.services.User.Get(int64(id))
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, user)
	if err != nil {
		h.serverError(w, err)
		return
	}
}

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}
