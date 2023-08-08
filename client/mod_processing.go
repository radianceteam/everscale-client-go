package client

// Manually implemented methods for mod_processing.
// Throws panic only when there is critical error.

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrNoCallbackSpecified = errors.New("no callback specified error when send_events enabled")

type EventCallback func(event *ProcessingEvent)

func handleEvents(responses <-chan *RawResponse, callback EventCallback, result interface{}) error {
	for r := range responses {
		switch r.Code { // nolint exhaustive
		case ResponseCodeCustom:
			event := &ProcessingEvent{}
			if err := json.Unmarshal(r.Data, event); err != nil {
				panic(err)
			}
			callback(event)
		case ResponseCodeError:
			return r.Error
		case ResponseCodeSuccess:
			if err := json.Unmarshal(r.Data, result); err != nil {
				panic(err)
			}

			return nil
		default:
			panic(fmt.Errorf("unknown response type code %v", r.Code))
		}
	}

	return nil
}

// ProcessingSendMessage Sends message to the network
//
// Sends message to the network and returns the last generated shard block of the destination account
// before the message was sent. It will be required later for message processing.
func (c *Client) ProcessingSendMessage(p *ParamsOfSendMessage, callback EventCallback) (*ResultOfSendMessage, error) {
	if p.SendEvents.Bool && callback == nil {
		return nil, ErrNoCallbackSpecified
	}
	responses, err := c.dllClient.resultsChannel("processing.send_message", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents.Bool {
		responses = NewDynamicallyBufferedResponses(responses)
	}
	result := &ResultOfSendMessage{}

	return result, handleEvents(responses, callback, result)
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
//   - The maximum block gen time is set to
//     `message_expiration_time + transaction_wait_timeout`.
//   - When maximum block gen time is reached the processing will
//     be finished with `MessageExpired` error.
//
// When the ABI header `expire` isn't present or `abi` parameter
// isn't specified, the processing uses `transaction waiting`
// strategy:
//   - The maximum block gen time is set to
//     `now() + transaction_wait_timeout`.
//
// - If maximum block gen time is reached and no result transaction is found
// the processing will exit with an error.
func (c *Client) ProcessingWaitForTransaction(p *ParamsOfWaitForTransaction, callback EventCallback) (*ResultOfProcessMessage, error) {
	if p.SendEvents.Bool && callback == nil {
		return nil, ErrNoCallbackSpecified
	}
	responses, err := c.dllClient.resultsChannel("processing.wait_for_transaction", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents.Bool {
		responses = NewDynamicallyBufferedResponses(responses)
	}

	result := &ResultOfProcessMessage{}

	return result, handleEvents(responses, callback, result)
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
func (c *Client) ProcessingProcessMessage(p *ParamsOfProcessMessage, callback EventCallback) (*ResultOfProcessMessage, error) {
	if p.SendEvents.Bool && callback == nil {
		return nil, ErrNoCallbackSpecified
	}
	responses, err := c.dllClient.resultsChannel("processing.process_message", p)
	if err != nil {
		return nil, err
	}
	if p.SendEvents.Bool {
		responses = NewDynamicallyBufferedResponses(responses)
	}

	result := &ResultOfProcessMessage{}

	return result, handleEvents(responses, callback, result)
}
