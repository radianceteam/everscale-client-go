package clienttest

import (
	"sync"
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

func TestClient_NetSubscribeCollection(t *testing.T) {
	a := assert.New(t)
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	} else {
		return
	}
	responses, handle, err := c.NetSubscribeCollection(&client.ParamsOfSubscribeCollection{
		Collection: "messages",
	})
	a.NoError(err, "subscribe_collection")
	a.NotNil(handle)
	a.NotNil(responses)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for r := range responses {
			a.NoError(r.Error, "no error when unsubscribe")
			a.Nil(r.Data, "no data when unsubscribe")
		}
	}()
	err = c.NetUnsubscribe(&client.ResultOfSubscribeCollection{Handle: handle.Handle})
	a.NoError(err, "unsubscribe")
	wg.Wait()
}
