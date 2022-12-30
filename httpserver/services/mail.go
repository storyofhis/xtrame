package services

import (
	"net/http"

	"github.com/storyofhis/xtrame/httpserver/controllers/views"
)

type emailSvc struct {
	to      []string
	cc      []string
	subject string
	message string
}

func NewEmailSvc(to []string, cc []string, subject string, message string) EmailSvc {
	return &emailSvc{
		to:      to,
		cc:      cc,
		subject: subject,
		message: message,
	}
}

func (svc *emailSvc) ConnectEmail() *views.Response {
	// err := config.ConnectMail(svc.to, svc.cc, svc.subject, svc.message)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return views.ErrorResponse(http.StatusInternalServerError, views.M_INVALID_CREDENTIALS, err)
	// }
	// log.Println("email sent!")
	return views.SuccessResponse(http.StatusOK, views.M_OK, nil)
}
