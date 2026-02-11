package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client maneja la comunicación con el servicio de autenticación
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient crea una nueva instancia del cliente de autenticación
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// ValidateTokenResponse representa la respuesta de validación del token
type ValidateTokenResponse struct {
	Valid     bool   `json:"valid"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Role      int    `json:"role"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiresAt int64  `json:"expires_at"`
}

// ValidateToken valida un token JWT con el servicio de autenticación
func (c *Client) ValidateToken(token string) (*ValidateTokenResponse, error) {
	url := fmt.Sprintf("%s/auth/validate", c.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token validation failed: %s", string(body))
	}

	var result ValidateTokenResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
