package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var (
	sesClient           *sesv2.Client
	sesFromEmail        string
	sesConfigurationSet string

	sesOnce    sync.Once
	sesInitErr error
)

// Mantém compatibilidade caso algum trecho antigo ainda chame SendEmail sem context.
// func SendEmail(to string, subject string, body string) error {
// 	return SendEmailWithContext(context.Background(), to, subject, body)
// }

func SendEmailWithContext(ctx context.Context, to string, subject string, body string) error {
	sesOnce.Do(func() {
		sesInitErr = initSES(ctx)
	})

	if sesInitErr != nil {
		return sesInitErr
	}

	to = strings.TrimSpace(to)
	subject = strings.TrimSpace(subject)

	if to == "" {
		return errors.New("email de destino vazio")
	}

	if subject == "" {
		return errors.New("assunto do email vazio")
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(sesFromEmail),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data:    aws.String(subject),
					Charset: aws.String("UTF-8"),
				},
				Body: &types.Body{
					Text: &types.Content{
						Data:    aws.String(body),
						Charset: aws.String("UTF-8"),
					},
				},
			},
		},
	}

	if sesConfigurationSet != "" {
		input.ConfigurationSetName = aws.String(sesConfigurationSet)
	}

	_, err := sesClient.SendEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("erro ao enviar email via SES: %w", err)
	}

	return nil
}

func initSES(ctx context.Context) error {
	from := strings.TrimSpace(os.Getenv("SES_FROM_EMAIL"))
	if from == "" {
		return errors.New("variável SES_FROM_EMAIL não configurada")
	}

	region := strings.TrimSpace(os.Getenv("SES_REGION"))

	var (
		cfg aws.Config
		err error
	)

	if region != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}

	if err != nil {
		return fmt.Errorf("erro ao carregar configuração AWS: %w", err)
	}

	sesClient = sesv2.NewFromConfig(cfg)
	sesFromEmail = from
	sesConfigurationSet = strings.TrimSpace(os.Getenv("SES_CONFIGURATION_SET"))

	return nil
}