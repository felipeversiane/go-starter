package user

import (
	"github.com/felipeversiane/go-starter/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup, db database.DatabaseInterface) *gin.RouterGroup {
	controller := NewUserController(NewUserService(NewUserRepository(db)))

	user := g.Group("/user")
	{
		user.POST("/", controller.InsertOneController)
		user.GET("/:id", controller.GetOneByIDController)
		user.GET("/get_user_by_email/:email", controller.GetOneByEmailController)
		user.GET("/", controller.GetAllController)
		user.DELETE("/:id", controller.DeleteController)
		user.PUT("/:id", controller.UpdateController)

	}

	return user
}
