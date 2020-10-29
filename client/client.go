package client

import (
	"encoding/json"

	"github.com/radianceteam/ton-client-go/spec"
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

type Config struct {
	Crypto  *CryptoConfig  `json:"crypto,omitempty"`
	ABI     *AbiConfig     `json:"abi,omitempty"`
	Network *NetworkConfig `json:"network,omitempty"`
}

func NewClient(config Config) (*Client, error) {
	rawConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	dllClient, err := newDLLClient(rawConfig)
	if err != nil {
		return nil, err
	}

	return &Client{dllClient: dllClient}, nil
}

type TransactionFees interface{}

// ClientGetAPIReference loads and parses JSON API spec.
func (c *Client) ClientGetAPIReference() (*spec.APIReference, error) {
	response := new(spec.APIReference)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.get_api_reference", nil, response)

	return response, err
}

// NetQueryCollectionRaw gives ability to unmarshall raw data yourself.
func (c *Client) NetQueryCollectionRaw(p *ParamsOfQueryCollection) ([]byte, error) {
	return c.dllClient.waitErrorOrResult("net.query_collection", p)
}

// NetSubscribeCollection Creates a subscription
//
// Triggers for each insert/update of data
// that satisfies the `filter` conditions.
// The projection fields are limited to `result` fields.
func (c *Client) NetSubscribeCollection(p *ParamsOfSubscribeCollection) (<-chan *RawResponse, *ResultOfSubscribeCollection, error) {
	responses, err := c.dllClient.resultsChannel("net.subscribe_collection", p)
	if err != nil {
		return nil, nil, err
	}

	data := <-responses
	if data.Error != nil {
		return nil, nil, data.Error
	}
	result := &ResultOfSubscribeCollection{}
	if err := json.Unmarshal(data.Data, result); err != nil {
		return nil, nil, err
	}

	return NewDynamicallyBufferedResponses(responses), result, nil
}

// ProcessingSendMessage Sends message to the network
//
// Sends message to the network and returns the last generated shard block
// (unmarshall to ResultOfSendMessage) of the destination account
// before the message was sent. It will be required later for message processing.
func (c *Client) ProcessingSendMessage(p *ParamsOfSendMessage) (<-chan *RawResponse, error) {
	responses, err := c.dllClient.resultsChannel("processing.send_message", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents {
		responses = NewDynamicallyBufferedResponses(responses)
	}

	return responses, err
}

// ProcessingWaitForTransaction Performs monitoring of the network for the result transaction
// of the external inbound message processing.
//
// `send_events` enables intermediate events, such as `WillFetchNextBlock`,
// `FetchNextBlockFailed` that may be useful for logging of new shard blocks creation
// during message processing.
//
// Note that presence of the `abi` parameter is critical for ABI
// compliant contracts. Message processing uses drastically
// different strategy for processing message for contracts which
// ABI includes "expire" header.
//
// When the ABI header `expire` is present, the processing uses
// `message expiration` strategy:
// - The maximum block gen time is set to
//   `message_expiration_time + transaction_wait_timeout`.
// - When maximum block gen time is reached the processing will
//   be finished with `MessageExpired` error.
//
// When the ABI header `expire` isn't present or `abi` parameter
// isn't specified, the processing uses `transaction waiting`
// strategy:
// - The maximum block gen time is set to
//   `now() + transaction_wait_timeout`.
//
// - If maximum block gen time is reached and no result transaction is found
// the processing will exit with an error.
func (c *Client) ProcessingWaitForTransaction(p *ParamsOfWaitForTransaction) (<-chan *RawResponse, error) {
	responses, err := c.dllClient.resultsChannel("processing.wait_for_transaction", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents {
		responses = NewDynamicallyBufferedResponses(responses)
	}

	return responses, err
}

// ProcessingProcessMessage Creates message, sends it to the network and monitors its processing.
//
// Creates ABI-compatible message,
// sends it to the network and monitors for the result transaction.
// Decodes the output message's bodies.
//
// If contract's ABI includes "expire" header then
// SDK implements retries in case of unsuccessful message delivery within the expiration
// timeout: SDK recreates the message, sends it and processes it again.
//
// The intermediate events, such as `WillFetchFirstBlock`, `WillSend`, `DidSend`,
// `WillFetchNextBlock`, etc - are switched on/off by `send_events` flag
// and logged into the supplied callback function.
// The retry configuration parameters are defined in config:
// <add correct config params here>
// pub const DEFAULT_EXPIRATION_RETRIES_LIMIT: i8 = 3; - max number of retries
// pub const DEFAULT_EXPIRATION_TIMEOUT: u32 = 40000;  - message expiration timeout in ms.
// pub const DEFAULT_....expiration_timeout_grow_factor... = 1.5 - factor that increases the expiration timeout for each retry
//
// If contract's ABI does not include "expire" header
// then if no transaction is found within the network timeout (see config parameter ), exits with error.
func (c *Client) ProcessingProcessMessage(p *ParamsOfProcessMessage) (<-chan *RawResponse, error) {
	responses, err := c.dllClient.resultsChannel("processing.process_message", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents {
		responses = NewDynamicallyBufferedResponses(responses)
	}

	return responses, err
}
