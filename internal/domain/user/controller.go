package user

import (
	"context"
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
	GetOneController(c *gin.Context)
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

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := controller.service.InsertOneService(req, ctxTimeout)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.NewCreatedResponse("User created successfully", gin.H{"id": id})
	c.JSON(response.Code, response)
}

func (controller *userController) GetOneController(c *gin.Context) {

}
func (controller *userController) GetAllController(c *gin.Context) {

}

func (controller *userController) UpdateController(c *gin.Context) {

}

func (controller *userController) DeleteController(c *gin.Context) {

}
