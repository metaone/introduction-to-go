package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
	"html/template"
)

type Item struct {
	Id int
	Value string
}

var (
	db *sql.DB
	createQuery *sql.Stmt
	readQuery *sql.Stmt
	updateQuery *sql.Stmt
	deleteQuery *sql.Stmt
)

const tpl = `
<h1>Items</h1>

<table>
	{{ range $key, $value := .}}
		<tr>
			<td>
				<form>
					<label for="value">ID: </lebel>
					<input type="text" value="{{ $value.Id }}" disabled>
				</form>
			</td>
			<td>
				<form method="POST" action="/items">
					<input type="hidden" name="action" value="PUT">
					<input type="hidden" name="id" value="{{ $value.Id }}">
					<label for="value">Value: </lebel>
					<input type="text" name="value" placeholder="Item value" value="{{ $value.Value }}">
					<input type="submit" value="Update">
				</form>
			</td>
			<td>
				<form method="POST" action="/items">
					<input type="hidden" name="action" value="DELETE">
					<input type="hidden" name="id" value="{{ $value.Id }}">
					<input type="submit" value="DELETE">
				</form>
			</td>
		</tr>
	{{ end }}
</table>

<hr />

<form method="POST" action="/items">
	<input type="hidden" name="action" value="POST">
	<input type="text" name="value" placeholder="Item value">
	<input type="submit" value="Add">
</form>
`


func main() {
	db, _ = sql.Open("mysql", "root:root@/crud")
	defer db.Close()

	createQuery, _ = db.Prepare("insert into items (value) values (?);")
	defer createQuery.Close()

	readQuery, _ = db.Prepare("select * from items;")
	defer readQuery.Close()

	updateQuery, _ = db.Prepare("update items set value=? where id=?;")
	defer updateQuery.Close()

	deleteQuery, _ = db.Prepare("delete from items where id=?;")
	defer deleteQuery.Close()

	http.HandleFunc("/items", handler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := template.New("index")
		t, _ = t.Parse(tpl)
		t.Execute(w, read())
	case "POST":
		switch r.FormValue("action") {
		case "POST":
			create(r.FormValue("value"))
		case "PUT":
			update(r.FormValue("id"), r.FormValue("value"))
		case "DELETE":
			delete(r.FormValue("id"))
		}
		http.Redirect(w, r, "/items", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/items", http.StatusSeeOther)
	}
}

func create(value string) {
	if value != "" {
		createQuery.Exec(value)
	}
}

func read() []Item {
	rows, err := readQuery.Query()
	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var r Item
		_ = rows.Scan(&r.Id, &r.Value)
		items = append(items, r)
	}

	return items
}

func update(id string, value string) {
	if id != "" && value != "" {
		updateQuery.Exec(value, id)
	}
}

func delete(id string) {
	if id != "" {
		deleteQuery.Exec(id)
	}
}
