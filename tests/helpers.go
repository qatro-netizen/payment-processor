package payment_processor

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type paymentService interface {
	VerifyCard(ctx context.Context, cardNumber string, cvc string, expiryDate string) (*response, error)
}

type response struct {
	Success bool `json:"success"`
}

type paymentGateway struct{}

func NewPaymentGateway() *paymentGateway {
	return &paymentGateway{}
}

func (pg *paymentGateway) VerifyCard(ctx context.Context, cardNumber string, cvc string, expiryDate string) (*response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://payment-gateway.com/verify-card", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(strings.NewReader(`{"card_number": "`+cardNumber+`","cvc": "`+cvc+`","expiry_date": "`+expiryDate+`"}`))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type paymentMethod struct {
	Card   *card `json:"card"`
	Cash   bool  `json:"cash"`
	Bank   bool  `json:"bank"`
	Other  bool  `json:"other"`
	Amount float64 `json:"amount"`
}

type card struct {
	Number     string `json:"number"`
	Cvc        string `json:"cvc"`
	ExpiryDate string `json:"expiry_date"`
}

type payment struct {
	Amount    float64 `json:"amount"`
	PaymentMethod   `json:"payment_method"`
}

type paymentRequest struct {
	Payment     payment     `json:"payment"`
	TransactionID string `json:"transaction_id"`
}

func (pg *paymentGateway) ProcessPayment(ctx context.Context, req *paymentRequest) (*response, error) {
	return pg.VerifyCard(ctx, req.Payment.PaymentMethod.Card.Number, req.Payment.PaymentMethod.Card.Cvc, req.Payment.PaymentMethod.Card.ExpiryDate)
}