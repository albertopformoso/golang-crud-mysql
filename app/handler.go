package app

import (
	"fmt"
	"net/http"
	"text/template"
	_ "github.com/go-sql-driver/mysql"
)

var tpl = template.Must(template.ParseGlob("app/templates/*.tmpl"))



func index (w http.ResponseWriter, r *http.Request) {
	stablishedConnection := dbConnection()
	insertRegistry, err := stablishedConnection.Prepare("INSERT INTO employees(name, email) VALUES('test', 'test@work.com')")

	if err != nil {
		fmt.Println("Something went wrong")
		panic(err.Error())
	}

	insertRegistry.Exec()

	tpl.ExecuteTemplate(w, "index", nil)
}

func createEmployee (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create", nil)
}