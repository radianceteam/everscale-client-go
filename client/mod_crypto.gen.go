package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Aug 21 05:19 UTC
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
	InvalidPublicKeyCryptoErrorCode           = 100
	InvalidSecretKeyCryptoErrorCode           = 101
	InvalidKeyCryptoErrorCode                 = 102
	InvalidFactorizeChallengeCryptoErrorCode  = 106
	InvalidBigIntCryptoErrorCode              = 107
	ScryptFailedCryptoErrorCode               = 108
	InvalidKeySizeCryptoErrorCode             = 109
	NaclSecretBoxFailedCryptoErrorCode        = 110
	NaclBoxFailedCryptoErrorCode              = 111
	NaclSignFailedCryptoErrorCode             = 112
	Bip39InvalidEntropyCryptoErrorCode        = 113
	Bip39InvalidPhraseCryptoErrorCode         = 114
	Bip32InvalidKeyCryptoErrorCode            = 115
	Bip32InvalidDerivePathCryptoErrorCode     = 116
	Bip39InvalidDictionaryCryptoErrorCode     = 117
	Bip39InvalidWordCountCryptoErrorCode      = 118
	MnemonicGenerationFailedCryptoErrorCode   = 119
	MnemonicFromEntropyFailedCryptoErrorCode  = 120
	SigningBoxNotRegisteredCryptoErrorCode    = 121
	InvalidSignatureCryptoErrorCode           = 122
	EncryptionBoxNotRegisteredCryptoErrorCode = 123
	InvalidIvSizeCryptoErrorCode              = 124
	UnsupportedCipherModeCryptoErrorCode      = 125
	CannotCreateCipherCryptoErrorCode         = 126
	EncryptDataErrorCryptoErrorCode           = 127
	DecryptDataErrorCryptoErrorCode           = 128
	IvRequiredCryptoErrorCode                 = 129
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

type EncryptionAlgorithm struct {
	// Should be any of
	// AesParams
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *EncryptionAlgorithm) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case AesParams:
		return json.Marshal(struct {
			AesParams
			Type string `json:"type"`
		}{
			value,
			"AES",
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
		var enumTypeValue AesParams
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

type AesParams struct {
	Mode CipherMode  `json:"mode"`
	Key  string      `json:"key"`
	Iv   null.String `json:"iv"` // optional
}

type AesInfo struct {
	Mode CipherMode  `json:"mode"`
	Iv   null.String `json:"iv"` // optional
}

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
	Nonce     string `json:"nonce"`
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
	// Public key - unprefixed 0-padded to 64 symbols hex string.
	Key string `json:"key"`
}

type ParamsOfMnemonicWords struct {
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
}

type ResultOfMnemonicWords struct {
	// The list of mnemonic words.
	Words string `json:"words"`
}

type ParamsOfMnemonicFromRandom struct {
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
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
	Dictionary null.Uint8 `json:"dictionary"` // optional
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
	Dictionary null.Uint8 `json:"dictionary"` // optional
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
	Dictionary null.Uint8 `json:"dictionary"` // optional
	// Word count.
	WordCount null.Uint8 `json:"word_count"` // optional
}

type ParamsOfHDKeyXPrvFromMnemonic struct {
	// String with seed phrase.
	Phrase string `json:"phrase"`
	// Dictionary identifier.
	Dictionary null.Uint8 `json:"dictionary"` // optional
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

type RegisteredSigningBox struct {
	// Handle of the signing box.
	Handle SigningBoxHandle `json:"handle"`
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

type RegisteredEncryptionBox struct {
	// Handle of the encryption box.
	Handle EncryptionBoxHandle `json:"handle"`
}

// Encryption box callbacks.

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
