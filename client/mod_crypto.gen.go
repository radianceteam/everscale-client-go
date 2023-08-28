package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Aug 23 13:53 UTC
//
// Mod crypto
//
// Crypto functions.

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/volatiletech/null"
)

const (
	InvalidPublicKeyCryptoErrorCode                    = 100
	InvalidSecretKeyCryptoErrorCode                    = 101
	InvalidKeyCryptoErrorCode                          = 102
	InvalidFactorizeChallengeCryptoErrorCode           = 106
	InvalidBigIntCryptoErrorCode                       = 107
	ScryptFailedCryptoErrorCode                        = 108
	InvalidKeySizeCryptoErrorCode                      = 109
	NaclSecretBoxFailedCryptoErrorCode                 = 110
	NaclBoxFailedCryptoErrorCode                       = 111
	NaclSignFailedCryptoErrorCode                      = 112
	Bip39InvalidEntropyCryptoErrorCode                 = 113
	Bip39InvalidPhraseCryptoErrorCode                  = 114
	Bip32InvalidKeyCryptoErrorCode                     = 115
	Bip32InvalidDerivePathCryptoErrorCode              = 116
	Bip39InvalidDictionaryCryptoErrorCode              = 117
	Bip39InvalidWordCountCryptoErrorCode               = 118
	MnemonicGenerationFailedCryptoErrorCode            = 119
	MnemonicFromEntropyFailedCryptoErrorCode           = 120
	SigningBoxNotRegisteredCryptoErrorCode             = 121
	InvalidSignatureCryptoErrorCode                    = 122
	EncryptionBoxNotRegisteredCryptoErrorCode          = 123
	InvalidIvSizeCryptoErrorCode                       = 124
	UnsupportedCipherModeCryptoErrorCode               = 125
	CannotCreateCipherCryptoErrorCode                  = 126
	EncryptDataErrorCryptoErrorCode                    = 127
	DecryptDataErrorCryptoErrorCode                    = 128
	IvRequiredCryptoErrorCode                          = 129
	CryptoBoxNotRegisteredCryptoErrorCode              = 130
	InvalidCryptoBoxTypeCryptoErrorCode                = 131
	CryptoBoxSecretSerializationErrorCryptoErrorCode   = 132
	CryptoBoxSecretDeserializationErrorCryptoErrorCode = 133
	InvalidNonceSizeCryptoErrorCode                    = 134
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[InvalidPublicKeyCryptoErrorCode] = "InvalidPublicKeyCryptoErrorCode"
	errorCodesToErrorTypes[InvalidSecretKeyCryptoErrorCode] = "InvalidSecretKeyCryptoErrorCode"
	errorCodesToErrorTypes[InvalidKeyCryptoErrorCode] = "InvalidKeyCryptoErrorCode"
	errorCodesToErrorTypes[InvalidFactorizeChallengeCryptoErrorCode] = "InvalidFactorizeChallengeCryptoErrorCode"
	errorCodesToErrorTypes[InvalidBigIntCryptoErrorCode] = "InvalidBigIntCryptoErrorCode"
	errorCodesToErrorTypes[ScryptFailedCryptoErrorCode] = "ScryptFailedCryptoErrorCode"
	errorCodesToErrorTypes[InvalidKeySizeCryptoErrorCode] = "InvalidKeySizeCryptoErrorCode"
	errorCodesToErrorTypes[NaclSecretBoxFailedCryptoErrorCode] = "NaclSecretBoxFailedCryptoErrorCode"
	errorCodesToErrorTypes[NaclBoxFailedCryptoErrorCode] = "NaclBoxFailedCryptoErrorCode"
	errorCodesToErrorTypes[NaclSignFailedCryptoErrorCode] = "NaclSignFailedCryptoErrorCode"
	errorCodesToErrorTypes[Bip39InvalidEntropyCryptoErrorCode] = "Bip39InvalidEntropyCryptoErrorCode"
	errorCodesToErrorTypes[Bip39InvalidPhraseCryptoErrorCode] = "Bip39InvalidPhraseCryptoErrorCode"
	errorCodesToErrorTypes[Bip32InvalidKeyCryptoErrorCode] = "Bip32InvalidKeyCryptoErrorCode"
	errorCodesToErrorTypes[Bip32InvalidDerivePathCryptoErrorCode] = "Bip32InvalidDerivePathCryptoErrorCode"
	errorCodesToErrorTypes[Bip39InvalidDictionaryCryptoErrorCode] = "Bip39InvalidDictionaryCryptoErrorCode"
	errorCodesToErrorTypes[Bip39InvalidWordCountCryptoErrorCode] = "Bip39InvalidWordCountCryptoErrorCode"
	errorCodesToErrorTypes[MnemonicGenerationFailedCryptoErrorCode] = "MnemonicGenerationFailedCryptoErrorCode"
	errorCodesToErrorTypes[MnemonicFromEntropyFailedCryptoErrorCode] = "MnemonicFromEntropyFailedCryptoErrorCode"
	errorCodesToErrorTypes[SigningBoxNotRegisteredCryptoErrorCode] = "SigningBoxNotRegisteredCryptoErrorCode"
	errorCodesToErrorTypes[InvalidSignatureCryptoErrorCode] = "InvalidSignatureCryptoErrorCode"
	errorCodesToErrorTypes[EncryptionBoxNotRegisteredCryptoErrorCode] = "EncryptionBoxNotRegisteredCryptoErrorCode"
	errorCodesToErrorTypes[InvalidIvSizeCryptoErrorCode] = "InvalidIvSizeCryptoErrorCode"
	errorCodesToErrorTypes[UnsupportedCipherModeCryptoErrorCode] = "UnsupportedCipherModeCryptoErrorCode"
	errorCodesToErrorTypes[CannotCreateCipherCryptoErrorCode] = "CannotCreateCipherCryptoErrorCode"
	errorCodesToErrorTypes[EncryptDataErrorCryptoErrorCode] = "EncryptDataErrorCryptoErrorCode"
	errorCodesToErrorTypes[DecryptDataErrorCryptoErrorCode] = "DecryptDataErrorCryptoErrorCode"
	errorCodesToErrorTypes[IvRequiredCryptoErrorCode] = "IvRequiredCryptoErrorCode"
	errorCodesToErrorTypes[CryptoBoxNotRegisteredCryptoErrorCode] = "CryptoBoxNotRegisteredCryptoErrorCode"
	errorCodesToErrorTypes[InvalidCryptoBoxTypeCryptoErrorCode] = "InvalidCryptoBoxTypeCryptoErrorCode"
	errorCodesToErrorTypes[CryptoBoxSecretSerializationErrorCryptoErrorCode] = "CryptoBoxSecretSerializationErrorCryptoErrorCode"
	errorCodesToErrorTypes[CryptoBoxSecretDeserializationErrorCryptoErrorCode] = "CryptoBoxSecretDeserializationErrorCryptoErrorCode"
	errorCodesToErrorTypes[InvalidNonceSizeCryptoErrorCode] = "InvalidNonceSizeCryptoErrorCode"
}

type (
	SigningBoxHandle    uint32
	EncryptionBoxHandle uint32
)

