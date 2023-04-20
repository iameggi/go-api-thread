func TestGetEmployee(t *testing.T) {
	employeeID := "123"
	employeeName := "John"
	employeeEmail := "john@example.com"
	employeeAddress := "123 Main Street"
	employeeAge := 30

	employee := Employee{
		ID:      employeeID,
		Name:    employeeName,
		Email:   employeeEmail,
		Address: employeeAddress,
		Age:     employeeAge,
	}

	employeeStore := NewEmployeeStore()
	employeeStore.CreateEmployee(employee)

	tests := []struct {
		name          string
		employeeID    string
		expectedError error
	}{
		{
			name:          "Employee exists",
			employeeID:    employeeID,
			expectedError: nil,
		},
		{
			name:          "Employee does not exist",
			employeeID:    "456",
			expectedError: ErrEmployeeNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := employeeStore.GetEmployee(test.employeeID)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestCreateEmployee(t *testing.T) {
	employeeID := "123"
	employeeName := "John"
	employeeEmail := "john@example.com"
	employeeAddress := "123 Main Street"
	employeeAge := 30

	employee := Employee{
		ID:      employeeID,
		Name:    employeeName,
		Email:   employeeEmail,
		Address: employeeAddress,
		Age:     employeeAge,
	}

	employeeStore := NewEmployeeStore()

	tests := []struct {
		name          string
		employee      Employee
		expectedError error
	}{
		{
			name:          "Create new employee",
			employee:      employee,
			expectedError: nil,
		},
		{
			name: "Create employee with existing ID",
			employee: Employee{
				ID:      employeeID,
				Name:    "Jane",
				Email:   "jane@example.com",
				Address: "456 Main Street",
				Age:     40,
			},
			expectedError: ErrEmployeeAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := employeeStore.CreateEmployee(test.employee)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
