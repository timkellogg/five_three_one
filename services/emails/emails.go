package emails

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/benmanns/goworker"
)

const (
	// DefaultConcurrency - how many workers running at once
	DefaultConcurrency = 2

	// DefaultURL - connection point to redis
	DefaultURL = "redis://localhost:6379/"

	// DefaultConnections - number of connections to redis instance
	DefaultConnections = 100

	// DefaultNamespace - namespace of the workers within redis instance
	DefaultNamespace = "resque:"

	// DefaultInterval - polling between checking for enqueued jobs
	DefaultInterval = 5.0

	// DefaultSenderName - name of email client sending
	DefaultSenderName = "Five Three One"
)

// EmailClient - sendgrid client that sends email
var EmailClient = sendgrid.Client{}

// EmailPayload - payload required to send email
type EmailPayload struct {
	FirstName     string
	LastName      string
	TextBody      string
	HTMLBody      string
	RecieverEmail string
	Subject       string
}

func (ep *EmailPayload) fullName() string {
	return ep.FirstName + " " + ep.LastName
}

// InitWorkers - establishes worker settings and connection
func InitWorkers() {
	settings := goworker.WorkerSettings{
		URI:            setURI(),
		Connections:    setMaxConnections(),
		Queues:         []string{"default"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    setConcurrency(),
		Namespace:      setNamespace(),
		Interval:       DefaultInterval,
	}
	goworker.SetSettings(settings)
	goworker.Register("Default", SendEmail)
}

func setURI() string {
	url := os.Getenv("WORKERS_URL")
	if url == "" {
		return DefaultURL
	}

	return url
}

func setMaxConnections() int {
	conns := os.Getenv("MAX_WORKERS_CONNECTIONS")

	num, err := strconv.Atoi(conns)
	if err != nil || conns == "" {
		return DefaultConnections
	}

	return num
}

func setConcurrency() int {
	concurrency := os.Getenv("WORKERS_CONCURRENCY")

	num, err := strconv.Atoi(concurrency)
	if err != nil || concurrency == "" {
		return DefaultConcurrency
	}

	return num
}

func setNamespace() string {
	namespace := os.Getenv("WORKERS_NAMESPACE")
	if namespace == "" {
		return DefaultNamespace
	}

	return namespace
}

func setSenderEmail() string {
	senderEmail := os.Getenv("DEFAULT_EMAIL_ADDRESS")
	if senderEmail == "" {
		return "test@test.com"
	}

	return senderEmail
}

// SendEmail - sends email over smpt server
func SendEmail(queue string, args ...interface{}) error {
	emailPayload := EmailPayload{
		FirstName:     args[0].(string),
		LastName:      args[1].(string),
		TextBody:      args[2].(string),
		HTMLBody:      args[3].(string),
		RecieverEmail: args[4].(string),
		Subject:       args[5].(string),
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	from := mail.NewEmail(emailPayload.fullName(), setSenderEmail())
	to := mail.NewEmail(DefaultSenderName, emailPayload.RecieverEmail)
	message := mail.NewSingleEmail(from, emailPayload.Subject, to, emailPayload.TextBody, emailPayload.HTMLBody)

	res, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
	fmt.Println(res.Headers)

	return nil
}
