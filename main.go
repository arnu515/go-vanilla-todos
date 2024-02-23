package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"todolist/todos/handlers"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("\\/todos\\/(\\d+){0,1}")
	isOnSingleTodo, todoId := false, uint(0)
	if pathMatch := re.FindStringSubmatch(r.URL.Path); len(pathMatch) == 2 && len(pathMatch[1]) > 0 {
		isOnSingleTodo = true
		todoIdFromPath, err := strconv.ParseUint(pathMatch[1], 0, 32)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		todoId = uint(todoIdFromPath)
	}
	fmt.Println(isOnSingleTodo, todoId)
	if r.Method == "GET" {
		if isOnSingleTodo {
			handlers.GetOne(todoId, w)
		} else {
			handlers.GetAll(w)
		}
	} else if r.Method == "POST" {
		if r.Header.Get("content-type") != "application/json" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if isOnSingleTodo {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		handlers.Create(w, r)
	} else if r.Method == "PUT" {
		if r.Header.Get("content-type") != "application/json" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if !isOnSingleTodo {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		handlers.Update(todoId, w, r)
	} else if r.Method == "DELETE" {
		if isOnSingleTodo {
			handlers.DeleteOne(todoId, w)
		} else {
			handlers.DeleteAll(w)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/todos/", todosHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	portFromEnv := os.Getenv("PORT")
	port := 5000
	if parsedPort, err := strconv.Atoi(portFromEnv); err == nil {
		port = parsedPort
	}
	fmt.Println("Starting the server on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
