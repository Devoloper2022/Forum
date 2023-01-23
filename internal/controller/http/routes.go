package http

import (
	"net/http"
)

const (
	// general
	home = "/"

	// auth
	signIn  = "/signin"
	signUP  = "/signup"
	refresh = "/refresh"
	// Logout  = ""

	// post
	createPost = "/post/create"
	deletePost = "/post/delete"
	updatePost = "/post/update"
	post       = "/post/"
	posts      = "/post/all"

	// user
	user = "/user/"

	// comment
	createComment = "/comment/create"
	deleteComment = "/comment/delete"
	comment       = "/comment"
)

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(home, h.Home)
	// post
	mux.HandleFunc(post, h.CreatePost)
	mux.HandleFunc(posts, h.ListPosts)
	mux.HandleFunc(createPost, h.CreatePost)
	mux.HandleFunc(deletePost, h.DeletePost)
	mux.HandleFunc(updatePost, h.UpdatePost)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
