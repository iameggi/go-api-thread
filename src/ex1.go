package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Address string `json:"address"`
	Age    int    `json:"age"`
}

var employees []Employee
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", addEmployee)
	http.HandleFunc("/employee/", updateEmployee)
	http.HandleFunc("/employee/delete/", deleteEmployee)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mutex.Lock()
		employees = append(employees, employee)
		mutex.Unlock()
		json.NewEncoder(w).Encode(employee)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		var employee Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, emp := range employees {
			if emp.ID == employee.ID {
				mutex.Lock()
				employees[i] = employee
				mutex.Unlock()
				json.NewEncoder(w).Encode(employee)
				return
			}
		}
		http.Error(w, "Employee not found", http.StatusNotFound)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Path[len("/employee/delete/"):]
		for i, emp := range employees {
			if emp.ID == id {
				mutex.Lock()
				employees = append(employees[:i], employees[i+1:]...)
				mutex.Unlock()
				return
			}
		}
		http.Error(w, "Employee not found", http.StatusNotFound)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
