package app

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var tpl = template.Must(template.ParseGlob("app/templates/*.tmpl"))
var StablishedConnection = dbConnection()

type Employee struct {
	Id    int
	Name  string
	Email string
}

func index(w http.ResponseWriter, r *http.Request) {
	registry, err := StablishedConnection.Query("SELECT * FROM employees")

	if err != nil {
		fmt.Println("Something went wrong")
		panic(err.Error())
	}

	employee := Employee{}
	arrEmployee := []Employee{}

	for registry.Next() {
		var id int
		var name, email string

		err = registry.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email

		arrEmployee = append(arrEmployee, employee)
	}

	tpl.ExecuteTemplate(w, "index", arrEmployee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create", nil)
}

func insertEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		name = strings.Title(strings.ToLower(name))

		insertRegistry, err := StablishedConnection.Prepare("INSERT INTO employees(name, email) VALUES(?, ?)")

		if err != nil {
			fmt.Println("Something went wrong")
			panic(err.Error())
		}

		insertRegistry.Exec(name, email)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := r.URL.Query().Get("id")

	deleteRegistry, err := StablishedConnection.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		fmt.Println("Something went wrong")
		panic(err.Error())
	}

	deleteRegistry.Exec(employeeId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := r.URL.Query().Get("id")

	registry, err := StablishedConnection.Query("SELECT * FROM employees WHERE id=?", employeeId)

	employee := Employee{}

	for registry.Next() {
		var id int
		var name, email string

		err = registry.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}

		employee.Id = id
		employee.Name = name
		employee.Email = email
	}

	tpl.ExecuteTemplate(w, "edit", employee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		name = strings.Title(strings.ToLower(name))

		modifyRegistry, err := StablishedConnection.Prepare("UPDATE employees SET name=?, email=? WHERE id=?")

		if err != nil {
			fmt.Println("Something went wrong")
			panic(err.Error())
		}

		modifyRegistry.Exec(name, email, id)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
