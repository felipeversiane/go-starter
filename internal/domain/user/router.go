package user

import (
	"github.com/felipeversiane/go-starter/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup, db database.DatabaseInterface) *gin.RouterGroup {
	controller := NewUserController(NewUserService(NewUserRepository(db)))

	person := g.Group("/user")
	{
		person.POST("/", controller.InsertOneController)

	}

	return person
}
