package server

import (
	sonik "forum/internal/controller/http"
	"forum/internal/repository"
	"forum/internal/service"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Serve() {
	db, err := repository.InitDB(repository.Config{
		Username: "sqlite3",
		DBName:   "./database/forum.db",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	if err := repository.CreateTables(db); err != nil {
		log.Println(err)
	}
	defer db.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := sonik.NewHandler(services, errorLog, infoLog)

	srv := &http.Server{
		Addr:     ":4000", ////10-11s
		ErrorLog: errorLog,
		Handler:  handlers.Routes(),
	}

	infoLog.Printf("Запуск веб-сервера на http://127.0.0.1:4000")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// type Server struct {
// 	httpServer *http.Server
// }

// func (s *Server) ServerRun(port string, handler http.Handler) error {
// 	s.httpServer = &http.Server{
// 		Addr:         ":" + port,
// 		Handler:      handler,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 5 * time.Second,
// 	}
// 	return s.httpServer.ListenAndServe()
// }
