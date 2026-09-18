package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jemaimedamine22-png/school-api/db"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
    p.Use(router)


	// API of department
	router.POST("/departments", server.createDepartment)
	router.GET("/departments/:id", server.getDepartment)
	router.GET("/departments", server.listDepartments)
	// API of students
	router.POST("/students", server.createStudent)
	router.GET("/students/:id", server.getStudent)
	router.GET("/students", server.listStudents)
	// API of professor
	router.POST("/professors", server.createProfessor)
	router.GET("/professors/:id", server.getProfessor)
	router.GET("/professors", server.listProfessors)
	//metrics
	//router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}