package internal

import (
	"log"
	"net/http"
)

const (
	welcome = "/"
	// auth
	signup = "/signup"
	signin = "/signin"
	// post
	posts = "/posts/"
	// category
	categories = "/categories/"
	// comment
	createComment       = "/create_comment"
	getCommentsByPostID = "/get_comments_by_post_id"
	getCommentByID      = "/get_comment_by_id"
	// reaction
	setPostReaction    = "/set_post_reaction"
	setCommentReaction = "/set_comment_reaction"

	profile = "/profile/"
)

func Serve() {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	Route(router, fileServer)

	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	log.Println("Server listening... http://127.0.0.1:9999")
	err := http.ListenAndServe("127.0.0.1:9999", router)
	log.Fatal(err)
}

func Route(router *http.ServeMux, fileServer http.Handler) {
	// auth
	router.HandleFunc(signup, signUp)
	router.HandleFunc(signin, signIn)

	router.HandleFunc(welcome, home)
	// router.HandleFunc(posts, app.post)
	// router.HandleFunc(signup, app.signUp)
	// router.HandleFunc(signin, app.signIn)
	// router.HandleFunc(profile, app.profile)
	// router.HandleFunc("/logout", app.logOut)
	// router.HandleFunc("/create-post", app.createPost)
	// router.HandleFunc("/comment/like/", app.likeComment)
	// router.HandleFunc("/post/like/", app.likePost)
	router.Handle("/static", http.NotFoundHandler())
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
}

// type Server struct {
// 	httpServer *http.Server
// }

// func (s *Server) Run(port string) error {
// 	s.httpServer = &http.Server{
// 		Addr:           ":" + port,
// 		MaxHeaderBytes: 1 << 20,
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 	}

// 	return s.httpServer.ListenAndServe()
// }

// func (s *Server) Shutdown(ctx context.Context) error {
// 	return s.httpServer.Shutdown(ctx)
// }
