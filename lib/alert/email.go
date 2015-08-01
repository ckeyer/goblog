// 通过邮件的方式告警
// 默认的发送邮箱是sf_monitor@163.com
package alert

import (
	"fmt"
	"net/smtp"
	"strings"
)

var (
	email *EmailAlerter
)

type EmailAlerter struct {
	auth smtp.Auth

	ContentType string
	User        string
	Password    string
	Host        string
	ToList      string
}

func NewEmailAlerter(user, password, host, content_type string, tolist []string) *EmailAlerter {
	hp := strings.Split(host, ":")
	switch content_type {
	case "html":
		content_type = "Content-Type: text/html; charset=UTF-8"
	default:
		content_type = "Content-Type: text/html; charset=UTF-8"
	}
	e := &EmailAlerter{
		User:        user,
		Password:    password,
		Host:        host,
		ContentType: content_type,
		ToList:      tolist,
	}
	e.auth = smtp.PlainAuth("", user, password, hp[0])
	return e
}
func NewDefaultEmailAlerter() *EmailAlerter {
	if email == nil {
		NewEmailAlerter("sf_monitor@163.com", "cqwsmdgupjmareyc", "smtp.163.com:25",
			"text", []string{"me@ckeyer.com", "ckeyer@yeah.net", "cjstudio@yeah.net"})
	}
	return email
}

func (e *EmailAlerter) Alert(lev int, msg string) error {
	switch lev {
	case ALERT_LEVEL_MINOR:
		subject := alert_title[ALERT_LEVEL_MINOR]
		return sendMail(subject, msg)
	case ALERT_LEVEL_MAJOR:
		subject := alert_title[ALERT_LEVEL_MAJOR]
		return sendMail(subject, msg)
	case ALERT_LEVEL_CRITICAL:
		subject := alert_title[ALERT_LEVEL_CRITICAL]
		return sendMail(subject, msg)
	// 默认为警告级别
	default:
		subject := alert_title[ALERT_LEVEL_WARING]
		return sendMail(subject, msg)
	}
}

func (e *EmailAlerter) sendMail(subject, body string) error {
	to := ""
	for i, v := range send_to {
		if i == 0 {
			to = v
		} else {
			to += ";" + v
		}
	}
	msg := []byte("To: " + to + "\r\nFrom: " + e.User + "<" + e.User + ">\r\nSubject: " + subject + "\r\n" +
		e.ContentType + "\r\n\r\n" + body)
	err := smtp.SendMail(e.Host, e.auth, e.User, e.ToList, msg)
	return err
}
