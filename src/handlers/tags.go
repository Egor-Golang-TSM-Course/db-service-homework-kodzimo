package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Tag struct {
	ID           int       `json:"id"`
	Tag          string    `json:"tag"`
	CreationDate time.Time `json:"creation_date"`
	PostsIDs     []int     `json:"posts_ids"`
}

// TagHandler хранит теги и обеспечивает потокобезопасный доступ
type TagHandler struct {
	sync.Mutex
	tags   map[int]Tag
	nextID int
}

// Добавление тега к посту. Тег нужно создать и добавить в мапу с тегами если его не существует.
func (h *TagHandler) addTagToPost(tag Tag, postId int, w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.tags)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// POST запрос
}

// Получение списка тегов
func (h *TagHandler) displayTags(w http.ResponseWriter, r *http.Request) {
	h.Lock()
	defer h.Unlock()

	err := json.NewEncoder(w).Encode(h.tags)
	if err != nil {
		http.Error(w, "json.NewEncoder(w).Encode error", http.StatusInternalServerError)
		return
	}

	// Здеcь должна быть описана логика фунции
	// GET запрос
}

// Запуск хендлера. Выполняется в мейне
func tagHandler() {
	handler := &TagHandler{
		tags:   make(map[int]Tag, 0),
		nextID: 1,
	}

	http.HandleFunc("/posts/{postId}/tags", handler.addTagToPost(tag, postId))
	http.HandleFunc("/tags", handler.displayTags)
}
