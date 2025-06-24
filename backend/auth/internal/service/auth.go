package service

import (
	"crypto/tls"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/errors"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/hash"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/jwt"
	"github.com/Bobr-Lord/react-go-shop/tree/main/backend/auth/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
)

func (s *Service) Register(req *models.RegisterRequest) (string, error) {
	hashPass, err := hash.HashPassword(req.Password)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.Password = hashPass
	token := uuid.New().String()
	//err = SendConfirmationEmailSendGrid(req.Email, token)
	go func() {
		err := SendEmailSMTP(req.Email, "Подтверждение", GetTemplate(token))
		if err != nil {
			logrus.Error(err)
		}
	}()

	id, err := s.repo.Register(req, token)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Login(req *models.LoginRequest) (string, error) {
	id, role, err := s.repo.Login(req)
	if err != nil {
		return "", err
	}

	privetKey, err := jwt.LoadRSAPrivateKeyPKCS8(s.cfg.PathPrivateKey)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	token, err := jwt.GenerateToken(id, role, privetKey)
	if err != nil {
		return "", errors.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return token, nil
}

func (s *Service) GetMe(id string) (*models.GetMeResponse, error) {
	return s.repo.GetMe(id)
}

func (s *Service) VerifyEmail(token string) error {
	return s.repo.VerifyEmail(token)
}

func SendEmailSMTP(toEmail, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "react.go.shop@gmail.com")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	pass := os.Getenv("EMAIL_API_PASS")
	d := gomail.NewDialer("smtp.gmail.com", 587, "react.go.shop@gmail.com", pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // ⚠️ для разработки

	return d.DialAndSend(m)
}

//func SendConfirmationEmailSendGrid(toEmail, token string) error {
//	fmt.Println("sending email to "+toEmail, token)
//	from := mail.NewEmail("React-Go Shop", "react.go.shop@gmail.com")
//	subject := "Подтверждение регистрации"
//	to := mail.NewEmail("", toEmail)
//
//	// Текст письма
//	plainText := "Подтвердите свою почту: http://localhost:8080/verify?token=" + token
//	htmlText := fmt.Sprintf(`
//		<p>Здравствуйте!</p>
//		<p>Благодарим за регистрацию в <strong>React-Go Shop</strong>.</p>
//		<p>Нажмите <a href="http://localhost:8080/verify?token=%s">сюда</a>, чтобы подтвердить ваш email.</p>
//		<p>Если вы не регистрировались, просто проигнорируйте это письмо.</p>
//	`, token)
//
//	// Создаём сообщение
//	message := mail.NewSingleEmail(from, subject, to, plainText, htmlText)
//
//	// Получаем ключ из окружения (заранее положи туда свой API ключ)
//	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
//
//	// Отправка
//	response, err := client.Send(message)
//	if err != nil {
//		return err
//	}
//	if response.StatusCode >= 400 {
//		return fmt.Errorf("sendgrid error: %s", response.Body)
//	}
//
//	return nil
//}
