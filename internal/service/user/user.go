package user

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
	"whatsapp-app/cache"
	"whatsapp-app/dto/request"
	"whatsapp-app/dto/response"
	configs "whatsapp-app/internal/config"
	"whatsapp-app/internal/repository"
	"whatsapp-app/internal/service"
	jwtPackage "whatsapp-app/internal/utils/jwt"
	models "whatsapp-app/model"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Register(ctx context.Context, userRegisterRequest request.UserRegisterDTO) (response.UserRegisterDTO, error)
	Login(ctx context.Context, userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error)
	//SendVerifyEmail(ctx context.Context, userVerifyRequest request.UserVerifyDTO) (response.UserVerifyDTO, error)
	VerifyUserEmail(ctx context.Context, userVerifyRequest request.UserVerifyEmailDTO) (response.UserVerifyDTO, error)
}

//TODO: Service 'e bakılacak
type UserService struct {
	repository repository.IUserRepository
	service.Service
	cache cache.ICache
}

func NewUserService(repository repository.IUserRepository, service service.Service, cache cache.ICache) IUserService {
	return &UserService{repository: repository, Service: service, cache: cache}
}

func (s *UserService) Register(ctx context.Context, request request.UserRegisterDTO) (response.UserRegisterDTO, error) {

	findSchoolID := strings.Split(request.Email, "@")
	var school_id int

	school_id, err := strconv.Atoi(findSchoolID[0])
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("Kullanıcı oluşturlamadı lütfen tekrar deneyiniz")
	}

	isExist := s.repository.IsDuplicateSchoolID(int32(school_id))
	if isExist {
		return response.UserRegisterDTO{}, errors.New("Kayıt olmak istediğiniz kullanıcı zaten var.")
	}

	apiKey := s.Utils.GenerateRandomString(40)

	var hashPassword string
	hashPassword, err = s.Utils.GeneratePassword(request.Password)
	if err != nil {
		err := errors.New("Şifre oluşturulamadı.")
		return response.UserRegisterDTO{}, err
	}

	passwordControl := s.Utils.EqualPassword(hashPassword, request.Password)
	if !passwordControl {
		return response.UserRegisterDTO{}, errors.New("Şifre doğru bir şekilde oluşturulamadı.")
	}

	newUser := models.User{
		Name:     request.Name,
		Password: hashPassword,
		SchoolID: int32(school_id),
		Email:    request.Email,
		ApiKey:   apiKey,
	}
	err = s.repository.CreateUser(&newUser)
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("Kullanıcı oluşturlamadı lütfen tekrar deneyiniz ")
	}

	responses := response.UserRegisterDTO{
		Name:     newUser.Name,
		Email:    newUser.Email,
		SchoolID: newUser.SchoolID,
	}

	code := s.Utils.RandNumber(100000, 999999)

	key := "user-email:" + request.Email + "code:" + strconv.Itoa(code)
	err = s.cache.Set(ctx, key, code, 180)
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("Oluşturulan kod kaydedilirken hata meydana geldi lütfen tekrar deneyiniz")
	}

	err = s.Utils.SendEmail(request.Email, code)
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("E-Posta gönderilirken bir hata meydana geldi.")
	}

	return responses, nil

}

func (s *UserService) Login(ctx context.Context, request request.UserLoginDTO) (response.UserLoginDTO, error) {

	school_id := s.Utils.SplitSchoolID(request.Email)
	user, err := s.repository.Login(int32(school_id))
	if err != nil {
		return response.UserLoginDTO{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return response.UserLoginDTO{}, errors.New("Kullanıcı adı veya parola hatalı lütfen tekrar deneyiniz.")
	}

	if !user.Verified {

		code := s.Utils.RandNumber(100000, 999999)
		key := "user-email:" + request.Email + "code:" + strconv.Itoa(code)
		err = s.cache.Set(ctx, key, code, 180)
		if err != nil {
			return response.UserLoginDTO{}, errors.New("Oluşturulan kod kaydedilirken hata meydana geldi lütfen tekrar deneyiniz")
		}

		err = s.Utils.SendEmail(request.Email, code)
		if err != nil {
			return response.UserLoginDTO{}, errors.New("E-Posta gönderilirken bir hata meydana geldi.")
		}
		var userLoginResponse response.UserLoginDTO
		userLoginResponse.Convert(user, "")

		return userLoginResponse, errors.New("Giriş yapmak için lütfen önce e-mail doğrulaması yapınız.")
	}

	claims := &jwtPackage.JwtCustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(configs.TokenSecret))
	if err != nil {
		return response.UserLoginDTO{}, err
	}

	var userLoginResponse response.UserLoginDTO
	userLoginResponse.Convert(user, tokenSigned)
	return userLoginResponse, nil
}

/*
func (s *UserService) SendVerifyEmail(ctx context.Context, request request.UserVerifyDTO) (response.UserVerifyDTO, error) {
	code := s.Utils.RandNumber(100000, 999999)

	isExist := s.repository.IsExistWithEmail(request.Email)
	if !isExist {
		return response.UserVerifyDTO{}, errors.New("Kullanıcı bulunamadı.")
	}

	key := "user-email:" + request.Email + "code:" + strconv.Itoa(code)
	err := s.cache.Set(ctx, key, code, 180)
	if err != nil {
		return response.UserVerifyDTO{}, errors.New("Oluşturulan kod kaydedilirken hata meydana geldi lütfen tekrar deneyiniz")
	}

	err = s.Utils.SendEmail(request.Email, code)
	if err != nil {
		return response.UserVerifyDTO{}, errors.New("E-Posta gönderilirken bir hata meydana geldi.")
	}
	message := "E-Posta başarıyla gönderildi."

	return response.UserVerifyDTO{Message: message}, nil
}*/

func (s *UserService) VerifyUserEmail(ctx context.Context, request request.UserVerifyEmailDTO) (response.UserVerifyDTO, error) {

	key := "user-email:" + request.Email + "code:" + request.Code
	isExist := s.cache.Get(ctx, key)
	if !isExist {
		return response.UserVerifyDTO{}, errors.New("Girilen kod hatalı veya süresi dolmuş.")
	}

	user, err := s.repository.FindUserWithEmail(request.Email)
	if err != nil {
		return response.UserVerifyDTO{}, errors.New("Kullanıcı bulunamadı")
	}
	user.Verified = true

	err = s.repository.UpdateUser(user)
	if err != nil {
		return response.UserVerifyDTO{}, errors.New("Kullanıcı güncellenemedi")
	}
	return response.UserVerifyDTO{
		Message: "Kullanıcı başarıyla doğrulandı, giriş yapabilirsiniz.",
	}, nil
}
