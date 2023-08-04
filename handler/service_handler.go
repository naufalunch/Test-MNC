package handler

import (
	"errors"
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	svcUsecase usecase.ServiceUsecase
}

func (svcHandler ServiceHandler) GetServiceById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	svc, err := svcHandler.svcUsecase.GetServiceById(id)
	if err != nil {
		fmt.Printf("ServiceHandler.GetServiceById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    svc,
	})
}

func (svcHandler ServiceHandler) InsertService(ctx *gin.Context) {
	svc := &model.ServiceModel{}
	err := ctx.ShouldBindJSON(&svc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(svc.Name) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
		})
		return
	}

	err = svcHandler.svcUsecase.InsertService(svc)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.InsertService() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ServiceHandler.InsertService() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data service",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) *ServiceHandler {
	svcHandler := &ServiceHandler{
		svcUsecase: svcUsecase,
	}
	srv.GET("/service/:id", RequireToken(), svcHandler.GetServiceById)
	srv.POST("/service", svcHandler.InsertService)
	return svcHandler
}
