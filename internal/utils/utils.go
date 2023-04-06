package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	configs "whatsapp-app/internal/config"
	models "whatsapp-app/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*/"

type IUtils interface {
	GetUser(c *echo.Context) models.User
	EqualPassword(old, new string) bool
	GeneratePassword(password string) (string, error)
	SplitSchoolID(email string) int
	RandNumber(start int, end int) int
	SendEmail(toEmail string, code int) error
	GenerateRandomString(n int) string
}

type Utils struct{}

func NewUtils() IUtils {
	return &Utils{}
}

func (u *Utils) GetUser(c *echo.Context) models.User {
	user := *(*c).Get("verifiedUser").(*models.User)
	return user
}

func (u *Utils) EqualPassword(old, new string) bool {
	passwordControl := bcrypt.CompareHashAndPassword([]byte(old), []byte(new))
	if passwordControl != nil {
		return false
	}
	return true
}

func (u *Utils) GeneratePassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(hasPassword), err
}

func EuToTime(StringDate string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", StringDate)
	return date, err
}

func (u *Utils) SplitSchoolID(email string) int {

	var school_id int
	if strings.Contains(email, "@") == true {
		findSchoolID := strings.Split(email, "@")
		school_id, _ = strconv.Atoi(findSchoolID[0])
	} else {
		school_id, _ = strconv.Atoi(email)
	}
	return school_id
}

func (u *Utils) RandNumber(start int, end int) int {
	code := rand.Intn(end-start) + start
	return code
}

func (u *Utils) SendEmail(toEmail string, code int) error {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", configs.EmailAddress)

	// Set E-Mail receivers
	m.SetHeader("To", toEmail)

	// Set E-Mail subject

	body := "Doğrulama kodunuz:" + strconv.Itoa(code)
	subject := "Hesabınızı doğrulamak için gelen doğrulama e-postası"

	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.office365.com", 587, configs.EmailAddress, configs.EmailPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (u *Utils) GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
