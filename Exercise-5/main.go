package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve the HTML form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<body>
    <h1>Create a New Task</h1>
    <form action="/tasks" method="POST">
        <label>Task Name:</label><br>
        <input type="text" name="task_name"><br><br>
        
        <label>Priority:</label><br>
        <input type="text" name="priority"><br><br>
        
        <button type="submit">Create Task</button>
    </form>
</body>
</html>
`
		fmt.Fprintf(w, html)
	})

	// Handle task creation
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse the form data from the request
			r.ParseForm()

			// Get the values the user submitted
			taskName := r.FormValue("task_name")
			priority := r.FormValue("priority")

			// Send back a response showing what we received
			fmt.Fprintf(w, "Task Created!\n\n")
			fmt.Fprintf(w, "Task Name: %s\n", taskName)
			fmt.Fprintf(w, "Priority: %s\n", priority)
		} else {
			fmt.Fprintf(w, "Please use POST method to create tasks")
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