// Encryption box information.
type EncryptionBoxInfo struct {
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Hdpath null.String `json:"hdpath"` // optional
	// Cryptographic algorithm, used by this encryption box.
	Algorithm null.String `json:"algorithm"` // optional
	// Options, depends on algorithm and specific encryption box implementation.
	Options json.RawMessage `json:"options"` // optional
	// Public information, depends on algorithm.
	Public json.RawMessage `json:"public"` // optional
}

type AesEncryptionAlgorithm struct {
	Value AesParamsEB `json:"value"`
}

type ChaCha20EncryptionAlgorithm struct {
	Value ChaCha20ParamsEB `json:"value"`
}

type NaclBoxEncryptionAlgorithm struct {
	Value NaclBoxParamsEB `json:"value"`
}

type NaclSecretBoxEncryptionAlgorithm struct {
	Value NaclSecretBoxParamsEB `json:"value"`
}

type EncryptionAlgorithm struct {
	// Should be any of
	// AesEncryptionAlgorithm
	// ChaCha20EncryptionAlgorithm
	// NaclBoxEncryptionAlgorithm
	// NaclSecretBoxEncryptionAlgorithm
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *EncryptionAlgorithm) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case AesEncryptionAlgorithm:
		return json.Marshal(struct {
			AesEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"AES",
		})

	case ChaCha20EncryptionAlgorithm:
		return json.Marshal(struct {
			ChaCha20EncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"ChaCha20",
		})

	case NaclBoxEncryptionAlgorithm:
		return json.Marshal(struct {
			NaclBoxEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"NaclBox",
		})

	case NaclSecretBoxEncryptionAlgorithm:
		return json.Marshal(struct {
			NaclSecretBoxEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"NaclSecretBox",
		})

	default:
		return nil, fmt.Errorf("unsupported type for EncryptionAlgorithm %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *EncryptionAlgorithm) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "AES":
		var enumTypeValue AesEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "ChaCha20":
		var enumTypeValue ChaCha20EncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "NaclBox":
		var enumTypeValue NaclBoxEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "NaclSecretBox":
		var enumTypeValue NaclSecretBoxEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for EncryptionAlgorithm %v", typeDescriptor.Type)
	}

	return nil
}

type CipherMode string

const (
	CbcCipherMode CipherMode = "CBC"
	CfbCipherMode CipherMode = "CFB"
	CtrCipherMode CipherMode = "CTR"
	EcbCipherMode CipherMode = "ECB"
	OfbCipherMode CipherMode = "OFB"
)

type AesParamsEB struct {
	Mode CipherMode  `json:"mode"`
	Key  string      `json:"key"`
	Iv   null.String `json:"iv"` // optional
}

type AesInfo struct {
	Mode CipherMode  `json:"mode"`
	Iv   null.String `json:"iv"` // optional
}

type ChaCha20ParamsEB struct {
	// 256-bit key.
	// Must be encoded with `hex`.
	Key string `json:"key"`
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type NaclBoxParamsEB struct {
	// 256-bit key.
	// Must be encoded with `hex`.
	TheirPublic string `json:"their_public"`
	// 256-bit key.
	// Must be encoded with `hex`.
	Secret string `json:"secret"`
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type NaclSecretBoxParamsEB struct {
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
}

// Crypto Box Secret.

// Creates Crypto Box from a random seed phrase. This option can be used if a developer doesn't want the seed phrase to leave the core library's memory, where it is stored encrypted.
// This type should be used upon the first wallet initialization, all further initializations
// should use `EncryptedSecret` type instead.
//
// Get `encrypted_secret` with `get_crypto_box_info` function and store it on your side.
type RandomSeedPhraseCryptoBoxSecret struct {
	Dictionary MnemonicDictionary `json:"dictionary"`
	Wordcount  uint8              `json:"wordcount"`
}

// Restores crypto box instance from an existing seed phrase. This type should be used when Crypto Box is initialized from a seed phrase, entered by a user.
// This type should be used only upon the first wallet initialization, all further
// initializations should use `EncryptedSecret` type instead.
//
// Get `encrypted_secret` with `get_crypto_box_info` function and store it on your side.
type PredefinedSeedPhraseCryptoBoxSecret struct {
	Phrase     string             `json:"phrase"`
	Dictionary MnemonicDictionary `json:"dictionary"`
	Wordcount  uint8              `json:"wordcount"`
}

// Use this type for wallet reinitializations, when you already have `encrypted_secret` on hands. To get `encrypted_secret`, use `get_crypto_box_info` function after you initialized your crypto box for the first time.
// It is an object, containing seed phrase or private key, encrypted with
// `secret_encryption_salt` and password from `password_provider`.
//
// Note that if you want to change salt or password provider, then you need to reinitialize
// the wallet with `PredefinedSeedPhrase`, then get `EncryptedSecret` via `get_crypto_box_info`,
// store it somewhere, and only after that initialize the wallet with `EncryptedSecret` type.
type EncryptedSecretCryptoBoxSecret struct {
	// It is an object, containing encrypted seed phrase or private key (now we support only seed phrase).
	EncryptedSecret string `json:"encrypted_secret"`
}

type CryptoBoxSecret struct {
	// Should be any of
	// RandomSeedPhraseCryptoBoxSecret
	// PredefinedSeedPhraseCryptoBoxSecret
	// EncryptedSecretCryptoBoxSecret
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *CryptoBoxSecret) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case RandomSeedPhraseCryptoBoxSecret:
		return json.Marshal(struct {
			RandomSeedPhraseCryptoBoxSecret
			Type string `json:"type"`
		}{
			value,
			"RandomSeedPhrase",
		})

	case PredefinedSeedPhraseCryptoBoxSecret:
		return json.Marshal(struct {
			PredefinedSeedPhraseCryptoBoxSecret
			Type string `json:"type"`
		}{
			value,
			"PredefinedSeedPhrase",
		})

	case EncryptedSecretCryptoBoxSecret:
		return json.Marshal(struct {
			EncryptedSecretCryptoBoxSecret
			Type string `json:"type"`
		}{
			value,
			"EncryptedSecret",
		})

	default:
		return nil, fmt.Errorf("unsupported type for CryptoBoxSecret %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *CryptoBoxSecret) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "RandomSeedPhrase":
		var enumTypeValue RandomSeedPhraseCryptoBoxSecret
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "PredefinedSeedPhrase":
		var enumTypeValue PredefinedSeedPhraseCryptoBoxSecret
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "EncryptedSecret":
		var enumTypeValue EncryptedSecretCryptoBoxSecret
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for CryptoBoxSecret %v", typeDescriptor.Type)
	}

	return nil
}

type CryptoBoxHandle uint32

type ChaCha20BoxEncryptionAlgorithm struct {
	Value ChaCha20ParamsCB `json:"value"`
}

type NaclBoxBoxEncryptionAlgorithm struct {
	Value NaclBoxParamsCB `json:"value"`
}

type NaclSecretBoxBoxEncryptionAlgorithm struct {
	Value NaclSecretBoxParamsCB `json:"value"`
}

type BoxEncryptionAlgorithm struct {
	// Should be any of
	// ChaCha20BoxEncryptionAlgorithm
	// NaclBoxBoxEncryptionAlgorithm
	// NaclSecretBoxBoxEncryptionAlgorithm
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BoxEncryptionAlgorithm) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case ChaCha20BoxEncryptionAlgorithm:
		return json.Marshal(struct {
			ChaCha20BoxEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"ChaCha20",
		})

	case NaclBoxBoxEncryptionAlgorithm:
		return json.Marshal(struct {
			NaclBoxBoxEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"NaclBox",
		})

	case NaclSecretBoxBoxEncryptionAlgorithm:
		return json.Marshal(struct {
			NaclSecretBoxBoxEncryptionAlgorithm
			Type string `json:"type"`
		}{
			value,
			"NaclSecretBox",
		})

	default:
		return nil, fmt.Errorf("unsupported type for BoxEncryptionAlgorithm %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BoxEncryptionAlgorithm) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "ChaCha20":
		var enumTypeValue ChaCha20BoxEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "NaclBox":
		var enumTypeValue NaclBoxBoxEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "NaclSecretBox":
		var enumTypeValue NaclSecretBoxBoxEncryptionAlgorithm
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for BoxEncryptionAlgorithm %v", typeDescriptor.Type)
	}

	return nil
}

