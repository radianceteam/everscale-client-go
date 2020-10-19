//nolint
package main

import (
	"encoding/json"

	"github.com/radianceteam/ton-client-go/client"
	"github.com/radianceteam/ton-client-go/spec"
)

func main() {
	c, err := client.NewClient(client.Config{})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	ref, err := c.ClientGetAPIReference()
	if err != nil {
		panic(err)
	}
	var api spec.APIReference
	err = json.Unmarshal(ref, &api)
	if err != nil {
		panic(err)
	}
	for _, m := range api.API.Modules {
		err = spec.GenModule("./tmp", m)
		if err != nil {
			panic(err)
		}
	}
}
