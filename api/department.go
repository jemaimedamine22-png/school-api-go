package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemaimedamine22-png/school-api/db"
)

type createDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createDepartment(ctx *gin.Context) {
	var req createDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateDepartmentParams{
		Name: req.Name,
	}

	department, err := server.store.CreateDepartment(ctx, arg.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, department)
}

// تم تعديل النوع هنا إلى int32 ليتطابق مع توليد sqlc لجدول departments
type getDepartmentRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getDepartment(ctx *gin.Context) {
	var req getDepartmentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := server.store.GetDepartment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "department not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, department)
}

func (server *Server) listDepartments(ctx *gin.Context) {
	departments, err := server.store.ListDepartments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, departments)
}
	
