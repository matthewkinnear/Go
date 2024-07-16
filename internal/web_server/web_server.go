package webserver

import (
	"html/template"
	types "my-pubsub-app/utils"
	"net/http"
)

var (
	users    []types.User
	userChan <-chan types.User
)

// SetUserChan sets the user channel for receiving updates
func SetUserChan(ch <-chan types.User) {
	userChan = ch
}

func StartWebServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Random Users</title>
				<link rel="icon" href="/static/favicon.ico" type="image/x-icon">
			</head>
			<body>
				<h1>Random Users</h1>
				<table border="1">
					<tr>
						<th>Name</th>
						<th>Gender</th>
						<th>Email</th>
					</tr>
					{{range .}}
					<tr>
						<td>{{.Name}}</td>
						<td>{{.Gender}}</td>
						<td>{{.Email}}</td>
					</tr>
					{{end}}
				</table>
			</body>
			</html>
		`))

		// Render the template with the current list of users
		tmpl.Execute(w, users)
	})

	// Update the users list whenever new data is published
	go func() {
		for user := range userChan {
			users = append(users, user)
		}
	}()

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
