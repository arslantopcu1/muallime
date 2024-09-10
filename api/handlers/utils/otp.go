package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"os"
)

func init() {
	log.Info("SendOtp init ")
}

func SendOtp(phone string) (string, error) {

	to := phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("whatsapp")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNTID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	resp, err := client.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("Otp failed to generate")
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		return *resp.Sid, nil
	}
}

func CheckOtp(phone, code string) error {
	to := phone
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNTID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Invalid otp")
	} else if *resp.Status == "approved" {
		return nil
	} else {
		return errors.New("Invalid otp")
	}
}
