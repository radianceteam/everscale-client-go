// +build !without_local_ton_node

package clienttest

import (
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/radianceteam/ton-client-go/client"
)

func TestClient_NetSuspendResume(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()
	a.NoError(c.NetSuspend(), "suspend")
	a.NoError(c.NetResume(), "resume")
}

func TestClient_NetSubscribeCollectionEmpty(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()
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
			a.NotNil(r, "no data when unsubscribe")
		}
	}()
	err = c.NetUnsubscribe(&client.ResultOfSubscribeCollection{Handle: handle.Handle})
	a.NoError(err, "unsubscribe")
	wg.Wait()
}

func TestClient_NetSubscribeForTransactionWithAddressees(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()
	abi, tvc := LoadTest(HelloContract, AbiV2)
	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "keys generation")

	deployParams := &client.ParamsOfEncodeMessage{
		Abi: abi,
		DeploySet: &client.DeploySet{
			Tvc: tvc,
		},
		Signer: client.Signer{Type: client.KeysSignerType, Keys: *keys},

		CallSet: &client.CallSet{FunctionName: "constructor"},
	}
	msg, err := c.AbiEncodeMessage(deployParams)
	a.NoError(err, "encode message")

	responses, handle, err := c.NetSubscribeCollection(&client.ParamsOfSubscribeCollection{
		Collection: "transactions",
		Filter: json.RawMessage(`{
			"account_addr": { "eq": "` + msg.Address + `" },
			"status": { "eq": 3 }
		}`),
		Result: "id account_addr",
	})
	a.NoError(err, "subscribe_collection")
	a.NotNil(handle, "subscription handle")
	a.NotNil(responses, "responses channel")
	err = GetGramsFromGiver(c, deployParams, msg)
	a.NoError(err, "get grams from giver")

	wg := sync.WaitGroup{}
	wg.Add(1)
	var transactions []interface{}
	go func() {
		defer wg.Done()
		for r := range responses {
			a.NotNil(r, "no data when unsubscribe")
			transactions = append(transactions, r)
		}
	}()
	time.Sleep(1 * time.Second) // wait for messages processing
	err = c.NetUnsubscribe(&client.ResultOfSubscribeCollection{Handle: handle.Handle})
	a.NoError(err, "unsubscribe")
	wg.Wait()
	a.Equal(2, len(transactions), "received transactions count")
	a.NotNil(transactions[0].(map[string]interface{})["id"], "first transaction id")
	a.NotNil(transactions[1].(map[string]interface{})["id"], "second transaction id")
}
