package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/descope/virtualwebauthn"
)

type Client struct {
	RelyingParty  virtualwebauthn.RelyingParty  `json:"relying_party,omitempty"`
	Authenticator virtualwebauthn.Authenticator `json:"authenticator,omitempty"`
	Credential    virtualwebauthn.Credential    `json:"credential,omitempty"`
}

func NewClient(name, domain, origin string) *Client {
	rp := virtualwebauthn.RelyingParty{
		Name:   name,
		ID:     domain,
		Origin: origin,
	}
	return &Client{
		RelyingParty:  rp,
		Authenticator: virtualwebauthn.NewAuthenticator(),
		Credential:    virtualwebauthn.NewCredential(virtualwebauthn.KeyTypeEC2),
	}
}

func LoadClient(filename string) (*Client, error) {
	c := new(Client)
	if err := c.Load(filename); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("client load: %w", err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err = dec.Decode(c); err != nil {
		return fmt.Errorf("client load: %w", err)
	}
	return nil
}

func (c *Client) Store(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("client store: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err = enc.Encode(c); err != nil {
		return fmt.Errorf("client store: %w", err)
	}
	return nil
}

func (c *Client) CreateAttestationResponse(options string) (string, error) {
	parsedAttestationOptions, err := virtualwebauthn.ParseAttestationOptions(options)
	if err != nil {
		return "", fmt.Errorf("client CreateAttestationResponse: %w", err)
	}
	return virtualwebauthn.CreateAttestationResponse(
		c.RelyingParty, c.Authenticator, c.Credential, *parsedAttestationOptions,
	), nil
}

func (c *Client) CreateAssertionResponse(options string) (string, error) {
	parsedAssertionOptions, err := virtualwebauthn.ParseAssertionOptions(options)
	if err != nil {
		return "", fmt.Errorf("client CreateAssertionResponse: %w", err)
	}
	return virtualwebauthn.CreateAssertionResponse(
		c.RelyingParty, c.Authenticator, c.Credential, *parsedAssertionOptions,
	), nil
}

/*
func main() {
	client := NewClient("Demo", "localhost", "http://localhost:8080")
	resp, err := client.CreateAttestationResponse([]byte(os.Args[1]))
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", resp)
}
*/
