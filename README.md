# ToDo list with Golang

## ✅ Tools

- GORM
- Gorilla Mux
- PostgreSQL
- Unit Test (testify)
- Docker
- Github Action CI/CD
- graphQL

## ✅ What wii you get after this project?

- By using the above tools, you will be able to get directions for how to use them.

## ✅ Routes

- Get Todos <-- /todo/ <-- GET
- Create Todo <-- /todo/ <-- POST
- Get Todo By ID <-- /todo/{todoid} <-- GET
- Delete Todo <-- /todo/{todoid} <-- DELETE
- Update Todo <-- /todo/{todoid} <-- UPDATE
- Complete Todo <-- /todo/{todoid} <-- PUT
- Incomplete Todo <-- /todo/{todoid} <-- PUT

## ✅ Project Structure

cmd --> main --> main.go <br>
pkg <br>
-->config --> app.go<br>
-->controllers --> todo-controller.go<br>
-->models --> todo.go<br>
-->routes --> todolist-routes.go<br>
-->utils --> utils.go<br>
