package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemaimedamine22-png/school-api/db"
)

type createProfessorRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int32  `json:"department_id" binding:"required,min=1"`
}

func (server *Server) createProfessor(ctx *gin.Context) {
	var req createProfessorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateProfessorParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		DepartmentID: sql.NullInt32{Int32: req.DepartmentID, Valid: true},
	}

	professor, err := server.store.CreateProfessor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professor)
}

type getProfessorRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProfessor(ctx *gin.Context) {
	var req getProfessorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	professor, err := server.store.GetProfessor(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "professor not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professor)
}

type listProfessorsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProfessors(ctx *gin.Context) {
	var req listProfessorsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListProfessorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	professors, err := server.store.ListProfessors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, professors)
}