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
		EnumTypeValue: client.ParamsOfQueryCollection{
			Collection: "collection_name",
		},
	}
	jsonBytes, err := json.MarshalIndent(&params, "", "  ")
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

func TestClient_Error(t *testing.T) {
	a := assert.New(t)
	a.EqualError(&client.Error{Code: client.AppRequestErrorErrorCode, Message: "some app error", Data: json.RawMessage(`{"val": "field"}`)},
		`sdk_error_code=26,
 sdk_error_code_description=AppRequestErrorErrorCode,
 sdk_error_msg='some app error',
 sdk_error_data: '{"val": "field"}'`)
}
