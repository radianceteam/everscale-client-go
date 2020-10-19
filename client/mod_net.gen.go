package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.683584 +0000 UTC
// Mod net
//  Network access.
//  Network access.

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
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Order      *[]OrderBy  `json:"order,omitempty"`
	Limit      *int        `json:"limit,omitempty"`
}

type ResultOfQueryCollection struct {
	Result []interface{} `json:"result"`
}

type ParamsOfWaitForCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
	Timeout    *int        `json:"timeout,omitempty"`
}

type ResultOfWaitForCollection struct {
	Result interface{} `json:"result"`
}

type ResultOfSubscribeCollection struct {
	Handle int `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	Collection string      `json:"collection"`
	Filter     interface{} `json:"filter,omitempty"`
	Result     string      `json:"result"`
}
