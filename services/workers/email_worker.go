package workers

// EmailPayload - payload required to send email
type EmailPayload struct {
	FirstName     string
	LastName      string
	Body          string
	RecieverEmail string
	SenderEmail   string
	Subject       string
}

// SendEmail - enqueues and sends email
func SendEmail(queue string) {

}
