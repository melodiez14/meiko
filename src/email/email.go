package email

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/smtp"
	"os"

	"github.com/domodwyer/mailyak"
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

// Attachment is a struct used for add new attachment
type Attachment struct {
	Title string
	Data  io.Reader
}

// Request struct for sending an email
type Request struct {
	*mailyak.MailYak
}

var (
	auth smtp.Auth
	c    Config
)

func PrintConfig() {
	fmt.Println(c)
}

// Init is used for initialize a new email connection
func Init(cfg Config) {
	c = cfg
	auth = smtp.PlainAuth("", c.UserName, c.Password, c.Address)
}

// NewRequest is used for make a new email request
func NewRequest(to string, subject string) *Request {

	addr := fmt.Sprintf("%s:%d", c.Address, c.Port)
	r := &Request{
		mailyak.New(addr, auth),
	}
	r.To(to)
	r.Subject(subject)
	r.From(c.UserName)
	r.FromName("Meiko")
	return r
}

// NewAttachment is a function for create new attachment by using title and path
func NewAttachment(title, path string) Attachment {

	data, err := os.Open(path)
	if err != nil {
		log.Println("Error: ", err)
		return Attachment{}
	}

	return Attachment{
		Title: title,
		Data:  data,
	}
}

// SetTemplate is used for set an email html content. The parameters are templatePath which contain the path of email template and data is a struct of data which used in email template
func (r *Request) SetTemplate(templatePath string, data interface{}) *Request {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error parsing data to template")
		return r
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println("Error parsing template to buffer")
		return r
	}
	r.HTML().Set(buf.String())
	return r
}

// SetSender used for change the name of sender
func (r *Request) SetSender(name, email string) *Request {
	r.From(email)
	r.FromName(name)
	return r
}

// SetTo used for change the name of sender
func (r *Request) SetTo(email string) *Request {
	r.To(email)
	return r
}

// SetAttachment used to add an attachment of email by using map[string]string. Example map["My Photo"]"/etc/myphoto.png"
func (r *Request) SetAttachment(attachment []Attachment) *Request {
	for _, v := range attachment {
		r.Attach(v.Title, v.Data)
	}
	return r
}

// Deliver action to send an email
func (r *Request) Deliver() {
	go func() {
		if err := r.Send(); err != nil {
			log.Println(err)
		}
	}()
}
