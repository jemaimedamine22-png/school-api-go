package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func getDepartments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all departments"})
}

func createDepartment(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Department created"})
}

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all students"})
}

func createStudent(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Student created"})
}

func getProfessors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all professors"})
}

func createProfessor(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Professor created"})
}