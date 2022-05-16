package client

// Manually implemented methods for mod_net.

import (
	"encoding/json"
)

// NetSubscribeCollection Creates a subscription with unmarshalled to interface{} data.
//
// Triggers for each insert/update of data
// that satisfies the `filter` conditions.
// The projection fields are limited to `result` fields.
func (c *Client) NetSubscribeCollection(p *ParamsOfSubscribeCollection) (<-chan json.RawMessage, *ResultOfSubscribeCollection, error) {
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

	responses = NewDynamicallyBufferedResponses(responses)

	unmarshalled := make(chan json.RawMessage, 1)
	go func() {
		var body struct {
			Result json.RawMessage `json:"result"`
		}
		for r := range responses {
			if err := json.Unmarshal(r.Data, &body); err != nil {
				panic(err)
			}
			unmarshalled <- body.Result
		}
		close(unmarshalled)
	}()

	return unmarshalled, result, nil
}

// NetSubscribe Creates a subscription.
//
// The subscription is a persistent communication channel between
// client and Everscale Network.
//
// ### Important Notes on Subscriptions
//
// Unfortunately sometimes the connection with the network brakes down.
// In this situation the library attempts to reconnect to the network.
// This reconnection sequence can take significant time.
// All of this time the client is disconnected from the network.
// Bad news is that all changes that happened while
// the client was disconnected are lost.
//
// Good news is that the client report errors to the callback when
// it loses and resumes connection.
//
// So, if the lost changes are important to the application then
// the application must handle these error reports.
//
// Library reports errors with `responseType` == 101
// and the error object passed via `params`.
// When the library has successfully reconnected
// the application receives callback with
// `responseType` == 101 and `params.code` == 614 (NetworkModuleResumed).
// Application can use several ways to handle this situation:
// - If application monitors changes for the single
// object (for example specific account):  application
// can perform a query for this object and handle actual data as a
// regular data from the subscription.
// - If application monitors sequence of some objects
// (for example transactions of the specific account): application must
// refresh all cached (or visible to user) lists where this sequences presents.
func (c *Client) NetSubscribe(p *ParamsOfSubscribe) (<-chan *RawResponse, *ResultOfSubscribeCollection, error) {
	responses, err := c.dllClient.resultsChannel("net.subscribe", p)
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
