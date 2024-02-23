package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	todoStore "todolist/todos/store"
)

var store *todoStore.TodoStore = todoStore.New()

func GetAll(w http.ResponseWriter) {
	t := store.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

func GetOne(id uint, w http.ResponseWriter) {
	t := store.GetOne(id)
	if t == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

func Create(w http.ResponseWriter, r *http.Request) {
	type PostBody struct {
		Text string `json:"text"`
	}
	var postBody PostBody
	err := json.NewDecoder(r.Body).Decode(&postBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := store.Create(postBody.Text)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

func Update(id uint, w http.ResponseWriter, r *http.Request) {
	type UpdateBody struct {
		Text      string `json:"text"`
		Completed bool   `json:"completed"`
	}
	var updateBody UpdateBody
	err := json.NewDecoder(r.Body).Decode(&updateBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := store.GetOne(id)
	if t == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	t.Text = updateBody.Text
	t.Completed = updateBody.Completed
	store.Update(id, *t)

	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

func DeleteOne(id uint, w http.ResponseWriter) {
	t := store.Delete(id)
	if t == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

func DeleteAll(w http.ResponseWriter) {
	todos := store.GetAll()

	for _, t := range todos {
		store.Delete(t.Id)
	}

	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}
