package usecase

import (
	"errors"
	"time"

	"ezustore/src/domain"
	"ezustore/src/infrastructure/auth"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct{ repo domain.UserRepository }

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase { return &userUsecase{repo: repo} }

func (u *userUsecase) Register(name, email, password string) error {
	if existing, _ := u.repo.GetByEmail(email); existing != nil && existing.ID != 0 {
		return errors.New("email sudah terdaftar")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &domain.User{Name: name, Email: email, Password: string(hashed)}
	return u.repo.Register(user)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("password salah")
	}
	// Use shared token generator (auth.GenerateToken) OR build inline for clarity
	return auth.GenerateToken(user.ID)
}

func (u *userUsecase) GetProfile(userID uint) (*domain.User, error) {
	user, err := u.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (u *userUsecase) UpdateProfile(userID uint, name, email string) error {
	user, err := u.repo.GetByID(userID)
	if err != nil {
		return err
	}
	user.Name = name
	if email != "" {
		user.Email = email
	}
	return u.repo.Update(user)
}

// (opsional) contoh generate token inline bila tak mau pakai auth.GenerateToken
func generateJWTInline(userID uint, key []byte) (string, error) {
	claims := jwt.MapClaims{"user_id": userID, "exp": time.Now().Add(24 * time.Hour).Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
