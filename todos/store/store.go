package store

import (
	"errors"
	"sync"

	"todolist/todos"
)

type TodoStore struct {
	todos  map[uint]todos.Todo
	mu     sync.Mutex
	nextId uint
}

func New() *TodoStore {
	return &TodoStore{
		todos:  make(map[uint]todos.Todo),
		nextId: 0,
	}
}

func (s *TodoStore) Create(text string) *todos.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextId
	s.nextId++
	t := todos.NewTodo(id, text)
	s.todos[id] = *t
	return t
}

func (s *TodoStore) GetOne(id uint) *todos.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, exists := s.todos[id]
	if !exists {
		return nil
	}
	return &t
}

func (s *TodoStore) GetAll() []todos.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	sl := make([]todos.Todo, 0, len(s.todos))
	for _, v := range s.todos {
		sl = append(sl, v)
	}
	return sl
}

func (s *TodoStore) Update(id uint, newTodo todos.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.todos[id]; !exists {
		return errors.New("Todo does not exist")
	}
	s.todos[id] = newTodo
	return nil
}

func (s *TodoStore) Delete(id uint) *todos.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t, exists := s.todos[id]; exists {
		delete(s.todos, id)
		return &t
	}
	return nil
}
