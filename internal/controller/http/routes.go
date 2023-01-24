package http

import (
	"net/http"
)

const (
	// general
	urlHome = "/"

	// auth
	urlSignIn  = "/signin"
	urlSignUP  = "/signup"
	urlRefresh = "/refresh"
	urlLogout  = ""

	// post
	urlPostCreate = "/post/create"
	urlPostDelete = "/post/delete"
	urlPostUpdate = "/post/update"
	urlPost       = "/post/"
	urlPosts      = "/post/all"

	// user
	urlUser        = "/user/"
	urlUserProfile = "/profile"
	urlUserUpdate  = "/user/update"
	urlUserDelete  = "/user/delete"

	// comment
	urlCommentCreate = "/comment/create"
	urlCommentDelete = "/comment/delete"
	urlCommentUpdate = "/comment/update "
	urlComment       = "/comment/"
	urlComments      = "/comment/all"
)

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(urlHome, h.Home)
	// post
	mux.HandleFunc(urlPost, h.GetPost)
	mux.HandleFunc(urlPosts, h.ListPosts)
	mux.HandleFunc(urlPostCreate, h.CreatePost)
	mux.HandleFunc(urlPostDelete, h.DeletePost)
	mux.HandleFunc(urlPostUpdate, h.UpdatePost)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
