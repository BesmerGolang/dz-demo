package email

import (
	"dz3/configs"
	"net"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendVerificationEmail(toEmail string, token string, conf *configs.EmailConfig) error {
	//создаем новое письмо
	letter := email.NewEmail()
	letter.From = "ChechService <" + conf.Email + ">"                   //служебная почта(откуда направляем письмо)
	letter.To = []string{toEmail}                                       // кому отправить
	letter.Subject = "Vertify your email"                               // тема письмак
	letter.Text = []byte("Link: http://localhost:8081/verify/" + token) // тело письма
	host, _, err := net.SplitHostPort(conf.Address)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", conf.Email, conf.Password, host)
	err = letter.Send(conf.Address, auth)
	return err

}
