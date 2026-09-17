package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jemaimedamine22-png/school-api/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	// API of department
	router.POST("/departments", server.createDepartment)
	router.GET("/departments/:id", server.getDepartment)
	router.GET("/departments", server.listDepartments)
	// API of students
	router.POST("/students", server.createStudent)
	router.GET("/students/:id", server.getStudent)
	router.GET("/students", server.listStudents)
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}