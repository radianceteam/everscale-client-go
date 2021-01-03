package clienttest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"

	"github.com/radianceteam/ton-client-go/client"
)

func TestModCrypto(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()

	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "call Client.version")
	a.Len(keys.Public, 64, "hex len")
	a.Len(keys.Secret, 64, "hex len")
}

// func TestModCryptoRegisterSigningBox(t *testing.T) {
//	a := assert.New(t)
//	c := NewTestClient()
//	defer c.Close()
//
//	_, err := c.CryptoRegisterSigningBox()
//	a.NoError(err, "CryptoRegisterSigningBox")
//}

func TestModCryptoMnemonicFromRandom(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{})
	a.NoError(err, "call crypto.mnemonic_from_random")
	a.Len(strings.Split(r.Phrase, " "), 12, "default phrase size")
	_, err = c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{WordCount: null.Uint8From(24)})
	a.NoError(err, "call crypto.mnemonic_from_random")
	_, err = c.CryptoMnemonicFromRandom(&client.ParamsOfMnemonicFromRandom{WordCount: null.Uint8From(13)})
	a.Error(err, "bip39 invalid wc")
}

func TestModCryptoMnemonicWords(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()

	r, err := c.CryptoMnemonicWords(&client.ParamsOfMnemonicWords{})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
	r, err = c.CryptoMnemonicWords(&client.ParamsOfMnemonicWords{Dictionary: null.Uint8From(1)})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
}
