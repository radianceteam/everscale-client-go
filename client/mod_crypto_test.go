package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModCrypto(t *testing.T) {
	a := assert.New(t)
	c, err := NewClient(Config{
		Network: NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	}
	a.NoError(err, "Client created")
	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "call Client.version")
	a.Len(keys.Public, 64, "hex len")
	a.Len(keys.Secret, 64, "hex len")
}
