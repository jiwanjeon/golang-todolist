# ToDo list with Golang

# Tools

- Gorm
- Json marshall, unmarshall
- Gorilla Mux
- postgreSQL

# Routes

- Get Todos <-- /todo/ <-- GET
- Create Todo <-- /todo/ <-- POST
- Get Todo By ID <-- /todo/{todoid} <-- GET
- Delete Todo <-- /todo/{todoid} <-- DELETE
- Update Todo <-- /todo/{todoid} <-- UPDATE
- Complete Todo <-- /todo/{todoid} <-- PUT
- Incomplete Todo <-- /todo/{todoid} <-- PUT

# Project Struct

cmd --> main --> main.go
pkg
-->config --> app.go
-->controllers --> todo-controller.go
-->models --> todo.go
-->routes --> todolist-routes.go
-->utils --> utils.go
