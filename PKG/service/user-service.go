package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/20WilliamPanjaitan/golang_gin_gorm_GWT/dto"
	"github.com/20WilliamPanjaitan/golang_gin_gorm_GWT/entity"
	"github.com/20WilliamPanjaitan/golang_gin_gorm_GWT/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
