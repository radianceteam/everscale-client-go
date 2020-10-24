package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func newTestClient() *Client {
	c, err := NewClient(Config{
		Network: NetworkConfig{ServerAddress: null.StringFrom("net.ton.dev")},
	})
	if err != nil {
		panic(err)
	}
	return c
}

func TestModCrypto(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "call Client.version")
	a.Len(keys.Public, 64, "hex len")
	a.Len(keys.Secret, 64, "hex len")
}

func TestModCryptoMnemonicFromRandom(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicFromRandom(&ParamsOfMnemonicFromRandom{})
	a.NoError(err, "call crypto.mnemonic_from_random")
	a.Len(strings.Split(r.Phrase, " "), 12, "default phrase size")
	r, err = c.CryptoMnemonicFromRandom(&ParamsOfMnemonicFromRandom{WordCount: null.IntFrom(24)})
	a.NoError(err, "call crypto.mnemonic_from_random")
	r, err = c.CryptoMnemonicFromRandom(&ParamsOfMnemonicFromRandom{WordCount: null.IntFrom(13)})
	a.Error(err, "bip39 invalid wc")
}

func TestModCryptoMnemonicWords(t *testing.T) {
	a := assert.New(t)
	c := newTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicWords(&ParamsOfMnemonicWords{})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
	r, err = c.CryptoMnemonicWords(&ParamsOfMnemonicWords{Dictionary: null.IntFrom(1)})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
}
