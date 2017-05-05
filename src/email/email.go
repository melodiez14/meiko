package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

// ************************************************* //

// func HowToUse() {
// 	templateData := struct {
// 		Name string
// 		URL  string
// 	}{
// 		Name: "Dhanush",
// 		URL:  "http://geektrust.in",
// 	}
// 	to := []string{
// 		"risal@live.com",
// 		"asepnurisk@gmail.com",
// 	}
// 	NewRequest(to, "This is subject!").
// 		SetTemplate("template.html", templateData).
// 		SendEmail()
// }

// https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d

// ************************************************* //

// Config is a struct for initialize the configuration
type Config struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     int16  `json:"port"`
}

// Request struct for sending an email
type Request struct {
	from    string
	to      string
	subject string
	body    string
}

var (
	auth smtp.Auth
	c    Config
)

// Init is used for initialize a new email connection
func Init(cfg Config) {
	c = cfg
	auth = smtp.PlainAuth("", c.UserName, c.Password, c.Address)
	templateData := struct{}{}
	to := "risal@live.com"

	NewRequest(to, "This is from meiko!", "template.html", templateData).
		Send()
}

// NewRequest is used for make a new email request
func NewRequest(to string, subject, templateFileName string, data interface{}) *Request {

	r := &Request{
		to:      to,
		subject: subject,
		from:    "Meiko <idms.notification@gmail.com>",
	}

	t, err := template.New("new").Parse("<html><body>Hello, this is the body</body></html>")
	if err != nil {
		log.Println("Error parsing data to template")
		return r
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println("Error parsing template to buffer")
		return r
	}

	r.body = buf.String()

	return r
}

// SetSender used for change the name of sender
func (r *Request) SetSender(name, email string) *Request {
	s := mail.Address{
		Name:    "",
		Address: "",
	}
	r.from = s.String()
	return r
}

// SetTo used for change the name of sender
func (r *Request) SetTo(name, email string) *Request {
	t := mail.Address{
		Name:    "",
		Address: "",
	}
	r.to = t.String()
	return r
}

// Send action to send an email
func (r *Request) Send() {

	if r == nil {
		return
	}

	addr := fmt.Sprintf("%s:%d", c.Address, c.Port)
	header := make(map[string]string)
	header["From"] = r.from
	header["To"] = r.to
	header["Subject"] = r.subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=utf-8"
	header["Content-Transfer-Encoding"] = "base64"

	b := ""
	for k, v := range header {
		b += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	b += "\r\n" + base64.StdEncoding.EncodeToString([]byte(r.body))

	msg := []byte(b)
	go func() {
		if err := smtp.SendMail(addr, auth, c.UserName, []string{r.to}, msg); err != nil {
			log.Println("Error to send email")
		}
	}()
}
