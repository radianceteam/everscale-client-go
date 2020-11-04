package clienttest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/radianceteam/ton-client-go/client"
)

func init() { // nolint gochecknoinits
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("02 Jan 15:04:05")
	l, _ := config.Build()
	zap.ReplaceGlobals(l)
}

func TestModClient(t *testing.T) {
	a := assert.New(t)
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	}
	a.NoError(err, "Client created")
	version, err := c.ClientVersion()
	a.NoError(err, "call Client.version")
	if !a.NotNil(version, "version response") {
		return
	}
	a.Equal("1.1.0", version.Version, "dll with specified version")
	ref, err := c.ClientGetAPIReference()
	a.NoError(err, "call Client.get_api_version")
	a.NotNil(ref, "ref not nil")
}
