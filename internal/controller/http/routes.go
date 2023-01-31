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
	// urlPostDelete = "/post/delete"
	// urlPostUpdate = "/post/update"
	urlPost = "/post/"
	// urlPosts    = "/post/all"
	urlPostLike = "/post/like"

	urlFilterCategory = "/filter/category/"
	urlFilterLike     = "/filter/like"
	urlFilterDislike  = "/filter/dislike/"

	// user
	urlUser = "/user/"
	// urlUserProfile = "/profile"
	// urlUserUpdate = "/user/update"
	// urlUserDelete = "/user/delete"

	// comment
	urlCommentCreate = "/comment/create"
	// urlCommentDelete = "/comment/delete"
	// urlCommentUpdate = "/comment/update "
	// urlComment     = "/comment/"
	urlComments    = "/comments/all/"
	urlCommentLike = "/comment/like"
)

func (h *Handler) Routes() *http.ServeMux {
	// Security Level
	m := new(middleware)
	m.addMidlleware(h.userIdentity)
	m.addMidlleware(h.logRequest)
	m.addMidlleware(h.secureHeaders)

	// Router
	mux := http.NewServeMux()

	mux.HandleFunc(urlHome, m.chain(h.Home)) // for all

	// auth
	mux.HandleFunc(urlSignUP, m.chain(h.SignUp))   // for all
	mux.HandleFunc(urlSignIn, m.chain(h.SignIn))   // for all
	mux.HandleFunc(urlRefresh, m.chain(h.Refresh)) // for auth
	mux.HandleFunc(urlLogout, m.chain(h.Logout))   // for auth

	// User
	mux.HandleFunc(urlUser, m.chain(h.GetUser)) // for auth
	// mux.HandleFunc(urlUserProfile, m.chain(h.profile)) // for owner
	// mux.HandleFunc(urlUserDelete, h.)       	// for owner
	// mux.HandleFunc(urlUserUpdate, h.)         // for owner

	// post
	mux.HandleFunc(urlPost, m.chain(h.GetPost)) // for all
	// mux.HandleFunc(urlPosts, m.chain(h.ListPosts))       // for all
	mux.HandleFunc(urlPostCreate, m.chain(h.CreatePost)) // for auth
	// mux.HandleFunc(urlPostDelete, m.chain(h.DeletePost)) // for owner
	// mux.HandleFunc(urlPostUpdate, m.chain(h.UpdatePost)) // for owner
	mux.HandleFunc(urlPostLike, m.chain(h.LikePost)) // for owner

	mux.HandleFunc(urlFilterCategory, m.chain(h.ListPosts))      // for all
	mux.HandleFunc(urlFilterDislike, m.chain(h.ListPostsByLike)) // for all
	mux.HandleFunc(urlFilterLike, m.chain(h.ListPostsByLike))    // for all

	// comment
	// mux.HandleFunc(urlComment, m.chain(h.GetPost))             // for all
	mux.HandleFunc(urlComments, m.chain(h.ListComments))       // for all
	mux.HandleFunc(urlCommentCreate, m.chain(h.CreateComment)) // for auth
	// mux.HandleFunc(urlCommentDelete, m.chain(h.DeletePost))    // for owner
	// mux.HandleFunc(urlCommentUpdate, m.chain(h.UpdatePost))    // for owner
	mux.HandleFunc(urlCommentLike, m.chain(h.LikeComment)) // for owner

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
