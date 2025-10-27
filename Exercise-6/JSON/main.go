package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task represents a task in our system
// The `json:"..."` tags tell Go how to convert between Go structs and JSON
type Task struct {
	Name     string `json:"name"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}

func main() {
	// Serve a simple HTML page with JavaScript to send JSON
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<body>
    <h1>Create Task (JSON)</h1>
    <input type="text" id="taskName" placeholder="Task name">
    <input type="text" id="priority" placeholder="Priority">
    <button onclick="createTask()">Create Task</button>
    
    <h2>Response:</h2>
    <pre id="response"></pre>

    <script>
    function createTask() {
        const taskName = document.getElementById('taskName').value;
        const priority = document.getElementById('priority').value;
        
        // Send JSON data to the server
        fetch('/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: taskName,
                priority: priority
            })
        })
        .then(response => response.json())
        .then(data => {
            document.getElementById('response').textContent = JSON.stringify(data, null, 2);
        });
    }
    </script>
</body>
</html>
`
		fmt.Fprintf(w, html)
	})

	// Handle JSON task creation
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		// Let's see what Content-Type the browser is sending
		fmt.Println("Content-Type:", r.Header.Get("Content-Type"))

		// Decode the JSON from the request body
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Received task:", task)

		// Process the task (in real app, save to database)
		task.Status = "created"

		// Send JSON response back
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
