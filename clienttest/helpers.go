package clienttest

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/guregu/null.v4"

	"github.com/radianceteam/ton-client-go/client"
)

func init() { // nolint gochecknoinits
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("02 Jan 15:04:05")
	l, _ := config.Build()
	zap.ReplaceGlobals(l)
}

const (
	EventsContract       = "Events"
	HelloContract        = "Hello"
	LimitWalletContract  = "LimitWallet"
	GiverContract        = "Giver"
	GiverWalletContract  = "GiverWallet"
	PiggyContract        = "Piggy"
	SubscriptionContract = "Subscription"
	WalletContract       = "Wallet"
)

type AbiVersion string

const (
	AbiV1 AbiVersion = "abi_v1"
	AbiV2 AbiVersion = "abi_v2"
)

const GiverAddress = "0:841288ed3b55d9cdafa806807f02a0ae0c169aa5edfe88a789a6482429756a94"

func LoadTest(name string, v AbiVersion) (client.Abi, string) {
	return LoadAbi(name, v), LoadTvc(name, v)
}

func LoadAbi(name string, v AbiVersion) client.Abi {
	content, err := ioutil.ReadFile("./contracts/" + string(v) + "/" + name + ".abi.json")
	if err != nil {
		panic(err)
	}
	abi := client.Abi{Type: client.ContractAbiType}
	if err = json.Unmarshal(content, &abi.Value); err != nil {
		panic(err)
	}

	return abi
}

func LoadTvc(name string, v AbiVersion) string {
	content, err := ioutil.ReadFile("./contracts/" + string(v) + "/" + name + ".tvc")
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(content)
}

func NewTestClient() *client.Client {
	c, err := client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: "http://localhost"},
	})
	if err != nil {
		panic(err)
	}

	return c
}

func signDetached(c *client.Client, data string, keys client.KeyPair) string {
	signKeys, _ := c.CryptoNaclSignKeypairFromSecretKey(&client.ParamsOfNaclSignKeyPairFromSecret{
		Secret: keys.Secret,
	})

	result, _ := c.CryptoNaclSignDetached(&client.ParamsOfNaclSign{Unsigned: data, Secret: signKeys.Secret})

	return result.Signature
}

func GetGramsFromGiver(c *client.Client, msgParams *client.ParamsOfEncodeMessage, msg *client.ResultOfEncodeMessage) error {
	abi := LoadAbi(GiverContract, AbiV1)
	params := client.ParamsOfEncodeMessage{
		Address: null.StringFrom(GiverAddress),
		Abi:     abi,
		CallSet: &client.CallSet{
			FunctionName: "sendGrams",
			Input: map[string]interface{}{
				"dest":   msg.Address,
				"amount": 500000000,
			},
		},
		Signer: client.Signer{Type: client.NoneSignerType},
	}

	res, err := c.ProcessingProcessMessage(&client.ParamsOfProcessMessage{
		MessageEncodeParams: params,
		SendEvents:          false,
	}, nil)
	if err != nil {
		return err
	}
	for _, msg := range res.OutMessages {
		_, err := c.BocParseMessage(&client.ParamsOfParse{
			Boc: msg,
		})
		if err != nil {
			return err
		}
	}
	_, err = c.ProcessingProcessMessage(&client.ParamsOfProcessMessage{
		MessageEncodeParams: params,
		SendEvents:          false,
	}, nil)

	return err
}
