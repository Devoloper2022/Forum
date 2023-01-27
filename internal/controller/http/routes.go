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
	urlLogout  = "/logout"

	// post
	urlPostCreate = "/post/create"
	urlPostDelete = "/post/delete"
	urlPostUpdate = "/post/update"
	urlPost       = "/post/"
	urlPosts      = "/post/all"
	urlPostLike   = "/post/like"
	urlFilter     = "/filter/all"

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
	urlComments      = "/comment/all/"
	urlCommentLike   = "/comment/like"
)

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(urlHome, h.Home) // for all

	// auth
	mux.HandleFunc(urlSignUP, h.SignUp)   // for all
	mux.HandleFunc(urlSignIn, h.SignIn)   // for all
	mux.HandleFunc(urlRefresh, h.Refresh) // for auth
	mux.HandleFunc(urlLogout, h.Logout)   // for auth

	// User
	mux.HandleFunc(urlUser, h.GetUser)        // for auth
	mux.HandleFunc(urlUserProfile, h.profile) // for owner
	// mux.HandleFunc(urlUserDelete, h.)       	// for owner
	// mux.HandleFunc(urlUserUpdate, h.)         // for owner

	// post
	mux.HandleFunc(urlPost, h.GetPost)          // for all
	mux.HandleFunc(urlPosts, h.ListPosts)       // for all
	mux.HandleFunc(urlPostCreate, h.CreatePost) // for auth
	mux.HandleFunc(urlPostDelete, h.DeletePost) // for owner
	mux.HandleFunc(urlPostUpdate, h.UpdatePost) // for owner
	mux.HandleFunc(urlPostLike, h.LikePost)     // for owner

	// comment
	mux.HandleFunc(urlComment, h.GetPost)             // for all
	mux.HandleFunc(urlComments, h.ListPosts)          // for all
	mux.HandleFunc(urlCommentCreate, h.CreateComment) // for auth
	mux.HandleFunc(urlCommentDelete, h.DeletePost)    // for owner
	mux.HandleFunc(urlCommentUpdate, h.UpdatePost)    // for owner
	mux.HandleFunc(urlCommentLike, h.LikeComment)     // for owner

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
