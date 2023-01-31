package http

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
	}

	if r.URL.Path != urlCommentCreate {
		h.notFound(w)
		return
	}

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			h.errorLog.Printf("error parse form:%v", err)
			return
		}

		text := r.PostFormValue("text")
		postId := r.PostFormValue("postId")
		id, err := strconv.Atoi(postId)

		if err != nil || id < 1 {
			h.notFound(w)
			return
		}

		if text == "" {
			h.clientError(w, 400)
			return
		}

		err = h.services.Comment.CreateComment(dto.CommentDto{
			Text:   text,
			User:   dto.UserDto{ID: user.ID},
			PostID: int64(id),
		})
		if err != nil {
			h.serverError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

// func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("UpdatePost из page"))
// }

// func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("UpdatePost из page"))
// }

// func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("UpdatePost из page"))
// }

func (h *Handler) ListComments(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlComments {
		h.notFound(w)
		return
	}

	if r.Method != "GET" {
		h.notFound(w)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	files := []string{
		"./ui/templates/comment/comments.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.serverError(w, err)
		return
	}

	comments, err := h.services.Comment.GetAllCommentsByPostId(int64(id))
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, comments)
	if err != nil {
		h.serverError(w, err)
	}
}

func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
	}
	w.Write([]byte("UpdatePost из page"))
}
