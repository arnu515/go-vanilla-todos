# Go Vanilla todos

This is a basic CRUD application providing a REST API that allows you to create/read/update/delete todo items.

A todo item has four fields:
```go
id         int     // The unique ID of the todo
text       string  // The actual text of the todo
completed  bool    // Whether the todo has been marked completed
created_at string  // An ISO datetime string denoting the creation time of the todo
```

It stores all the todos in a `map` which is stored in-memory.

## Available API routes

```http
GET /todos  # Get all todos in an array

GET /todos/:id  # Get a single todo

POST /todos  # Create a todo
Content-Type: application/json

{ "text": "Example text" }

PUT /todos/:id  # Update a todo
Content-Type: application/json

{ "text": "Example text", "completed": true }

DELETE /todos/:id  # Delete a todo

DELETE /todos  # Delete all todos
```

## Run the server locally

Clone the repository and run `make` to start the server at port `5000`. This port can be changed by setting the `PORT` environment variable.
