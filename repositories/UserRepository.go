package repositories

import (
	"github.com/Leeroyakbar/bowlnow-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindById(id uuid.UUID) (*models.User, error)
	FindByUserName(userName string) (*models.User, error)
	FindAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) FindById(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := repo.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByUserName(userName string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("user_name = ?", userName).First(&user).Error
	return &user, err
}

func (repo *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}