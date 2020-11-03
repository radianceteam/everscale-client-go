package clienttest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/radianceteam/ton-client-go/client"
)

var keys = client.KeyPair{
	Public: "4c7c408ff1ddebb8d6405ee979c716a14fdd6cc08124107a61d3c25597099499",
	Secret: "cc8929d635719612a9478b9cd17675a39cfad52d8959e8a177389b8c0b9122a7",
}

func TestClient_AbiEncodeMessage(t *testing.T) {
	a := assert.New(t)
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err == nil {
		defer c.Close()
	} else {
		panic(err)
	}
	a.Nil(nil)
	//msg := &ParamsOfEncodeMessage{
	//	Abi: Abi{Type: HandleAbiType, Content: 0},
	//}
}
