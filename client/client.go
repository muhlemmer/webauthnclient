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
	pc := new(portableClient)
	if err := pc.load(filename); err != nil {
		return nil, err
	}
	return &Client{
		RelyingParty:  pc.RelyingParty,
		Authenticator: pc.Authenticator,
		Credential:    pc.Credential.ToCredential(),
	}, nil
}

func (c *Client) Store(filename string) error {
	pc := &portableClient{
		RelyingParty:  c.RelyingParty,
		Authenticator: c.Authenticator,
		Credential:    c.Credential.ExportToPortableCredential(),
	}
	return pc.store(filename)
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

type portableClient struct {
	RelyingParty  virtualwebauthn.RelyingParty       `json:"relying_party,omitempty"`
	Authenticator virtualwebauthn.Authenticator      `json:"authenticator,omitempty"`
	Credential    virtualwebauthn.PortableCredential `json:"credential,omitempty"`
}

func (c *portableClient) load(filename string) error {
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

func (c *portableClient) store(filename string) error {
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
