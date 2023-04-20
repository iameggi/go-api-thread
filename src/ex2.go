package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Address string `json:"address"`
	Age    int    `json:"age"`
}

var employees []Employee

func main() {
	http.HandleFunc("/employees", employeesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getEmployees(w, r)
	case "POST":
		addEmployee(w, r)
	case "PUT":
		updateEmployee(w, r)
	case "DELETE":
		deleteEmployee(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var emp Employee
	err = json.Unmarshal(reqBody, &emp)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}
	emp.ID = uuid.New().String()
	employees = append(employees, emp)
	json.NewEncoder(w).Encode(emp)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var emp Employee
	err = json.Unmarshal(reqBody, &emp)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.ID == emp.ID {
			employees[i] = emp
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var emp Employee
	err = json.Unmarshal(reqBody, &emp)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.ID == emp.ID {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Fprintf(w, "Employee with ID %v has been deleted", emp.ID)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}
