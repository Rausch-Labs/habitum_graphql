package services

import (
	"errors"
	"gorm.io/gorm"
	"github.com/suisuss/habitum_graphQL/models"
	"strings"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)
type UserServiceI interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUserByID(id string) error
	GetUserByID(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	CheckIfUserExists(username string) error
	CheckPasswordHash(hash, password string) bool
	HashPassword(password string) (string, error)
	GenerateJWT(id, username string) (string, error)
	ParseJWT(tokenString string) (jwt.MapClaims, error)
}

type UserServiceS struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceS {
	return &UserServiceS{
		db,
	}
}

func (s *UserServiceS) CreateUser(user *models.User) (*models.User, error) {


	err := s.db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("user [SOMETHING] already taken")
		}
		return nil, err
	}

	return user, nil

}

func (s *UserServiceS) UpdateUser(user *models.User) (*models.User, error) {

	err := s.db.Save(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("user [SOMETHING] already taken")
		}
		return nil, err
	}

	return user, nil

}

func (s *UserServiceS) DeleteUserByID(id string) error {

	user := &models.User{}

	err := s.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceS) GetUserByID(id string) (*models.User, error) {

	user := &models.User{}

	err := s.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceS) GetUserByUsername(username string) (*models.User, error) {

	user := &models.User{}

	err := s.db.Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceS) CheckIfUserExists(username string) error {
	// Check if username exists
	err := s.db.Where("username = ?", username).Error
	if err != nil {
		return err
	}
	return err
}

func (s *UserServiceS) GetAllUsers() ([]*models.User, error) {

	var users []*models.User

	fmt.Println("here")

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserServiceS) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	return string(bytes), err
}
func (s *UserServiceS) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Parsing token
func (s *UserServiceS) ParseJWT(tokenString string) (jwt.MapClaims, error) {
	var JWTSECRETKEY = os.Getenv("JWTSECRETKEY")
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Don't forget to validate the alg is what you expect:
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
    return []byte(JWTSECRETKEY), nil
  })
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return claims,nil
  } 
  return nil,err
}

func (s *UserServiceS) GenerateJWT(id, username string) (string, error) {
	//Generating token
	var (
		maxAge = 60*60*2
		JWTSECRETKEY = os.Getenv("JWTSECRETKEY")
	)

  claims := jwt.MapClaims{
		"user-id": id,
    "username": username,
    "exp": time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // expiration time, must be set,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString([]byte(JWTSECRETKEY))
	
  if err != nil {
    fmt.Println(err)
  }

	return tokenString, err
}