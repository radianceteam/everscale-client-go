package client

// DON'T EDIT THIS FILE is generated 20 Oct 20 13:40 UTC
//
// Mod client
//
// Provides information about library.

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type NetworkConfig struct {
	ServerAddress            string              `json:"server_address"`
	MessageRetriesCount      null.Int            `json:"message_retries_count"`      // optional
	MessageProcessingTimeout null.Int            `json:"message_processing_timeout"` // optional
	WaitForTimeout           null.Int            `json:"wait_for_timeout"`           // optional
	OutOfSyncThreshold       decimal.NullDecimal `json:"out_of_sync_threshold"`      // optional
	AccessKey                null.String         `json:"access_key"`                 // optional
}

type CryptoConfig struct {
	FishParam null.String `json:"fish_param"` // optional
}

type AbiConfig struct {
	MessageExpirationTimeout           null.Int `json:"message_expiration_timeout"`             // optional
	MessageExpirationTimeoutGrowFactor null.Int `json:"message_expiration_timeout_grow_factor"` // optional
}

type ResultOfVersion struct {
	// core version.
	Version string `json:"version"`
}

func (c *Client) ClientVersion() (*ResultOfVersion, error) {
	response := new(ResultOfVersion)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.version", nil, response)

	return response, err
}
