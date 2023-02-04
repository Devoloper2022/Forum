package http

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlCommentCreate {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		text := r.PostFormValue("text")
		text1 := strings.ReplaceAll(text, " ", "")

		if text1 == "" {
			h.errorHandler(w, http.StatusBadRequest, "Invalid input")
			return
		}

		postId := r.PostFormValue("postId")
		id, err := strconv.Atoi(postId)

		if err != nil || id < 1 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		err = h.services.Comment.CreateComment(dto.CommentDto{
			Text:   text,
			User:   dto.UserDto{ID: user.ID},
			PostID: int64(id),
		})
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) ListComments(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlComments {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != "GET" {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
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
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	comments, err := h.services.Comment.GetAllCommentsByPostId(int64(id))
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = ts.Execute(w, comments)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlCommentLike {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		likeID := r.PostFormValue("id")

		id, err := strconv.Atoi(likeID)

		if err != nil || id < 0 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		commentId := r.PostFormValue("commentId")

		cid, err := strconv.Atoi(commentId)

		if err != nil || cid < 1 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		postID := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postID)

		if err != nil || id < 0 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		err = h.services.LikeComment(models.CommentLike{
			ID:        int64(id),
			CommentID: int64(cid),
			UserID:    user.ID,
			Like:      true,
			DisLike:   false,
		})

		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/comments/all?id=%d", pid), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) DislikeComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlCommentDislike {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		likeID := r.PostFormValue("id")

		id, err := strconv.Atoi(likeID)

		if err != nil || id < 0 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		commentId := r.PostFormValue("commentId")

		cid, err := strconv.Atoi(commentId)

		if err != nil || cid < 1 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		postID := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postID)

		if err != nil || id < 0 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		err = h.services.DislikeComment(models.CommentLike{
			ID:        int64(id),
			CommentID: int64(cid),
			UserID:    user.ID,
			Like:      false,
			DisLike:   true,
		})

		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/comments/all?id=%d", pid), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}
