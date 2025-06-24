package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Leeroyakbar/bowlnow-backend/dto"
	"github.com/Leeroyakbar/bowlnow-backend/models"
	"github.com/Leeroyakbar/bowlnow-backend/repositories"
	"github.com/Leeroyakbar/bowlnow-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterFromForm(c *fiber.Ctx) (*dto.RegisterResponse, error)
	Login(userName string, password string) (*dto.LoginResponse, error)
	GetById(id uuid.UUID) (*models.User, error)
	GetByUserName(userName string) (*models.User, error)
	GetAll() ([]models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterFromForm(c *fiber.Ctx) (*dto.RegisterResponse, error) {

	form, err := c.MultipartForm()
	if err != nil {
		return nil, errors.New("invalid form data")
	}

	fileHeader := form.File["image"]
	if len(fileHeader) == 0 {
		return nil, errors.New("image is required")
	}

	values := form.Value
	userName := values["user_name"][0]
	fullName := values["full_name"][0]
	password := values["password"][0]
	roleID := values["role_id"][0]
	guestFlag, _ := strconv.Atoi(values["guest_flag"][0])

	existing, err := s.repo.FindByUserName(userName)
	if (existing != nil && existing.UserName == userName) || err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("username already exists")
	}

	// Simpan file
	file := fileHeader[0]
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := fmt.Sprintf("uploads/users/%s", fileName)
	if err := c.SaveFile(file, path); err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to save file")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		UserID:       uuid.New(),
		UserName:     userName,
		FullName:     fullName,
		Password:     string(hashed),
		RoleID:       uuid.MustParse(roleID),
		GuestFlag:    guestFlag,
		Image:        path,
		DeleteStatus: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	res := &dto.RegisterResponse{
		UserID:   user.UserID,
		RoleID:   user.RoleID,
		FullName: user.FullName,
		UserName: user.UserName,
	}
	return res, s.repo.Create(user)
}

func (service *userService) Login(userName string, password string) (*dto.LoginResponse, error) {
	user, err := service.repo.FindByUserName(userName)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// compare hashing password from db
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	// generate token
	token, err := utils.GenerateJWT(user.UserID, user.Role.RoleName)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	res := &dto.LoginResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		FullName: user.FullName,
		RoleName: user.Role.RoleName,
		Token:    token,
	}
	return res, nil
}

func (service *userService) GetById(id uuid.UUID) (*models.User, error) {
	return service.repo.FindById(id)
}

func (service *userService) GetByUserName(userName string) (*models.User, error) {
	return service.repo.FindByUserName(userName)
}

func (service *userService) GetAll() ([]models.User, error) {
	return service.repo.FindAll()
}
