package repository

import (
	"SplitWise/internal/domain/entity"
	"SplitWise/internal/logger"
	"sync"
)

type UserRepository struct {
	Users map[string]*entity.User
	Mutex sync.RWMutex
}

func (r *UserRepository) RegisterUser(userID, name, email string) *entity.User {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	if user, exists := r.Users[userID]; exists {
		return user
	}
	user := entity.User{ID: userID, Name: name, Email: email}
	r.Users[userID] = &user
	logger.GetLogger("UserRepository").Infof("Registered user %s", userID)
	return &user
}

func (r *UserRepository) GetUser(userId string) (*entity.User, bool) {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	user, ok := r.Users[userId]
	return user, ok
}
