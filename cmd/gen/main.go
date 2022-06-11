//nolint
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/radianceteam/everscale-client-go/spec"
)

func main() {
	file, err := os.Open("./api-spec.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ref, err := ioutil.ReadAll(file)
	var api spec.APIReference
	// fmt.Println("ref", ref)
	err = json.Unmarshal(ref, &api)
	if err != nil {
		panic(err)
	}
	for _, m := range api.Modules {
		err = spec.GenModule("./client", m)
		if err != nil {
			panic(err)
		}
	}
}
