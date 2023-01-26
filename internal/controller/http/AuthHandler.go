package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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

	posts, err := h.services.GetAllPosts()
	if err != nil {
		h.errorLog.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, posts)
	if err != nil {
		h.serverError(w, err)
	}

	w.Write([]byte("home page"))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlSignIn {
		h.notFound(w)
		return
	}
	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/signIn.html")
		if err != nil {
			log.Printf("Create Post: Execute:%v", err)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			h.serverError(w, err)
			return
		}
	} else if r.Method == "POST" {
		email := r.PostFormValue("email")
		pass := r.PostFormValue("password")
		fmt.Println(email)
		fmt.Println(pass)

		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)

	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlSignUP {
		h.notFound(w)
		return
	}

	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/signUp.html")
		if err != nil {
			log.Printf("Create Post: Execute:%v", err)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			h.serverError(w, err)
			return
		}
	} else if r.Method == "POST" {
		email := r.PostFormValue("email")
		username := r.PostFormValue("username")
		pass := r.PostFormValue("password")
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(pass)

		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)

	} else {
		log.Println("Create Post: Method not allowed")
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlRefresh {
		h.notFound(w)
		return
	}
	w.Write([]byte("Refresh page"))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlLogout {
		h.notFound(w)
		return
	}
	w.Write([]byte("Logout page"))
}
