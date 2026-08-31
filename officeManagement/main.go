package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Email       string `json:"email"`
}

var employees = []Employee{
	{ID: 1, Name: "Rahul", Department: "IT", Designation: "Developer", Email: "rahul@gmail.com"},
	{ID: 2, Name: "Aman", Department: "HR", Designation: "Manager", Email: "aman@gmail.com"},
}

func main() {
	r := gin.Default()

	r.GET("/employees", getEmployees)
	r.GET("/employees/:id", getEmployeeByID)
	r.POST("/employees", addEmployee)
	r.PUT("/employees/:id", updateEmployee)
	r.DELETE("/employees/:id", deleteEmployee)

	r.Run(":8080")
}

func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func getEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, emp := range employees {
		if emp.ID == id {
			c.JSON(http.StatusOK, emp)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

func addEmployee(c *gin.Context) {
	var newEmployee Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	c.JSON(http.StatusCreated, newEmployee)
}

func updateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedEmployee Employee

	if err := c.BindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			updatedEmployee.ID = id
			employees[i] = updatedEmployee
			c.JSON(http.StatusOK, updatedEmployee)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}

func deleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
}
