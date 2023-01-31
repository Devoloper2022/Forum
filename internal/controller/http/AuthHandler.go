package http

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != urlHome {
		h.notFound(w)
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
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	posts, err := h.services.GetAllPosts()
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = ts.Execute(w, dto.Index{
		List: categories,
		Post: posts,
	})
	if err != nil {
		h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user != (models.User{}) {
		http.Redirect(w, r, urlHome, http.StatusSeeOther)
	}
	if r.URL.Path != urlSignUP {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/signUp.html")
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return
		}

		email := r.PostFormValue("email")
		email = strings.ReplaceAll(email, " ", "")

		username := r.PostFormValue("username")
		username = strings.ReplaceAll(username, " ", "")

		pass := r.PostFormValue("password")
		pass = strings.ReplaceAll(pass, " ", "")

		repass := r.PostFormValue("repassw")
		repass = strings.ReplaceAll(repass, " ", "")

		if email == "" || username == "" || pass == "" || repass == "" || repass != pass {
			h.errorHandler(w, http.StatusBadRequest, "Not valid input ")
			return
		}

		err = h.services.CreateUser(models.User{
			Username: username,
			Email:    email,
			Password: pass,
		})
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		http.Redirect(w, r, urlSignIn, http.StatusSeeOther)

	} else {
		h.errorLog.Println(http.StatusText(http.StatusMethodNotAllowed))
		h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(key).(models.User)
	if user != (models.User{}) {
		http.Redirect(w, r, urlHome, http.StatusSeeOther)
	}
	if r.URL.Path != urlSignIn {
		h.errorHandler(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/templates/signIn.html")
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			h.errorHandler(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return
		}

		email := r.PostFormValue("email")
		email = strings.ReplaceAll(email, " ", "")

		pass := r.PostFormValue("password")
		pass = strings.ReplaceAll(pass, " ", "")

		if email == "" || pass == "" {
			h.errorHandler(w, http.StatusBadRequest, "Not valid input ")
			return
		}

		cook, err := h.services.Autorization.GenerateToken(dto.Credentials{
			Username: email,
			Password: pass,
		})
		if err != nil {
			h.errorHandler(w, http.StatusBadRequest, "Not valid input ")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   cook.Token,
			Expires: cook.Expiry,
			Path:    urlHome,
		})

		http.Redirect(w, r, urlHome, http.StatusSeeOther)

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
	user := r.Context().Value(key).(models.User)
	if user == (models.User{}) {
		h.errorHandler(w, http.StatusBadRequest, "can't log-out,without log-in")
		return
	}

	if r.URL.Path != urlLogout {
		h.notFound(w)
		return
	}

	if r.Method != "POST" {
		token, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				h.errorHandler(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}
		}
		err = h.services.DeleteToken(token.Value)
		if err != nil {
			h.errorHandler(w, http.StatusInternalServerError, err.Error())
			return
		}
		c := &http.Cookie{
			Name:    "token",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, urlHome, http.StatusSeeOther)
	}
	w.Write([]byte("Logout page"))
}
