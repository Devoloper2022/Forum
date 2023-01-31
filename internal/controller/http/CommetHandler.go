package http

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"html/template"
	"log"
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
		h.notFound(w)
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

		postId := r.PostFormValue("postId")
		id, err := strconv.Atoi(postId)

		if err != nil || id < 1 {
			h.notFound(w)
			return
		}

		if text1 == "" {
			h.clientError(w, 400)
			return
		}

		err = h.services.Comment.CreateComment(dto.CommentDto{
			Text:   text,
			User:   dto.UserDto{ID: user.ID},
			PostID: int64(id),
		})
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", id), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

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
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	comments, err := h.services.Comment.GetAllCommentsByPostId(int64(id))
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
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
		h.notFound(w)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("error parse form :", err)
			return
		}
		likeID := r.PostFormValue("id")

		id, err := strconv.Atoi(likeID)

		if err != nil || id < 0 {
			h.notFound(w)
			return
		}

		commentId := r.PostFormValue("commentId")

		cid, err := strconv.Atoi(commentId)

		if err != nil || cid < 1 {
			h.notFound(w)
			return
		}

		postID := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postID)

		if err != nil || id < 0 {
			h.notFound(w)
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
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/comments/all?id=%d", pid), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) DislikeComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlCommentDislike {
		h.notFound(w)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("error parse form :", err)
			return
		}
		likeID := r.PostFormValue("id")

		id, err := strconv.Atoi(likeID)

		if err != nil || id < 0 {
			h.notFound(w)
			return
		}

		commentId := r.PostFormValue("commentId")

		cid, err := strconv.Atoi(commentId)

		if err != nil || cid < 1 {
			h.notFound(w)
			return
		}

		postID := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postID)

		if err != nil || id < 0 {
			h.notFound(w)
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
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}
