package helper

import (
	"bytes"
	"context"
	"f-commerce/config"
	"fmt"
	"text/template"
	"time"

	"github.com/mailersend/mailersend-go"
	"golang.org/x/exp/rand"
)

func GenerateOTP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(uint64(seed)))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000)) // Generate 6 digit OTP
	return otp
}

type EmailOTP struct {
	OTP string
}

func SendOTPEmail(to []mailersend.Recipient, otp string) error {

	cfg, err := config.SetConfig()
	if err != nil {
		return err
	}

	ms := mailersend.NewMailersend(cfg.SmtpEmail.ApiKey)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tmpl, err := template.ParseFiles("view/otp_template.html")
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	from := mailersend.From{
		Name:  cfg.SmtpEmail.SmtpUser,
		Email: cfg.SmtpEmail.FromEmail,
	}

	data := EmailOTP{OTP: otp}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(to)
	message.SetHTML(body.String())
	message.SetSubject("OTP Verification")
	message.SetTags(tags)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	messageID := res.Header.Get("X-Message-Id")
	if messageID == "" {
		return fmt.Errorf("failed to retrieve message ID from response")
	}

	return nil
}
