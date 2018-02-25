package emails

import (
	"testing"

	"github.com/benmanns/goworker"
)

var testEmail = EmailPayload{
	FirstName:     "Test",
	LastName:      "Tester",
	TextBody:      "This is a test.",
	HTMLBody:      "<p>this is a test</p>",
	RecieverEmail: "tim.kellogg@gmail.com",
	Subject:       "Email Test",
}

func TestSendEmail(t *testing.T) {
	InitWorkers()

	if err := goworker.Work(); err != nil {
		t.Errorf("Worker failed to start: %v", err)
	}

	err := SendEmail("default", testEmail)
	if err != nil {
		t.Error(err)
	}
}
