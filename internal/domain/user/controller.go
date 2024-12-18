package user

import (
	"context"
	"net/http"
	"time"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
	"github.com/felipeversiane/go-starter/internal/infra/config/validation"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service UserServiceInterface
}

type UserControllerInterface interface {
	InsertOneController(c *gin.Context)
	GetOneByIDController(c *gin.Context)
	GetOneByEmailController(c *gin.Context)
	GetAllController(c *gin.Context)
	UpdateController(c *gin.Context)
	DeleteController(c *gin.Context)
}

func NewUserController(service UserServiceInterface) UserControllerInterface {
	return &userController{service}
}

func (controller *userController) InsertOneController(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationError := validation.ValidateError(err)
		c.JSON(validationError.Code, validationError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := controller.service.InsertOneService(req, ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.NewSuccessResponse(http.StatusCreated, gin.H{"id": id})
	c.JSON(response.Code, response)
}

func (controller *userController) GetOneByIDController(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		validationError := response.NewBadRequestError("ID is required")
		c.JSON(validationError.Code, validationError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := controller.service.GetOneByIDService(id, ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.NewSuccessResponse(http.StatusOK, user)
	c.JSON(response.Code, response)
}

func (controller *userController) GetOneByEmailController(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		validationError := response.NewBadRequestError("Email is required")
		c.JSON(validationError.Code, validationError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user, err := controller.service.GetOneByEmailService(email, ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.NewSuccessResponse(http.StatusOK, user)
	c.JSON(response.Code, response)
}

func (controller *userController) GetAllController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	users, err := controller.service.GetAllService(ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.NewSuccessResponse(http.StatusOK, users)
	c.JSON(response.Code, response)
}

func (controller *userController) UpdateController(c *gin.Context) {

	var req UserUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationError := validation.ValidateError(err)
		c.JSON(validationError.Code, validationError)
		return
	}

	id := c.Param("id")
	if id == "" {
		validationError := response.NewBadRequestError("ID is required")
		c.JSON(validationError.Code, validationError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := controller.service.UpdateService(id, req, ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func (controller *userController) DeleteController(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		validationError := response.NewBadRequestError("ID is required")
		c.JSON(validationError.Code, validationError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := controller.service.DeleteService(id, ctx)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
