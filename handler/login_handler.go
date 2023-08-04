package handler

import (
	"goclean/model"
	"goclean/usecase"
	"goclean/utils/authutils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	userUc usecase.UserUsecase
}

func (l LoginHandler) login(ctx *gin.Context) {
	loginUserName := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&loginUserName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	usr, errGetName := l.userUc.GetUserByName(loginUserName.Username)
	if usr == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name is Invalid",
		})
		return
	}
	if errGetName != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginUserName.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": err.Error(),
		})
		return
	}

	temp, err := authutils.GenerateToken(loginUserName.Username)
	if err != nil {
		log.Println("Tokne Invalid")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": temp,
	})
}

func NewLoginHandler(lg *gin.Engine, loginUsecase usecase.UserUsecase) {
	loginHandler := &LoginHandler{
		userUc: loginUsecase,
	}
	lg.POST("/login", loginHandler.login)
}
