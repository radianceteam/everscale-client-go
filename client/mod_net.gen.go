package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:44:28.807717 +0000 UTC
// Mod net
// Network access.
// Network access.

type OrderBy struct {
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type ParamsOfQueryCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures)
	Collection string `json:"collection"`
	// collection filter
	Filter interface{} `json:"filter,omitempty"`
	// projection (result) string
	Result string `json:"result"`
	// sorting order
	Order *[]OrderBy `json:"order,omitempty"`
	// number of documents to return
	Limit *int `json:"limit,omitempty"`
}

type ResultOfQueryCollection struct {
	// objects that match provided criteria
	Result []interface{} `json:"result"`
}

type ParamsOfWaitForCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures)
	Collection string `json:"collection"`
	// collection filter
	Filter interface{} `json:"filter,omitempty"`
	// projection (result) string
	Result string `json:"result"`
	// query timeout
	Timeout *int `json:"timeout,omitempty"`
}

type ResultOfWaitForCollection struct {
	// first found object that match provided criteria
	Result interface{} `json:"result"`
}

type ResultOfSubscribeCollection struct {
	// handle to subscription. It then can be used in `get_next_subscription_data` function
	// and must be closed with `unsubscribe`
	Handle int `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures)
	Collection string `json:"collection"`
	// collection filter
	Filter interface{} `json:"filter,omitempty"`
	// projection (result) string
	Result string `json:"result"`
}
