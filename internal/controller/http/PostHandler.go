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

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}
	if r.URL.Path != urlPostCreate {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/post/createPost.html")
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		categories, err := h.services.GetAllCategories()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = ts.Execute(w, categories)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		title := r.PostFormValue("title")
		title1 := strings.ReplaceAll(title, " ", "")

		text := r.PostFormValue("text")
		text1 := strings.ReplaceAll(text, " ", "")

		categories := r.Form["categories"]

		if title1 == "" || text1 == "" || categories == nil {
			h.errorHandler(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		pid, err := h.services.CreatePost(dto.PostDto{Title: title, Text: text, User: dto.UserDto{ID: user.ID}}, categories)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlPost {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != "GET" {
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	ts, err := template.ParseFiles("./ui/templates/post/post.html")
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	post, err := h.services.GetPost(int64(id))
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = ts.Execute(w, post)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlFilterCategory {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	files := []string{
		"./ui/templates/index.html",
	}
	// as
	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	//

	posts, err := h.services.GetAllPostsByCategory(int64(id))
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})

	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (h *Handler) ListPostsByLike(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path == urlFilterDislike || r.URL.Path == urlFilterLike {
	} else {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != "GET" {
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	files := []string{
		"./ui/templates/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	var posts []dto.PostDto
	if r.URL.Path == urlFilterDislike {
		posts, err = h.services.Post.GetAllPostsByLike("dislike")
	} else {
		posts, err = h.services.Post.GetAllPostsByLike("like")
	}

	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})

	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlPostLike {
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

		postId := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postId)

		if err != nil || pid < 1 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		err = h.services.LikePost.LikePost(models.PostLike{
			ID:      int64(id),
			PostID:  int64(pid),
			UserID:  user.ID,
			Like:    true,
			DisLike: false,
		})

		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlPostDislike {
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

		postId := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postId)

		if err != nil || pid < 1 {
			h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
			return
		}

		err = h.services.LikePost.DislikePost(models.PostLike{
			ID:      int64(id),
			PostID:  int64(pid),
			UserID:  user.ID,
			DisLike: true,
			Like:    false,
		})

		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) ListPostsByDate(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlFilterDate {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method != "GET" {
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	files := []string{
		"./ui/templates/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	posts, err := h.services.Post.GetAllPostsByDate()
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})

	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}
