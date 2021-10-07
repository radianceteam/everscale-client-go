//go:build !without_local_ton_node
// +build !without_local_ton_node

package clienttest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/radianceteam/ton-client-go/client"
)

func TestClient_ProcessingWaitMessage(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()
	abi, tvc := LoadTest(EventsContract, AbiV2)
	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "keys generation")
	deployParams := &client.ParamsOfEncodeMessage{
		Abi: abi,
		DeploySet: &client.DeploySet{
			Tvc: tvc,
		},
		Signer: client.Signer{EnumTypeValue: client.KeysSigner{Keys: *keys}},

		CallSet: &client.CallSet{FunctionName: "constructor"},
	}
	msg, err := c.AbiEncodeMessage(deployParams)
	a.NoError(err, "encode message")
	err = GetGramsFromGiver(c, deployParams, msg)
	a.NoError(err, "give grams")
	events := make(chan *client.ProcessingEvent, 20)
	sendResult, err := c.ProcessingSendMessage(&client.ParamsOfSendMessage{
		Message:    msg.Message,
		SendEvents: true,
		Abi:        &abi,
	}, func(e *client.ProcessingEvent) { events <- e })
	a.NoError(err, "processing.send_message")
	a.NotEqual(0, len(sendResult.ShardBlockID))

	waitForTransactionResult, err := c.ProcessingWaitForTransaction(&client.ParamsOfWaitForTransaction{
		Message:      msg.Message,
		ShardBlockID: sendResult.ShardBlockID,
		SendEvents:   true,
		Abi:          &abi,
	}, func(e *client.ProcessingEvent) { events <- e })
	close(events)
	a.NoError(err, "processing.wait_for_transaction")
	a.Len(waitForTransactionResult.OutMessages, 0, "empty wait_for_transaction messages")

	order := []client.ProcessingEvent{
		{EnumTypeValue: client.WillFetchFirstBlockProcessingEvent{}},
		{EnumTypeValue: client.WillSendProcessingEvent{}},
		{EnumTypeValue: client.DidSendProcessingEvent{}},
	}
	for i, et := range order {
		e, ok := <-events
		a.True(ok, "event %d", i)
		a.IsType(et.EnumTypeValue, e.EnumTypeValue, "type event %d", i)
	}
	for e := range events {
		if _, ok := e.EnumTypeValue.(client.WillFetchNextBlockProcessingEvent); !ok {
			break
		}
	}
}
