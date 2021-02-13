package client

// Manually implemented methods for mod_net.

import (
	"encoding/json"
	"fmt"
)

type ParamsOfQueryOperation struct {
	Value interface{} // any of ParamsOfQueryCollection, ParamsOfWaitForCollection or ParamsOfAggregateCollection
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p ParamsOfQueryOperation) MarshalJSON() ([]byte, error) {
	switch value := (p.Value).(type) {
	case ParamsOfQueryCollection:
		return json.Marshal(struct {
			ParamsOfQueryCollection
			Type string `json:"type"`
		}{
			value,
			"QueryCollection",
		})
	case ParamsOfWaitForCollection:
		return json.Marshal(struct {
			ParamsOfWaitForCollection
			Type string `json:"type"`
		}{
			value,
			"WaitForCollection",
		})
	case ParamsOfAggregateCollection:
		return json.Marshal(struct {
			ParamsOfAggregateCollection
			Type string `json:"type"`
		}{
			value,
			"AggregateCollection",
		})
	default:
		panic(fmt.Sprintf("unsupported type for ParamsOfQueryOperation %v", p.Value))
	}
}

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
