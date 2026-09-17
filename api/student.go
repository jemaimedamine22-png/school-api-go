package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemaimedamine22-png/school-api/db"
)

type createStudentRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int32  `json:"department_id" binding:"required,min=1"`
}

func (server *Server) createStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// استخدام الـ Struct الذي تفضله وتولده sqlc تلقائياً لوجود أكثر من حقل
	arg := db.CreateStudentParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		DepartmentID: sql.NullInt32{Int32: req.DepartmentID, Valid: true},
	}

	student, err := server.store.CreateStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

type getStudentRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStudent(ctx *gin.Context) {
	var req getStudentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := server.store.GetStudent(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

type listStudentsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listStudents(ctx *gin.Context) {
	var req listStudentsRequest
	// استخدام ShouldBindQuery لأننا نستقبل البيانات عبر الـ Query Parameters مثل ?page_id=1&page_size=5
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListStudentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	students, err := server.store.ListStudents(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}