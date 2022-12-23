package usecase

import (
	"backend/model"
	"backend/repository"
)

func V1GetUser(userRepository repository.User, id int) (*model.User, error) {
	return userRepository.GetUser(id)
}

func V1PostUser(userRepository repository.User, user *model.User) (int, error) {
	return userRepository.PostUser(user)
}
