package controller

import (
	"fmt"
	"kemiskinan/helper"
	"kemiskinan/model"
	"kemiskinan/request"
	"kemiskinan/responses"
	"kemiskinan/service"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) GetUsers(cntx *gin.Context) {
	var users, err = c.userService.FindAll()
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
	}

	var usersResponse []responses.UserResponse

	for _, user := range users {
		var userResponse = helper.ConvertToUserResponse(user)
		usersResponse = append(usersResponse, userResponse)
	}

	cntx.JSON(http.StatusOK, usersResponse)
}

func (c *userController) GetUser(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	var user, err = c.userService.FindById(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak ditemukan",
		})
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, userResponse)
}

func (c *userController) CreateUser(cntx *gin.Context) {
	var userRequest request.CreateUserRequest

	var err = cntx.ShouldBindJSON(&userRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := c.userService.Create(userRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": cntx.Error(err),
		})
		return
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusCreated, gin.H{
		"data": userResponse,
	})
}

func (c *userController) UpdateUser(cntx *gin.Context) {
	var userRequest request.UpdateUserRequest

	var err = cntx.ShouldBindJSON(&userRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	user, err := c.userService.Update(id, userRequest)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"errors": cntx.Error(err),
		})
		return
	}

	var userResponse = helper.ConvertToUserResponse(user)

	cntx.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (c *userController) DeleteUser(cntx *gin.Context) {
	var idString = cntx.Param("id")
	var id, _ = strconv.Atoi(idString)

	_, err := c.userService.Delete(id)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Data gagal dihapus",
		})
		return
	}

	cntx.JSON(http.StatusOK, gin.H{
		"status": "Data berhasil dihapus",
	})
}

func (c *userController) LoginUser(cntx *gin.Context) {
	var loginRequest request.LoginUserRequest

	var err = cntx.ShouldBindJSON(&loginRequest)
	if err != nil {
		var errorMessages = []string{}

		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage = fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	user, err := c.userService.Login(loginRequest.Username)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Kesalahan pada username atau password",
		})

		return
	}

	var comparePassword = helper.CheckPasswordHash(user.Password, loginRequest.Password)
	if !comparePassword {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Kesalahan pada username atau password",
		})

		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uId":       user.Id,
		"uUsername": user.Username,
		"uEmail":    user.Email,
		"uNoHP":     user.NoHp,
		"uRole":     user.Role,

		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("APP_SECRET_KEY")))
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal generate token",
		})

		return
	}

	//send cookies back
	cntx.SetSameSite(http.SameSiteLaxMode)
	//false on secure (6th parameter) maybe want to change it to "true" when it goes online
	cntx.SetCookie("Authorization", accessTokenString, 3600*1, "", "", false, true)

	cntx.JSON(http.StatusOK, gin.H{})

}

func (r *userController) Validate(cntx *gin.Context) {
	var user, _ = cntx.Get("user")

	var userResponse = helper.ConvertToUserResponse(user.(model.User))

	// var userRole = user.(model.User).Role

	cntx.JSON(http.StatusOK, gin.H{
		"user": userResponse,
	})
}
