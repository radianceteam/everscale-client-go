package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/radianceteam/ton-client-go/client"
)

func main() {
	fmt.Println("before")
	c, err := client.NewClient(client.Config{
		Network: client.NetworkConfig{ServerAddress: "net.ton.dev"},
	})
	if err != nil {
		fmt.Println("err", err)

		return
	}
	defer c.Close()

	version, err := c.ClientVersion()
	fmt.Println("version", version, err)

	pair, err := c.CryptoGenerateRandomSignKeys()
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		seedBytes, _ := hex.DecodeString(pair.Secret)
		publicBytes, _ := hex.DecodeString(pair.Secret)
		private := ed25519.NewKeyFromSeed(seedBytes)
		fmt.Println("private", len(private))
		fmt.Println("public", private.Public())
		fmt.Println("public", publicBytes)
		fmt.Println("success", pair)
	}

	res, err := c.NetQueryCollection(client.NetQueryCollectionParams{
		Collection: "accounts",
		Limit:      20,
		Result:     "id balance(format:DEC)",
	})

	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("success", len(res))
	}
	for i := 0; i < 1000000; i++ {
		r, err := c.NetQueryCollection(client.NetQueryCollectionParams{
			Collection: "accounts",
			Limit:      20,
			Result:     "id balance(format:DEC)",
		})
		fmt.Println(len(r), err, i)
	}
	time.Sleep(60 * time.Second)

	fmt.Println("after")
}
