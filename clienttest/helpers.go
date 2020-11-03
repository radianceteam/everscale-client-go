package clienttest

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/radianceteam/ton-client-go/client"
)

const (
	EventsContract       = "Events"
	HelloContract        = "Hello"
	LimitWalletContract  = "LimitWallet"
	PiggyContract        = "Piggy"
	SubscriptionContract = "Subscription"
	WalletContract       = "Wallet"
)

func LoadTest(name string) (client.Abi, string) {
	content, err := ioutil.ReadFile("./contracts/abi_v2/" + name + ".abi.json")
	if err != nil {
		panic(err)
	}
	abi := client.Abi{Type: client.ContractAbiType}
	if err = json.Unmarshal(content, &abi.Value); err != nil {
		panic(err)
	}
	content, err = ioutil.ReadFile("./contracts/abi_v2/" + name + ".tvc")
	if err != nil {
		panic(err)
	}

	return abi, base64.StdEncoding.EncodeToString(content)
}

func signDetached(c *client.Client, data string, keys client.KeyPair) string {
	signKeys, _ := c.CryptoNaclSignKeypairFromSecretKey(&client.ParamsOfNaclSignKeyPairFromSecret{
		Secret: keys.Secret,
	})

	result, _ := c.CryptoNaclSignDetached(&client.ParamsOfNaclSign{Unsigned: data, Secret: signKeys.Secret})

	return result.Signature
}
