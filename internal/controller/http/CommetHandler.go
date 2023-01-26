package http

import "net/http"

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) ListComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}

func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdatePost из page"))
}
