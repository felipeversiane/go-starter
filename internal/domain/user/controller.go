package user

import (
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

}

func (controller *userController) GetOneController(c *gin.Context) {

}
func (controller *userController) GetAllController(c *gin.Context) {

}

func (controller *userController) UpdateController(c *gin.Context) {

}

func (controller *userController) DeleteController(c *gin.Context) {

}