type ChaCha20ParamsCB struct {
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type NaclBoxParamsCB struct {
	// 256-bit key.
	// Must be encoded with `hex`.
	TheirPublic string `json:"their_public"`
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type NaclSecretBoxParamsCB struct {
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
}

type MnemonicDictionary int

const (

	// TON compatible dictionary.
	TonMnemonicDictionary MnemonicDictionary = 0
	// English BIP-39 dictionary.
	EnglishMnemonicDictionary MnemonicDictionary = 1
	// Chinese simplified BIP-39 dictionary.
	ChineseSimplifiedMnemonicDictionary MnemonicDictionary = 2
	// Chinese traditional BIP-39 dictionary.
	ChineseTraditionalMnemonicDictionary MnemonicDictionary = 3
	// French BIP-39 dictionary.
	FrenchMnemonicDictionary MnemonicDictionary = 4
	// Italian BIP-39 dictionary.
	ItalianMnemonicDictionary MnemonicDictionary = 5
	// Japanese BIP-39 dictionary.
	JapaneseMnemonicDictionary MnemonicDictionary = 6
	// Korean BIP-39 dictionary.
	KoreanMnemonicDictionary MnemonicDictionary = 7
	// Spanish BIP-39 dictionary.
	SpanishMnemonicDictionary MnemonicDictionary = 8
)

type ParamsOfFactorize struct {
	// Hexadecimal representation of u64 composite number.
	Composite string `json:"composite"`
}

type ResultOfFactorize struct {
	// Two factors of composite or empty if composite can't be factorized.
	Factors []string `json:"factors"`
}

type ParamsOfModularPower struct {
	// `base` argument of calculation.
	Base string `json:"base"`
	// `exponent` argument of calculation.
	Exponent string `json:"exponent"`
	// `modulus` argument of calculation.
	Modulus string `json:"modulus"`
}

type ResultOfModularPower struct {
	// Result of modular exponentiation.
	ModularPower string `json:"modular_power"`
}

type ParamsOfTonCrc16 struct {
	// Input data for CRC calculation.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfTonCrc16 struct {
	// Calculated CRC for input data.
	Crc uint16 `json:"crc"`
}

type ParamsOfGenerateRandomBytes struct {
	// Size of random byte array.
	Length uint32 `json:"length"`
}

type ResultOfGenerateRandomBytes struct {
	// Generated bytes encoded in `base64`.
	Bytes string `json:"bytes"`
}

type ParamsOfConvertPublicKeyToTonSafeFormat struct {
	// Public key - 64 symbols hex string.
	PublicKey string `json:"public_key"`
}

type ResultOfConvertPublicKeyToTonSafeFormat struct {
	// Public key represented in TON safe format.
	TonPublicKey string `json:"ton_public_key"`
}

type KeyPair struct {
	// Public key - 64 symbols hex string.
	Public string `json:"public"`
	// Private key - u64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfSign struct {
	// Data that must be signed encoded in `base64`.
	Unsigned string `json:"unsigned"`
	// Sign keys.
	Keys KeyPair `json:"keys"`
}

type ResultOfSign struct {
	// Signed data combined with signature encoded in `base64`.
	Signed string `json:"signed"`
	// Signature encoded in `hex`.
	Signature string `json:"signature"`
}

type ParamsOfVerifySignature struct {
	// Signed data that must be verified encoded in `base64`.
	Signed string `json:"signed"`
	// Signer's public key - 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfVerifySignature struct {
	// Unsigned data encoded in `base64`.
	Unsigned string `json:"unsigned"`
}

type ParamsOfHash struct {
	// Input data for hash calculation.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type ResultOfHash struct {
	// Hash of input `data`.
	// Encoded with 'hex'.
	Hash string `json:"hash"`
}

type ParamsOfScrypt struct {
	// The password bytes to be hashed. Must be encoded with `base64`.
	Password string `json:"password"`
	// Salt bytes that modify the hash to protect against Rainbow table attacks. Must be encoded with `base64`.
	Salt string `json:"salt"`
	// CPU/memory cost parameter.
	LogN uint8 `json:"log_n"`
	// The block size parameter, which fine-tunes sequential memory read size and performance.
	R uint32 `json:"r"`
	// Parallelization parameter.
	P uint32 `json:"p"`
	// Intended output length in octets of the derived key.
	DkLen uint32 `json:"dk_len"`
}

type ResultOfScrypt struct {
	// Derived key.
	// Encoded with `hex`.
	Key string `json:"key"`
}

type ParamsOfNaclSignKeyPairFromSecret struct {
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfNaclSign struct {
	// Data that must be signed encoded in `base64`.
	Unsigned string `json:"unsigned"`
	// Signer's secret key - unprefixed 0-padded to 128 symbols hex string (concatenation of 64 symbols secret and 64 symbols public keys). See `nacl_sign_keypair_from_secret_key`.
	Secret string `json:"secret"`
}

type ResultOfNaclSign struct {
	// Signed data, encoded in `base64`.
	Signed string `json:"signed"`
}

type ParamsOfNaclSignOpen struct {
	// Signed data that must be unsigned.
	// Encoded with `base64`.
	Signed string `json:"signed"`
	// Signer's public key - unprefixed 0-padded to 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfNaclSignOpen struct {
	// Unsigned data, encoded in `base64`.
	Unsigned string `json:"unsigned"`
}

type ResultOfNaclSignDetached struct {
	// Signature encoded in `hex`.
	Signature string `json:"signature"`
}

type ParamsOfNaclSignDetachedVerify struct {
	// Unsigned data that must be verified.
	// Encoded with `base64`.
	Unsigned string `json:"unsigned"`
	// Signature that must be verified.
	// Encoded with `hex`.
	Signature string `json:"signature"`
	// Signer's public key - unprefixed 0-padded to 64 symbols hex string.
	Public string `json:"public"`
}

type ResultOfNaclSignDetachedVerify struct {
	// `true` if verification succeeded or `false` if it failed.
	Succeeded bool `json:"succeeded"`
}

type ParamsOfNaclBoxKeyPairFromSecret struct {
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfNaclBox struct {
	// Data that must be encrypted encoded in `base64`.
	Decrypted string `json:"decrypted"`
	// Nonce, encoded in `hex`.
	Nonce string `json:"nonce"`
	// Receiver's public key - unprefixed 0-padded to 64 symbols hex string.
	TheirPublic string `json:"their_public"`
	// Sender's private key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ResultOfNaclBox struct {
	// Encrypted data encoded in `base64`.
	Encrypted string `json:"encrypted"`
}

type ParamsOfNaclBoxOpen struct {
	// Data that must be decrypted.
	// Encoded with `base64`.
	Encrypted string `json:"encrypted"`
	// Nonce.
	Nonce string `json:"nonce"`
	// Sender's public key - unprefixed 0-padded to 64 symbols hex string.
	TheirPublic string `json:"their_public"`
	// Receiver's private key - unprefixed 0-padded to 64 symbols hex string.
	Secret string `json:"secret"`
}

type ResultOfNaclBoxOpen struct {
	// Decrypted data encoded in `base64`.
	Decrypted string `json:"decrypted"`
}

type ParamsOfNaclSecretBox struct {
	// Data that must be encrypted.
	// Encoded with `base64`.
	Decrypted string `json:"decrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfNaclSecretBoxOpen struct {
	// Data that must be decrypted.
	// Encoded with `base64`.
	Encrypted string `json:"encrypted"`
	// Nonce in `hex`.
	Nonce string `json:"nonce"`
	// Secret key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfMnemonicWords struct {
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
}

type ResultOfMnemonicWords struct {
	// The list of mnemonic words.
	Words string `json:"words"`
}

type ParamsOfMnemonicFromRandom struct {
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicFromRandom struct {
	// String of mnemonic words.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicFromEntropy struct {
	// Entropy bytes.
	// Hex encoded.
	Entropy string `json:"entropy"`
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicFromEntropy struct {
	// Phrase.
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicVerify struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
	// Word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfMnemonicVerify struct {
	// Flag indicating if the mnemonic is valid or not.
	Valid bool `json:"valid"`
}

type ParamsOfMnemonicDeriveSignKeys struct {
	// Phrase.
	Phrase string `json:"phrase"`
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Path null.String `json:"path"` // optional
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
	// Word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ParamsOfHDKeyXPrvFromMnemonic struct {
	// String with seed phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary *MnemonicDictionary `json:"dictionary"` // optional
	// Mnemonic word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ResultOfHDKeyXPrvFromMnemonic struct {
	// Serialized extended master private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
	// Child index (see BIP-0032).
	ChildIndex uint32 `json:"child_index"`
	// Indicates the derivation of hardened/not-hardened key (see BIP-0032).
	Hardened bool `json:"hardened"`
}

type ResultOfHDKeyDeriveFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrvPath struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
	// Derivation path, for instance "m/44'/396'/0'/0/0".
	Path string `json:"path"`
}

type ResultOfHDKeyDeriveFromXPrvPath struct {
	// Derived serialized extended private key.
	XPrv string `json:"xprv"`
}

type ParamsOfHDKeySecretFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ResultOfHDKeySecretFromXPrv struct {
	// Private key - 64 symbols hex string.
	Secret string `json:"secret"`
}

type ParamsOfHDKeyPublicFromXPrv struct {
	// Serialized extended private key.
	XPrv string `json:"xprv"`
}

type ResultOfHDKeyPublicFromXPrv struct {
	// Public key - 64 symbols hex string.
	Public string `json:"public"`
}

type ParamsOfChaCha20 struct {
	// Source data to be encrypted or decrypted.
	// Must be encoded with `base64`.
	Data string `json:"data"`
	// 256-bit key.
	// Must be encoded with `hex`.
	Key string `json:"key"`
	// 96-bit nonce.
	// Must be encoded with `hex`.
	Nonce string `json:"nonce"`
}

type ResultOfChaCha20 struct {
	// Encrypted/decrypted data.
	// Encoded with `base64`.
	Data string `json:"data"`
}

type ParamsOfCreateCryptoBox struct {
	// Salt used for secret encryption. For example, a mobile device can use device ID as salt.
	SecretEncryptionSalt string `json:"secret_encryption_salt"`
	// Cryptobox secret.
	Secret CryptoBoxSecret `json:"secret"`
}

type RegisteredCryptoBox struct {
	Handle CryptoBoxHandle `json:"handle"`
}

// Interface that provides a callback that returns an encrypted password, used for cryptobox secret encryption.
// To secure the password while passing it from application to the library,
// the library generates a temporary key pair, passes the pubkey
// to the passwordProvider, decrypts the received password with private key,
// and deletes the key pair right away.
//
// Application should generate a temporary nacl_box_keypair
// and encrypt the password with naclbox function using nacl_box_keypair.secret
// and encryption_public_key keys + nonce = 24-byte prefix of encryption_public_key.

type GetPasswordParamsOfAppPasswordProvider struct {
	// Temporary library pubkey, that is used on application side for password encryption, along with application temporary private key and nonce. Used for password decryption on library side.
	EncryptionPublicKey string `json:"encryption_public_key"`
}

type ParamsOfAppPasswordProvider struct {
	// Should be any of
	// GetPasswordParamsOfAppPasswordProvider
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppPasswordProvider) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetPasswordParamsOfAppPasswordProvider:
		return json.Marshal(struct {
			GetPasswordParamsOfAppPasswordProvider
			Type string `json:"type"`
		}{
			value,
			"GetPassword",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppPasswordProvider %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppPasswordProvider) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetPassword":
		var enumTypeValue GetPasswordParamsOfAppPasswordProvider
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppPasswordProvider %v", typeDescriptor.Type)
	}

	return nil
}

type GetPasswordResultOfAppPasswordProvider struct {
	// Password, encrypted and encoded to base64. Crypto box uses this password to decrypt its secret (seed phrase).
	EncryptedPassword string `json:"encrypted_password"`
	// Hex encoded public key of a temporary key pair, used for password encryption on application side.
	// Used together with `encryption_public_key` to decode `encrypted_password`.
	AppEncryptionPubkey string `json:"app_encryption_pubkey"`
}

type ResultOfAppPasswordProvider struct {
	// Should be any of
	// GetPasswordResultOfAppPasswordProvider
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppPasswordProvider) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetPasswordResultOfAppPasswordProvider:
		return json.Marshal(struct {
			GetPasswordResultOfAppPasswordProvider
			Type string `json:"type"`
		}{
			value,
			"GetPassword",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppPasswordProvider %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppPasswordProvider) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetPassword":
		var enumTypeValue GetPasswordResultOfAppPasswordProvider
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ResultOfAppPasswordProvider %v", typeDescriptor.Type)
	}

	return nil
}

type ResultOfGetCryptoBoxInfo struct {
	// Secret (seed phrase) encrypted with salt and password.
	EncryptedSecret string `json:"encrypted_secret"`
}

type ResultOfGetCryptoBoxSeedPhrase struct {
	Phrase     string             `json:"phrase"`
	Dictionary MnemonicDictionary `json:"dictionary"`
	Wordcount  uint8              `json:"wordcount"`
}

type ParamsOfGetSigningBoxFromCryptoBox struct {
	// Crypto Box Handle.
	Handle uint32 `json:"handle"`
	// HD key derivation path.
	// By default, Everscale HD path is used.
	Hdpath null.String `json:"hdpath"` // optional
	// Store derived secret for this lifetime (in ms). The timer starts after each signing box operation. Secrets will be deleted immediately after each signing box operation, if this value is not set.
	SecretLifetime null.Uint32 `json:"secret_lifetime"` // optional
}

type RegisteredSigningBox struct {
	// Handle of the signing box.
	Handle SigningBoxHandle `json:"handle"`
}

type ParamsOfGetEncryptionBoxFromCryptoBox struct {
	// Crypto Box Handle.
	Handle uint32 `json:"handle"`
	// HD key derivation path.
	// By default, Everscale HD path is used.
	Hdpath null.String `json:"hdpath"` // optional
	// Encryption algorithm.
	Algorithm BoxEncryptionAlgorithm `json:"algorithm"`
	// Store derived secret for encryption algorithm for this lifetime (in ms). The timer starts after each encryption box operation. Secrets will be deleted (overwritten with zeroes) after each encryption operation, if this value is not set.
	SecretLifetime null.Uint32 `json:"secret_lifetime"` // optional
}

type RegisteredEncryptionBox struct {
	// Handle of the encryption box.
	Handle EncryptionBoxHandle `json:"handle"`
}

// Signing box callbacks.

// Get signing box public key.
type GetPublicKeyParamsOfAppSigningBox struct{}

// Sign data.
type SignParamsOfAppSigningBox struct {
	// Data to sign encoded as base64.
	Unsigned string `json:"unsigned"`
}

type ParamsOfAppSigningBox struct {
	// Should be any of
	// GetPublicKeyParamsOfAppSigningBox
	// SignParamsOfAppSigningBox
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppSigningBox) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetPublicKeyParamsOfAppSigningBox:
		return json.Marshal(struct {
			GetPublicKeyParamsOfAppSigningBox
			Type string `json:"type"`
		}{
			value,
			"GetPublicKey",
		})

	case SignParamsOfAppSigningBox:
		return json.Marshal(struct {
			SignParamsOfAppSigningBox
			Type string `json:"type"`
		}{
			value,
			"Sign",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppSigningBox %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppSigningBox) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetPublicKey":
		var enumTypeValue GetPublicKeyParamsOfAppSigningBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Sign":
		var enumTypeValue SignParamsOfAppSigningBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppSigningBox %v", typeDescriptor.Type)
	}

	return nil
}

// Returning values from signing box callbacks.

// Result of getting public key.
type GetPublicKeyResultOfAppSigningBox struct {
	// Signing box public key.
	PublicKey string `json:"public_key"`
}

// Result of signing data.
type SignResultOfAppSigningBox struct {
	// Data signature encoded as hex.
	Signature string `json:"signature"`
}

type ResultOfAppSigningBox struct {
	// Should be any of
	// GetPublicKeyResultOfAppSigningBox
	// SignResultOfAppSigningBox
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppSigningBox) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetPublicKeyResultOfAppSigningBox:
		return json.Marshal(struct {
			GetPublicKeyResultOfAppSigningBox
			Type string `json:"type"`
		}{
			value,
			"GetPublicKey",
		})

	case SignResultOfAppSigningBox:
		return json.Marshal(struct {
			SignResultOfAppSigningBox
			Type string `json:"type"`
		}{
			value,
			"Sign",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppSigningBox %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppSigningBox) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetPublicKey":
		var enumTypeValue GetPublicKeyResultOfAppSigningBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Sign":
		var enumTypeValue SignResultOfAppSigningBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ResultOfAppSigningBox %v", typeDescriptor.Type)
	}

	return nil
}

type ResultOfSigningBoxGetPublicKey struct {
	// Public key of signing box.
	// Encoded with hex.
	Pubkey string `json:"pubkey"`
}

type ParamsOfSigningBoxSign struct {
	// Signing Box handle.
	SigningBox SigningBoxHandle `json:"signing_box"`
	// Unsigned user data.
	// Must be encoded with `base64`.
	Unsigned string `json:"unsigned"`
}

type ResultOfSigningBoxSign struct {
	// Data signature.
	// Encoded with `hex`.
	Signature string `json:"signature"`
}

// Interface for data encryption/decryption.

// Get encryption box info.
type GetInfoParamsOfAppEncryptionBox struct{}

// Encrypt data.
type EncryptParamsOfAppEncryptionBox struct {
	// Data, encoded in Base64.
	Data string `json:"data"`
}

// Decrypt data.
type DecryptParamsOfAppEncryptionBox struct {
	// Data, encoded in Base64.
	Data string `json:"data"`
}

type ParamsOfAppEncryptionBox struct {
	// Should be any of
	// GetInfoParamsOfAppEncryptionBox
	// EncryptParamsOfAppEncryptionBox
	// DecryptParamsOfAppEncryptionBox
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppEncryptionBox) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetInfoParamsOfAppEncryptionBox:
		return json.Marshal(struct {
			GetInfoParamsOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"GetInfo",
		})

	case EncryptParamsOfAppEncryptionBox:
		return json.Marshal(struct {
			EncryptParamsOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"Encrypt",
		})

	case DecryptParamsOfAppEncryptionBox:
		return json.Marshal(struct {
			DecryptParamsOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"Decrypt",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppEncryptionBox %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppEncryptionBox) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetInfo":
		var enumTypeValue GetInfoParamsOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Encrypt":
		var enumTypeValue EncryptParamsOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Decrypt":
		var enumTypeValue DecryptParamsOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppEncryptionBox %v", typeDescriptor.Type)
	}

	return nil
}

// Returning values from signing box callbacks.

// Result of getting encryption box info.
type GetInfoResultOfAppEncryptionBox struct {
	Info EncryptionBoxInfo `json:"info"`
}

// Result of encrypting data.
type EncryptResultOfAppEncryptionBox struct {
	// Encrypted data, encoded in Base64.
	Data string `json:"data"`
}

// Result of decrypting data.
type DecryptResultOfAppEncryptionBox struct {
	// Decrypted data, encoded in Base64.
	Data string `json:"data"`
}

type ResultOfAppEncryptionBox struct {
	// Should be any of
	// GetInfoResultOfAppEncryptionBox
	// EncryptResultOfAppEncryptionBox
	// DecryptResultOfAppEncryptionBox
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppEncryptionBox) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case GetInfoResultOfAppEncryptionBox:
		return json.Marshal(struct {
			GetInfoResultOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"GetInfo",
		})

	case EncryptResultOfAppEncryptionBox:
		return json.Marshal(struct {
			EncryptResultOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"Encrypt",
		})

	case DecryptResultOfAppEncryptionBox:
		return json.Marshal(struct {
			DecryptResultOfAppEncryptionBox
			Type string `json:"type"`
		}{
			value,
			"Decrypt",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppEncryptionBox %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppEncryptionBox) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "GetInfo":
		var enumTypeValue GetInfoResultOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Encrypt":
		var enumTypeValue EncryptResultOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Decrypt":
		var enumTypeValue DecryptResultOfAppEncryptionBox
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ResultOfAppEncryptionBox %v", typeDescriptor.Type)
	}

	return nil
}

type ParamsOfEncryptionBoxGetInfo struct {
	// Encryption box handle.
	EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
}

type ResultOfEncryptionBoxGetInfo struct {
	// Encryption box information.
	Info EncryptionBoxInfo `json:"info"`
}

type ParamsOfEncryptionBoxEncrypt struct {
	// Encryption box handle.
	EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
	// Data to be encrypted, encoded in Base64.
	Data string `json:"data"`
}

type ResultOfEncryptionBoxEncrypt struct {
	// Encrypted data, encoded in Base64.
	// Padded to cipher block size.
	Data string `json:"data"`
}

type ParamsOfEncryptionBoxDecrypt struct {
	// Encryption box handle.
	EncryptionBox EncryptionBoxHandle `json:"encryption_box"`
	// Data to be decrypted, encoded in Base64.
	Data string `json:"data"`
}

type ResultOfEncryptionBoxDecrypt struct {
	// Decrypted data, encoded in Base64.
	Data string `json:"data"`
}

type ParamsOfCreateEncryptionBox struct {
	// Encryption algorithm specifier including cipher parameters (key, IV, etc).
	Algorithm EncryptionAlgorithm `json:"algorithm"`
}

// Integer factorization.
// Performs prime factorization – decomposition of a composite number
// into a product of smaller prime integers (factors).
// See [https://en.wikipedia.org/wiki/Integer_factorization].
func (c *Client) CryptoFactorize(p *ParamsOfFactorize) (*ResultOfFactorize, error) {
	result := new(ResultOfFactorize)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.factorize", p, result)

	return result, err
}

// Modular exponentiation.
// Performs modular exponentiation for big integers (`base`^`exponent` mod `modulus`).
// See [https://en.wikipedia.org/wiki/Modular_exponentiation].
func (c *Client) CryptoModularPower(p *ParamsOfModularPower) (*ResultOfModularPower, error) {
	result := new(ResultOfModularPower)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.modular_power", p, result)

	return result, err
}

// Calculates CRC16 using TON algorithm.
func (c *Client) CryptoTonCrc16(p *ParamsOfTonCrc16) (*ResultOfTonCrc16, error) {
	result := new(ResultOfTonCrc16)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.ton_crc16", p, result)

	return result, err
}

// Generates random byte array of the specified length and returns it in `base64` format.
func (c *Client) CryptoGenerateRandomBytes(p *ParamsOfGenerateRandomBytes) (*ResultOfGenerateRandomBytes, error) {
	result := new(ResultOfGenerateRandomBytes)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_bytes", p, result)

	return result, err
}

// Converts public key to ton safe_format.
func (c *Client) CryptoConvertPublicKeyToTonSafeFormat(p *ParamsOfConvertPublicKeyToTonSafeFormat) (*ResultOfConvertPublicKeyToTonSafeFormat, error) {
	result := new(ResultOfConvertPublicKeyToTonSafeFormat)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.convert_public_key_to_ton_safe_format", p, result)

	return result, err
}

// Generates random ed25519 key pair.
func (c *Client) CryptoGenerateRandomSignKeys() (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_sign_keys", nil, result)

	return result, err
}

// Signs a data using the provided keys.
func (c *Client) CryptoSign(p *ParamsOfSign) (*ResultOfSign, error) {
	result := new(ResultOfSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sign", p, result)

	return result, err
}

// Verifies signed data using the provided public key. Raises error if verification is failed.
func (c *Client) CryptoVerifySignature(p *ParamsOfVerifySignature) (*ResultOfVerifySignature, error) {
	result := new(ResultOfVerifySignature)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.verify_signature", p, result)

	return result, err
}

// Calculates SHA256 hash of the specified data.
func (c *Client) CryptoSha256(p *ParamsOfHash) (*ResultOfHash, error) {
	result := new(ResultOfHash)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha256", p, result)

	return result, err
}

// Calculates SHA512 hash of the specified data.
func (c *Client) CryptoSha512(p *ParamsOfHash) (*ResultOfHash, error) {
	result := new(ResultOfHash)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.sha512", p, result)

	return result, err
}

// Perform `scrypt` encryption.
// Derives key from `password` and `key` using `scrypt` algorithm.
// See [https://en.wikipedia.org/wiki/Scrypt].
//
// # Arguments
// - `log_n` - The log2 of the Scrypt parameter `N`
// - `r` - The Scrypt parameter `r`
// - `p` - The Scrypt parameter `p`
// # Conditions
// - `log_n` must be less than `64`
// - `r` must be greater than `0` and less than or equal to `4294967295`
// - `p` must be greater than `0` and less than `4294967295`
// # Recommended values sufficient for most use-cases
// - `log_n = 15` (`n = 32768`)
// - `r = 8`
// - `p = 1`.
func (c *Client) CryptoScrypt(p *ParamsOfScrypt) (*ResultOfScrypt, error) {
	result := new(ResultOfScrypt)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.scrypt", p, result)

	return result, err
}

// Generates a key pair for signing from the secret key.
// **NOTE:** In the result the secret key is actually the concatenation
// of secret and public keys (128 symbols hex string) by design of [NaCL](http://nacl.cr.yp.to/sign.html).
// See also [the stackexchange question](https://crypto.stackexchange.com/questions/54353/).
func (c *Client) CryptoNaclSignKeypairFromSecretKey(p *ParamsOfNaclSignKeyPairFromSecret) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_keypair_from_secret_key", p, result)

	return result, err
}

// Signs data using the signer's secret key.
func (c *Client) CryptoNaclSign(p *ParamsOfNaclSign) (*ResultOfNaclSign, error) {
	result := new(ResultOfNaclSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign", p, result)

	return result, err
}

// Verifies the signature and returns the unsigned message.
// Verifies the signature in `signed` using the signer's public key `public`
// and returns the message `unsigned`.
//
// If the signature fails verification, crypto_sign_open raises an exception.
func (c *Client) CryptoNaclSignOpen(p *ParamsOfNaclSignOpen) (*ResultOfNaclSignOpen, error) {
	result := new(ResultOfNaclSignOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_open", p, result)

	return result, err
}

// Signs the message using the secret key and returns a signature.
// Signs the message `unsigned` using the secret key `secret`
// and returns a signature `signature`.
func (c *Client) CryptoNaclSignDetached(p *ParamsOfNaclSign) (*ResultOfNaclSignDetached, error) {
	result := new(ResultOfNaclSignDetached)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_detached", p, result)

	return result, err
}

// Verifies the signature with public key and `unsigned` data.
func (c *Client) CryptoNaclSignDetachedVerify(p *ParamsOfNaclSignDetachedVerify) (*ResultOfNaclSignDetachedVerify, error) {
	result := new(ResultOfNaclSignDetachedVerify)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_sign_detached_verify", p, result)

	return result, err
}

// Generates a random NaCl key pair.
func (c *Client) CryptoNaclBoxKeypair() (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair", nil, result)

	return result, err
}

// Generates key pair from a secret key.
func (c *Client) CryptoNaclBoxKeypairFromSecretKey(p *ParamsOfNaclBoxKeyPairFromSecret) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_keypair_from_secret_key", p, result)

	return result, err
}

// Public key authenticated encryption.
// Encrypt and authenticate a message using the senders secret key, the receivers public
// key, and a nonce.
func (c *Client) CryptoNaclBox(p *ParamsOfNaclBox) (*ResultOfNaclBox, error) {
	result := new(ResultOfNaclBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box", p, result)

	return result, err
}

// Decrypt and verify the cipher text using the receivers secret key, the senders public key, and the nonce.
func (c *Client) CryptoNaclBoxOpen(p *ParamsOfNaclBoxOpen) (*ResultOfNaclBoxOpen, error) {
	result := new(ResultOfNaclBoxOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_box_open", p, result)

	return result, err
}

// Encrypt and authenticate message using nonce and secret key.
func (c *Client) CryptoNaclSecretBox(p *ParamsOfNaclSecretBox) (*ResultOfNaclBox, error) {
	result := new(ResultOfNaclBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box", p, result)

	return result, err
}

// Decrypts and verifies cipher text using `nonce` and secret `key`.
func (c *Client) CryptoNaclSecretBoxOpen(p *ParamsOfNaclSecretBoxOpen) (*ResultOfNaclBoxOpen, error) {
	result := new(ResultOfNaclBoxOpen)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.nacl_secret_box_open", p, result)

	return result, err
}

// Prints the list of words from the specified dictionary.
func (c *Client) CryptoMnemonicWords(p *ParamsOfMnemonicWords) (*ResultOfMnemonicWords, error) {
	result := new(ResultOfMnemonicWords)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_words", p, result)

	return result, err
}

// Generates a random mnemonic from the specified dictionary and word count.
func (c *Client) CryptoMnemonicFromRandom(p *ParamsOfMnemonicFromRandom) (*ResultOfMnemonicFromRandom, error) {
	result := new(ResultOfMnemonicFromRandom)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_random", p, result)

	return result, err
}

// Generates mnemonic from pre-generated entropy.
func (c *Client) CryptoMnemonicFromEntropy(p *ParamsOfMnemonicFromEntropy) (*ResultOfMnemonicFromEntropy, error) {
	result := new(ResultOfMnemonicFromEntropy)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_from_entropy", p, result)

	return result, err
}

// Validates a mnemonic phrase.
// The phrase supplied will be checked for word length and validated according to the checksum
// specified in BIP0039.
func (c *Client) CryptoMnemonicVerify(p *ParamsOfMnemonicVerify) (*ResultOfMnemonicVerify, error) {
	result := new(ResultOfMnemonicVerify)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_verify", p, result)

	return result, err
}

// Derives a key pair for signing from the seed phrase.
// Validates the seed phrase, generates master key and then derives
// the key pair from the master key and the specified path.
func (c *Client) CryptoMnemonicDeriveSignKeys(p *ParamsOfMnemonicDeriveSignKeys) (*KeyPair, error) {
	result := new(KeyPair)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.mnemonic_derive_sign_keys", p, result)

	return result, err
}

// Generates an extended master private key that will be the root for all the derived keys.
func (c *Client) CryptoHdkeyXprvFromMnemonic(p *ParamsOfHDKeyXPrvFromMnemonic) (*ResultOfHDKeyXPrvFromMnemonic, error) {
	result := new(ResultOfHDKeyXPrvFromMnemonic)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_xprv_from_mnemonic", p, result)

	return result, err
}

// Returns extended private key derived from the specified extended private key and child index.
func (c *Client) CryptoHdkeyDeriveFromXprv(p *ParamsOfHDKeyDeriveFromXPrv) (*ResultOfHDKeyDeriveFromXPrv, error) {
	result := new(ResultOfHDKeyDeriveFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv", p, result)

	return result, err
}

// Derives the extended private key from the specified key and path.
func (c *Client) CryptoHdkeyDeriveFromXprvPath(p *ParamsOfHDKeyDeriveFromXPrvPath) (*ResultOfHDKeyDeriveFromXPrvPath, error) {
	result := new(ResultOfHDKeyDeriveFromXPrvPath)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_derive_from_xprv_path", p, result)

	return result, err
}

// Extracts the private key from the serialized extended private key.
func (c *Client) CryptoHdkeySecretFromXprv(p *ParamsOfHDKeySecretFromXPrv) (*ResultOfHDKeySecretFromXPrv, error) {
	result := new(ResultOfHDKeySecretFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_secret_from_xprv", p, result)

	return result, err
}

// Extracts the public key from the serialized extended private key.
func (c *Client) CryptoHdkeyPublicFromXprv(p *ParamsOfHDKeyPublicFromXPrv) (*ResultOfHDKeyPublicFromXPrv, error) {
	result := new(ResultOfHDKeyPublicFromXPrv)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.hdkey_public_from_xprv", p, result)

	return result, err
}

// Performs symmetric `chacha20` encryption.
func (c *Client) CryptoChacha20(p *ParamsOfChaCha20) (*ResultOfChaCha20, error) {
	result := new(ResultOfChaCha20)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.chacha20", p, result)

	return result, err
}

// Creates a Crypto Box instance.
// Crypto Box is a root crypto object, that encapsulates some secret (seed phrase usually)
// in encrypted form and acts as a factory for all crypto primitives used in SDK:
// keys for signing and encryption, derived from this secret.
//
// Crypto Box encrypts original Seed Phrase with salt and password that is retrieved
// from `password_provider` callback, implemented on Application side.
//
// When used, decrypted secret shows up in core library's memory for a very short period
// of time and then is immediately overwritten with zeroes.

func (c *Client) CryptoCreateCryptoBox(p *ParamsOfCreateCryptoBox, app AppPasswordProvider) (*RegisteredCryptoBox, error) { // nolint dupl
	result := new(RegisteredCryptoBox)
	responses, err := c.dllClient.resultsChannel("crypto.create_crypto_box", p)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestCryptoCreateCryptoBox(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestCryptoCreateCryptoBox(payload []byte, app AppPasswordProvider) { // nolint dupl
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppPasswordProvider
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	var appResponse interface{}
	// appResponse, err := app.Request(appParams)

	switch value := (appParams.EnumTypeValue).(type) {
	case GetPasswordParamsOfAppPasswordProvider:
		appResponse, err = app.GetPasswordRequest(value)

	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue)
	}

	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&ResultOfAppPasswordProvider{EnumTypeValue: appResponse})
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err == nil || errors.Is(err, ErrContextIsClosed) {
		return
	}
	panic(err)
}

// Removes Crypto Box. Clears all secret data.
func (c *Client) CryptoRemoveCryptoBox(p *RegisteredCryptoBox) error {
	_, err := c.dllClient.waitErrorOrResult("crypto.remove_crypto_box", p)

	return err
}

// Get Crypto Box Info. Used to get `encrypted_secret` that should be used for all the cryptobox initializations except the first one.
func (c *Client) CryptoGetCryptoBoxInfo(p *RegisteredCryptoBox) (*ResultOfGetCryptoBoxInfo, error) {
	result := new(ResultOfGetCryptoBoxInfo)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_crypto_box_info", p, result)

	return result, err
}

// Get Crypto Box Seed Phrase.
// Attention! Store this data in your application for a very short period of time and overwrite it with zeroes ASAP.
func (c *Client) CryptoGetCryptoBoxSeedPhrase(p *RegisteredCryptoBox) (*ResultOfGetCryptoBoxSeedPhrase, error) {
	result := new(ResultOfGetCryptoBoxSeedPhrase)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_crypto_box_seed_phrase", p, result)

	return result, err
}

// Get handle of Signing Box derived from Crypto Box.
func (c *Client) CryptoGetSigningBoxFromCryptoBox(p *ParamsOfGetSigningBoxFromCryptoBox) (*RegisteredSigningBox, error) {
	result := new(RegisteredSigningBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_signing_box_from_crypto_box", p, result)

	return result, err
}

// Gets Encryption Box from Crypto Box.
// Derives encryption keypair from cryptobox secret and hdpath and
// stores it in cache for `secret_lifetime`
// or until explicitly cleared by `clear_crypto_box_secret_cache` method.
// If `secret_lifetime` is not specified - overwrites encryption secret with zeroes immediately after
// encryption operation.
func (c *Client) CryptoGetEncryptionBoxFromCryptoBox(p *ParamsOfGetEncryptionBoxFromCryptoBox) (*RegisteredEncryptionBox, error) {
	result := new(RegisteredEncryptionBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_encryption_box_from_crypto_box", p, result)

	return result, err
}

// Removes cached secrets (overwrites with zeroes) from all signing and encryption boxes, derived from crypto box.
func (c *Client) CryptoClearCryptoBoxSecretCache(p *RegisteredCryptoBox) error {
	_, err := c.dllClient.waitErrorOrResult("crypto.clear_crypto_box_secret_cache", p)

	return err
}

// Register an application implemented signing box.

func (c *Client) CryptoRegisterSigningBox(app AppSigningBox) (*RegisteredSigningBox, error) { // nolint dupl
	result := new(RegisteredSigningBox)
	responses, err := c.dllClient.resultsChannel("crypto.register_signing_box", nil)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestCryptoRegisterSigningBox(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestCryptoRegisterSigningBox(payload []byte, app AppSigningBox) { // nolint dupl
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppSigningBox
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	var appResponse interface{}
	// appResponse, err := app.Request(appParams)

	switch value := (appParams.EnumTypeValue).(type) {
	case GetPublicKeyParamsOfAppSigningBox:
		appResponse, err = app.GetPublicKeyRequest(value)

	case SignParamsOfAppSigningBox:
		appResponse, err = app.SignRequest(value)

	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue)
	}

	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&ResultOfAppSigningBox{EnumTypeValue: appResponse})
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err == nil || errors.Is(err, ErrContextIsClosed) {
		return
	}
	panic(err)
}

// Creates a default signing box implementation.
func (c *Client) CryptoGetSigningBox(p *KeyPair) (*RegisteredSigningBox, error) {
	result := new(RegisteredSigningBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.get_signing_box", p, result)

	return result, err
}

// Returns public key of signing key pair.
func (c *Client) CryptoSigningBoxGetPublicKey(p *RegisteredSigningBox) (*ResultOfSigningBoxGetPublicKey, error) {
	result := new(ResultOfSigningBoxGetPublicKey)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.signing_box_get_public_key", p, result)

	return result, err
}

// Returns signed user data.
func (c *Client) CryptoSigningBoxSign(p *ParamsOfSigningBoxSign) (*ResultOfSigningBoxSign, error) {
	result := new(ResultOfSigningBoxSign)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.signing_box_sign", p, result)

	return result, err
}

// Removes signing box from SDK.
func (c *Client) CryptoRemoveSigningBox(p *RegisteredSigningBox) error {
	_, err := c.dllClient.waitErrorOrResult("crypto.remove_signing_box", p)

	return err
}

// Register an application implemented encryption box.

func (c *Client) CryptoRegisterEncryptionBox(app AppEncryptionBox) (*RegisteredEncryptionBox, error) { // nolint dupl
	result := new(RegisteredEncryptionBox)
	responses, err := c.dllClient.resultsChannel("crypto.register_encryption_box", nil)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestCryptoRegisterEncryptionBox(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestCryptoRegisterEncryptionBox(payload []byte, app AppEncryptionBox) { // nolint dupl
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppEncryptionBox
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	var appResponse interface{}
	// appResponse, err := app.Request(appParams)

	switch value := (appParams.EnumTypeValue).(type) {
	case GetInfoParamsOfAppEncryptionBox:
		appResponse, err = app.GetInfoRequest(value)

	case EncryptParamsOfAppEncryptionBox:
		appResponse, err = app.EncryptRequest(value)

	case DecryptParamsOfAppEncryptionBox:
		appResponse, err = app.DecryptRequest(value)

	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue)
	}

	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&ResultOfAppEncryptionBox{EnumTypeValue: appResponse})
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err == nil || errors.Is(err, ErrContextIsClosed) {
		return
	}
	panic(err)
}

// Removes encryption box from SDK.
func (c *Client) CryptoRemoveEncryptionBox(p *RegisteredEncryptionBox) error {
	_, err := c.dllClient.waitErrorOrResult("crypto.remove_encryption_box", p)

	return err
}

// Queries info from the given encryption box.
func (c *Client) CryptoEncryptionBoxGetInfo(p *ParamsOfEncryptionBoxGetInfo) (*ResultOfEncryptionBoxGetInfo, error) {
	result := new(ResultOfEncryptionBoxGetInfo)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.encryption_box_get_info", p, result)

	return result, err
}

// Encrypts data using given encryption box Note.
// Block cipher algorithms pad data to cipher block size so encrypted data can be longer then original data. Client should store the original data size after encryption and use it after
// decryption to retrieve the original data from decrypted data.
func (c *Client) CryptoEncryptionBoxEncrypt(p *ParamsOfEncryptionBoxEncrypt) (*ResultOfEncryptionBoxEncrypt, error) {
	result := new(ResultOfEncryptionBoxEncrypt)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.encryption_box_encrypt", p, result)

	return result, err
}

// Decrypts data using given encryption box Note.
// Block cipher algorithms pad data to cipher block size so encrypted data can be longer then original data. Client should store the original data size after encryption and use it after
// decryption to retrieve the original data from decrypted data.
func (c *Client) CryptoEncryptionBoxDecrypt(p *ParamsOfEncryptionBoxDecrypt) (*ResultOfEncryptionBoxDecrypt, error) {
	result := new(ResultOfEncryptionBoxDecrypt)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.encryption_box_decrypt", p, result)

	return result, err
}

// Creates encryption box with specified algorithm.
func (c *Client) CryptoCreateEncryptionBox(p *ParamsOfCreateEncryptionBox) (*RegisteredEncryptionBox, error) {
	result := new(RegisteredEncryptionBox)

	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.create_encryption_box", p, result)

	return result, err
}
