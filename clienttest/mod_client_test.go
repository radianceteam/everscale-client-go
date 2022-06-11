package clienttest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"

	"github.com/radianceteam/everscale-client-go/client"
)

func TestModClient(t *testing.T) {
	a := assert.New(t)
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: null.NewString("net.ton.dev", true)},
	}, client.WrapperConfig{MaxCGOConcurrentThreads: 10})
	if err == nil {
		defer c.Close()
	}
	a.NoError(err, "Client created")
	version, err := c.ClientVersion()
	a.NoError(err, "call Client.version")
	if !a.NotNil(version, "version response") {
		return
	}
	a.True(strings.HasPrefix(version.Version, "1."), "dll with major version 1")
	ref, err := c.ClientGetAPIReference()
	a.NoError(err, "call Client.get_api_version")
	a.NotNil(ref, "ref not nil")
}
