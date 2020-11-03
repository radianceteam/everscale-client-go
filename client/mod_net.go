package client

// Manually implemented methods for mod_net.

import "encoding/json"

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
