package client

// DON'T EDIT THIS FILE is generated 2020-10-19 11:18:28.293188 +0000 UTC
//
// Mod client
//
// Provides information about library.

import (
	"github.com/shopspring/decimal"
)

type NetworkConfig struct {
	ServerAddress            string           `json:"server_address"`
	MessageRetriesCount      *int             `json:"message_retries_count,omitempty"`
	MessageProcessingTimeout *int             `json:"message_processing_timeout,omitempty"`
	WaitForTimeout           *int             `json:"wait_for_timeout,omitempty"`
	OutOfSyncThreshold       *decimal.Decimal `json:"out_of_sync_threshold,omitempty"`
	AccessKey                *string          `json:"access_key,omitempty"`
}

type CryptoConfig struct {
	FishParam *string `json:"fish_param,omitempty"`
}

type AbiConfig struct {
	MessageExpirationTimeout           *int `json:"message_expiration_timeout,omitempty"`
	MessageExpirationTimeoutGrowFactor *int `json:"message_expiration_timeout_grow_factor,omitempty"`
}

type ResultOfVersion struct {
	// core version.
	Version string `json:"version"`
}

func (c *Client) Version() {}
