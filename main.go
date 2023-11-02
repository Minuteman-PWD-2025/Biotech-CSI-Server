package main

import (
	"fmt"
	"html/template"
	_ "github.com/lib/pq"
)

type User struct {
	ID int
	Name  string
	Email string
	Password string
	Privilege int
}
// for debugging the sending and recieving
var tmpl = template.Must(template.New("index").Parse(`
<html>
<head>
    <title>User List</title>
</head>
<body>
    <table border="1">
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Email</th>
			<th>Password</th>
			<th>Privilege</th>

        </tr>
        {{range .}}
            <tr>
                <td>{{.ID}}</td>
                <td>{{.Name}}</td>
                <td>{{.Email}}</td>
				<td>.Password</td>
				<td>.Privilege</td>
            </tr>
        {{end}}
    </table>
</body>
</html>
`))

func main() {
	fmt.Printf("Starting Server...\n")
}
