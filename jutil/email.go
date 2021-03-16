package jutil

import (
	"gopkg.in/gomail.v2"
)

//邮件发送工具
type EmailUtil struct {
	Host     string
	Port     int
	Username string
	Password string

	Dialer *gomail.Dialer
}

//连接
func (u *EmailUtil) Connect() {
	u.Dialer = gomail.NewDialer(u.Host, u.Port, u.Username, u.Password)
}

//发送邮件
func (u *EmailUtil) Send(to string, subject string, body string) {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", u.Username)
	//接收人
	m.SetHeader("To", to)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", body)
	//附件
	//m.Attach("./myIpPic.png")
	// 发送邮件
	if err := u.Dialer.DialAndSend(m); err != nil {
		panic(err)
	}
}

//获取实例
func NewEmailUtil() *EmailUtil  {
	config:=NewConfig()
	host:=config.GetString("email.default.host")
	port:=config.GetInt("email.default.port")
	username:=config.GetString("email.default.username")
	password:=config.GetString("email.default.password")
	util:=&EmailUtil{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
	util.Connect()
	return util
}