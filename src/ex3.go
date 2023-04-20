package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Model Employee
type Employee struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Address string    `json:"address"`
	Age    int       `json:"age"`
}

var employees []Employee
var mu sync.Mutex


func getEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(employees)
}


func getEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}


func createEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = uuid.New()
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}


func updateEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID.String() == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = item.ID
			employees = append(employees, employee)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	json.NewEncoder(w).Encode(employees)
}


func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID.String() == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {

	r := mux.NewRouter()

	
	employees = append(employees, Employee{ID: uuid.New(), Name: "John Doe", Email: "john.doe@example.com", Address: "123 Main St", Age: 30})
	employees = append(employees, Employee{ID: uuid.New(), Name: "Jane Doe", Email: "jane.doe@example.com", Address: "456 Second St", Age: 35})

	
	r.HandleFunc("/api/employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/employees", createEmployee).Methods("POST")
	r.HandleFunc("/api/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", deleteEmployee).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", r))
}
