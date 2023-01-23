package http

import (
	"fmt"
	"forum/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/post/createPost.html")
		if err != nil {
			log.Printf("Create Post: Execute:%v", err)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else if r.Method == "POST" {

		title := "История про улитку"
		content := "Улитка выползла из ."
		var userID int64 = 60

		id, err := h.services.CreatePost(models.Post{Title: title, Text: content, UserID: userID})
		if err != nil {
			h.serverError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/", id), http.StatusSeeOther)
		w.Write([]byte("CreatePost page"))
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePost из page"))
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	fmt.Fprintf(w, "GetPost ID %d...", id)
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ListPosts из page"))
}
