package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModClient(t *testing.T) {
	a := assert.New(t)
	c, err := NewClient(Config{
		Network: NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	}
	fmt.Println("Client", c, err)
	a.NoError(err, "Client created")
	version, err := c.ClientVersion()
	a.NoError(err, "call Client.version")
	if !a.NotNil(version, "version response") {
		return
	}
	a.Equal("1.0.0", version.Version, "dll with specified version")
	ref, err := c.ClientGetAPIReference()
	a.NoError(err, "call Client.get_api_version")
	a.NotNil(ref, "ref not nil")
}
