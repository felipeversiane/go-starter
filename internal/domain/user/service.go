package user

type userService struct {
	repository UserRepositoryInterface
}

type UserServiceInterface interface {
}

func NewUserService(repository UserRepositoryInterface) UserServiceInterface {
	return &userService{repository}
}
