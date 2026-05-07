package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"sync"
	"time"
)

const defaultBrevoTransactionalEmailURL = "https://api.brevo.com/v3/smtp/email"

var (
	brevoOnce       sync.Once
	brevoConfigData brevoConfig
	brevoInitErr    error

	brevoHTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}
)

type brevoConfig struct {
	APIKey    string
	FromEmail string
	FromName  string
	APIURL    string
}

type brevoEmailRequest struct {
	Sender      brevoEmailSender      `json:"sender"`
	To          []brevoEmailRecipient `json:"to"`
	Subject     string                `json:"subject"`
	TextContent string                `json:"textContent"`
}

type brevoEmailSender struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type brevoEmailRecipient struct {
	Email string `json:"email"`
}

// Mantém compatibilidade com qualquer trecho antigo que ainda chame SendEmail sem context.
func SendEmail(to string, subject string, body string) error {
	return SendEmailWithContext(context.Background(), to, subject, body)
}

func SendEmailWithContext(ctx context.Context, to string, subject string, body string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	brevoOnce.Do(func() {
		brevoConfigData, brevoInitErr = loadBrevoConfig()
	})

	if brevoInitErr != nil {
		return brevoInitErr
	}

	payload, err := buildBrevoEmailPayload(brevoConfigData, to, subject, body)
	if err != nil {
		return err
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao montar payload do email Brevo: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		brevoConfigData.APIURL,
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		return fmt.Errorf("erro ao criar request para Brevo: %w", err)
	}

	req.Header.Set("api-key", brevoConfigData.APIKey)
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")

	resp, err := brevoHTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao enviar email via Brevo: %w", err)
	}
	defer resp.Body.Close()

	responseBody := readLimitedResponseBody(resp.Body)

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		if responseBody == "" {
			responseBody = http.StatusText(resp.StatusCode)
		}

		return fmt.Errorf(
			"erro ao enviar email via Brevo: status %d: %s",
			resp.StatusCode,
			responseBody,
		)
	}

	return nil
}

func loadBrevoConfig() (brevoConfig, error) {
	cfg := brevoConfig{
		APIKey:    strings.TrimSpace(os.Getenv("BREVO_API_KEY")),
		FromEmail: strings.TrimSpace(os.Getenv("BREVO_FROM_EMAIL")),
		FromName:  strings.TrimSpace(os.Getenv("BREVO_FROM_NAME")),
		APIURL:    strings.TrimSpace(os.Getenv("BREVO_API_URL")),
	}

	if cfg.APIKey == "" {
		return cfg, errors.New("variável BREVO_API_KEY não configurada")
	}

	if cfg.FromEmail == "" {
		return cfg, errors.New("variável BREVO_FROM_EMAIL não configurada")
	}

	parsedFrom, err := mail.ParseAddress(cfg.FromEmail)
	if err != nil {
		return cfg, fmt.Errorf("BREVO_FROM_EMAIL inválido: %w", err)
	}

	cfg.FromEmail = parsedFrom.Address

	if cfg.APIURL == "" {
		cfg.APIURL = defaultBrevoTransactionalEmailURL
	}

	return cfg, nil
}

func buildBrevoEmailPayload(cfg brevoConfig, to string, subject string, body string) (brevoEmailRequest, error) {
	to = strings.TrimSpace(to)
	subject = sanitizeEmailHeader(subject)
	body = strings.TrimSpace(body)

	if to == "" {
		return brevoEmailRequest{}, errors.New("email de destino vazio")
	}

	parsedTo, err := mail.ParseAddress(to)
	if err != nil {
		return brevoEmailRequest{}, fmt.Errorf("email de destino inválido: %w", err)
	}

	if subject == "" {
		return brevoEmailRequest{}, errors.New("assunto do email vazio")
	}

	if body == "" {
		return brevoEmailRequest{}, errors.New("corpo do email vazio")
	}

	return brevoEmailRequest{
		Sender: brevoEmailSender{
			Name:  cfg.FromName,
			Email: cfg.FromEmail,
		},
		To: []brevoEmailRecipient{
			{
				Email: parsedTo.Address,
			},
		},
		Subject:     subject,
		TextContent: body,
	}, nil
}

func sanitizeEmailHeader(value string) string {
	value = strings.ReplaceAll(value, "\r", " ")
	value = strings.ReplaceAll(value, "\n", " ")

	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func readLimitedResponseBody(body io.Reader) string {
	data, err := io.ReadAll(io.LimitReader(body, 4096))
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}