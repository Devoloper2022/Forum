package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	h.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	w.Write([]byte("CreatePost page"))
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletePost из page"))
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	fmt.Fprintf(w, "GetPost ID %d...", id)
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ListPosts из page"))
}
