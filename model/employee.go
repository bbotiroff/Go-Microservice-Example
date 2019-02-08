package model

import (
	"github.com/gocql/gocql"
)

// Employee is database model
type Employee struct {
	EmployeeID gocql.UUID
	FirstName  string
	LastName   string
	Email      string
	Gender     string
}
