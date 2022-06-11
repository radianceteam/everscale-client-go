package client

import (
	"encoding/json"

	"github.com/radianceteam/everscale-client-go/spec"
)

type Client struct {
	dllClient
}

func (c *Client) Close() {
	if c == nil {
		return
	}
	c.dllClient.close()
}

type WrapperConfig struct {
	// Prevents from spawning unlimited system threads
	// when EVER-SDK external code is called via CGO
	// for more details please see `func cgocall(fn, arg unsafe.Pointer) int32`
	// implementation and comments in `runtime` package.
	MaxCGOConcurrentThreads uint
}

func NewClient(config Config, wrapperConfig WrapperConfig) (*Client, error) {
	rawConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	MaxCGOConcurrentThreads := wrapperConfig.MaxCGOConcurrentThreads
	if MaxCGOConcurrentThreads == 0 {
		MaxCGOConcurrentThreads = 1 // to prevent dead-lock before CGO call.
	}
	dllClient, err := newDLLClient(rawConfig, MaxCGOConcurrentThreads)
	if err != nil {
		return nil, err
	}

	return &Client{dllClient: dllClient}, nil
}

// ClientGetAPIReference loads and parses JSON API spec.
func (c *Client) ClientGetAPIReference() (*spec.APIReference, error) {
	response := new(spec.APIReference)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.get_api_reference", nil, response)

	return response, err
}
