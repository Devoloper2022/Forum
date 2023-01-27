package http

import (
	"fmt"
	dto "forum/internal/DTO"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlCommentCreate {
		h.notFound(w)
		return
	}

	// if r.Method == "GET" {
	// 	ts, err := template.ParseFiles("./ui/templates/comment/createComment.html")
	// 	if err != nil {
	// 		h.errorLog.Printf("Create Post: Execute:%v", err)
	// 		return
	// 	}

	// 	err = ts.Execute(w, nil)
	// 	if err != nil {
	// 		h.serverError(w, err)
	// 		return
	// 	}
	// } else
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
		var userID int64 = 1 /// DEL

		err = h.services.Comment.CreateComment(dto.CommentDto{
			Text:   text,
			User:   dto.UserDto{ID: userID},
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
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}
