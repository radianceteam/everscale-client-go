package clienttest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/radianceteam/ton-client-go/client"
)

func TestClient_ParamsOfQueryOperation_MarshalJSON(t *testing.T) {
	a := assert.New(t)
	params := client.ParamsOfQueryOperation{
		Value: client.ParamsOfQueryCollection{
			Collection: "collection_name",
		},
	}
	jsonBytes, err := json.MarshalIndent(params, "", "  ")
	fmt.Println(string(jsonBytes))
	a.NoError(err, "marshal")
	a.JSONEq(`{
  "collection": "collection_name",
  "filter": null,
  "result": "",
  "order": null,
  "limit": null,
  "type": "QueryCollection"
}`, string(jsonBytes), "add type field")
}
