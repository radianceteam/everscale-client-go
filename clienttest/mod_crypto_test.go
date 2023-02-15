package clienttest

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"

	"github.com/radianceteam/everscale-client-go/client"
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

type appSigningBoxMock struct {
	Public  string
	Private string
}

func (app *appSigningBoxMock) GetPublicKeyRequest(_ client.GetPublicKeyParamsOfAppSigningBox) (client.GetPublicKeyResultOfAppSigningBox, error) {
	return client.GetPublicKeyResultOfAppSigningBox{PublicKey: app.Public}, nil
}

func (app *appSigningBoxMock) SignRequest(p client.SignParamsOfAppSigningBox) (client.SignResultOfAppSigningBox, error) {
	seedBytes, err := hex.DecodeString(app.Private)
	if err != nil {
		return client.SignResultOfAppSigningBox{}, err
	}
	privateKey := ed25519.NewKeyFromSeed(seedBytes)
	data, err := base64.StdEncoding.DecodeString(p.Unsigned)
	signature := hex.EncodeToString(ed25519.Sign(privateKey, data))

	return client.SignResultOfAppSigningBox{Signature: signature}, err
}

func (app *appSigningBoxMock) Notify(client.ParamsOfAppSigningBox) {
	panic("notify")
}

func TestModCryptoRegisterSigningBox(t *testing.T) {
	a := assert.New(t)
	c := NewTestClient()
	defer c.Close()

	keys, err := c.CryptoGenerateRandomSignKeys()
	a.NoError(err, "call Client.version")
	a.Len(keys.Public, 64, "public hex len")
	a.Len(keys.Secret, 64, "secret hex len")

	handle, err := c.CryptoRegisterSigningBox(&appSigningBoxMock{
		Private: keys.Secret,
		Public:  keys.Public,
	})

	a.NoError(err, "CryptoRegisterSigningBox")
	a.NotZero(handle.Handle, "CryptoRegisterSigningBox handle")
	keyResult, err := c.CryptoSigningBoxGetPublicKey(&client.RegisteredSigningBox{Handle: handle.Handle})
	a.NoError(err, "CryptoSigningBoxGetPublicKey")
	a.Equal(keyResult.Pubkey, keys.Public, "public key is the same")

	messageToSign := []byte("test message")
	signResult, err := c.CryptoSigningBoxSign(&client.ParamsOfSigningBoxSign{
		SigningBox: handle.Handle,
		Unsigned:   base64.StdEncoding.EncodeToString(messageToSign),
	})
	a.NoError(err, "CryptoSigningBoxSign")
	pubKeyBytes, err := hex.DecodeString(keyResult.Pubkey)
	a.NoError(err, "hex.DecodeString(keyResult.Pubkey)")
	signatureBytes, err := hex.DecodeString(signResult.Signature)
	a.NoError(err, "hex.DecodeString(signResult.Signature)")

	a.True(ed25519.Verify(pubKeyBytes, messageToSign, signatureBytes))
	a.NoError(c.CryptoRemoveSigningBox(&client.RegisteredSigningBox{Handle: handle.Handle}))
}

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
	dict := client.EnglishMnemonicDictionary
	r, err = c.CryptoMnemonicWords(&client.ParamsOfMnemonicWords{Dictionary: &dict})
	a.NoError(err, "call crypto.mnemonic_words")
	a.Len(strings.Split(r.Words, " "), 2048, "default dictionary size")
}
