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

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}
	if r.URL.Path != urlPostCreate {
		h.notFound(w)
		return
	}

	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/post/createPost.html")
		if err != nil {
			log.Printf("Create Post: Execute:%v", err)
			return
		}

		categories, err := h.services.GetAllCategories()
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
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
			log.Println("error parse form :", err)
			return
		}
		title := r.PostFormValue("title")
		title1 := strings.ReplaceAll(title, " ", "")

		text := r.PostFormValue("text")
		text1 := strings.ReplaceAll(text, " ", "")

		categories := r.Form["categories"]

		if title1 == "" || text1 == "" || categories == nil {
			h.clientError(w, 400)
			return
		}

		pid, err := h.services.CreatePost(dto.PostDto{Title: title, Text: text, User: dto.UserDto{ID: user.ID}}, categories)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlPost {
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

	ts, err := template.ParseFiles("./ui/templates/post/post.html")
	if err != nil {
		log.Printf("Get Post: Execute:%v", err)
		return
	}

	post, err := h.services.GetPost(int64(id))
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, post)
	if err != nil {
		h.serverError(w, err)
		return
	}
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlFilterCategory {
		h.notFound(w)
		return
	}

	files := []string{
		"./ui/templates/index.html",
	}
	// as
	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.serverError(w, err)
		return
	}

	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}
	//

	posts, err := h.services.GetAllPostsByCategory(int64(id))
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})

	if err != nil {
		h.serverError(w, err)
	}
}

func (h *Handler) ListPostsByLike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlFilterDislike || r.URL.Path != urlFilterLike {
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

	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	var posts []dto.PostDto
	if r.URL.Path == urlFilterDislike {
		posts, err = h.services.Post.GetAllPostsByLike("like")
	} else {
		posts, err = h.services.Post.GetAllPostsByLike("dislike")
	}

	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})

	if err != nil {
		h.serverError(w, err)
	}
}

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlPostLike {
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

		postId := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postId)

		if err != nil || pid < 1 {
			h.notFound(w)
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
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		http.Redirect(w, r, fmt.Sprintf(urlSignIn), http.StatusSeeOther)
		return
	}

	if r.URL.Path != urlPostDislike {
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

		postId := r.PostFormValue("postId")

		pid, err := strconv.Atoi(postId)

		if err != nil || pid < 1 {
			h.notFound(w)
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
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/post?id=%d", pid), http.StatusSeeOther)
	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}
