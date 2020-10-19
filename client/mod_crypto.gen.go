package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.682304 +0000 UTC
// Mod crypto
//  Crypto functions.
//  Crypto functions.

type SigningBoxHandle struct {
	int `json:""`
}

type ParamsOfFactorize struct {
	Composite string `json:"composite"`
}

type ResultOfFactorize struct {
	Factors []string `json:"factors"`
}

type ParamsOfModularPower struct {
	Base     string `json:"base"`
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

type ResultOfModularPower struct {
	ModularPower string `json:"modular_power"`
}

type ParamsOfTonCrc16 struct {
	Data string `json:"data"`
}

type ResultOfTonCrc16 struct {
	Crc int `json:"crc"`
}

type ParamsOfGenerateRandomBytes struct {
	Length int `json:"length"`
}

type ResultOfGenerateRandomBytes struct {
	Bytes string `json:"bytes"`
}

type ParamsOfConvertPublicKeyToTonSafeFormat struct {
	PublicKey string `json:"public_key"`
}

type ResultOfConvertPublicKeyToTonSafeFormat struct {
	TonPublicKey string `json:"ton_public_key"`
}

type KeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

type ParamsOfSign struct {
	Unsigned string  `json:"unsigned"`
	Keys     KeyPair `json:"keys"`
}

type ResultOfSign struct {
	Signed    string `json:"signed"`
	Signature string `json:"signature"`
}

type ParamsOfVerifySignature struct {
	Signed string `json:"signed"`
	Public string `json:"public"`
}

type ResultOfVerifySignature struct {
	Unsigned string `json:"unsigned"`
}

type ParamsOfHash struct {
	Data string `json:"data"`
}

type ResultOfHash struct {
	Hash string `json:"hash"`
}

type ParamsOfScrypt struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
	LogN     int    `json:"log_n"`
	R        int    `json:"r"`
	P        int    `json:"p"`
	DkLen    int    `json:"dk_len"`
}

type ResultOfScrypt struct {
	Key string `json:"key"`
}

type ParamsOfNaclSignKeyPairFromSecret struct {
	Secret string `json:"secret"`
}

type ParamsOfNaclSign struct {
	Unsigned string `json:"unsigned"`
	Secret   string `json:"secret"`
}

type ResultOfNaclSign struct {
	Signed string `json:"signed"`
}

type ParamsOfNaclSignOpen struct {
	Signed string `json:"signed"`
	Public string `json:"public"`
}

type ResultOfNaclSignOpen struct {
	Unsigned string `json:"unsigned"`
}

type ResultOfNaclSignDetached struct {
	Signature string `json:"signature"`
}

type ParamsOfNaclBoxKeyPairFromSecret struct {
	Secret string `json:"secret"`
}

type ParamsOfNaclBox struct {
	Decrypted   string `json:"decrypted"`
	Nonce       string `json:"nonce"`
	TheirPublic string `json:"their_public"`
	Secret      string `json:"secret"`
}

type ResultOfNaclBox struct {
	Encrypted string `json:"encrypted"`
}

type ParamsOfNaclBoxOpen struct {
	Encrypted   string `json:"encrypted"`
	Nonce       string `json:"nonce"`
	TheirPublic string `json:"their_public"`
	Secret      string `json:"secret"`
}

type ResultOfNaclBoxOpen struct {
	Decrypted string `json:"decrypted"`
}

type ParamsOfNaclSecretBox struct {
	Decrypted string `json:"decrypted"`
	Nonce     string `json:"nonce"`
	Key       string `json:"key"`
}

type ParamsOfNaclSecretBoxOpen struct {
	Encrypted string `json:"encrypted"`
	Nonce     string `json:"nonce"`
	Key       string `json:"key"`
}

type ParamsOfMnemonicWords struct {
	Dictionary *int `json:"dictionary,omitempty"`
}

type ResultOfMnemonicWords struct {
	Words string `json:"words"`
}

type ParamsOfMnemonicFromRandom struct {
	Dictionary *int `json:"dictionary,omitempty"`
	WordCount  *int `json:"word_count,omitempty"`
}

type ResultOfMnemonicFromRandom struct {
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicFromEntropy struct {
	Entropy    string `json:"entropy"`
	Dictionary *int   `json:"dictionary,omitempty"`
	WordCount  *int   `json:"word_count,omitempty"`
}

type ResultOfMnemonicFromEntropy struct {
	Phrase string `json:"phrase"`
}

type ParamsOfMnemonicVerify struct {
	Phrase     string `json:"phrase"`
	Dictionary *int   `json:"dictionary,omitempty"`
	WordCount  *int   `json:"word_count,omitempty"`
}

type ResultOfMnemonicVerify struct {
	Valid bool `json:"valid"`
}

type ParamsOfMnemonicDeriveSignKeys struct {
	Phrase     string  `json:"phrase"`
	Path       *string `json:"path,omitempty"`
	Dictionary *int    `json:"dictionary,omitempty"`
	WordCount  *int    `json:"word_count,omitempty"`
}

type ParamsOfHDKeyXPrvFromMnemonic struct {
	Phrase string `json:"phrase"`
}

type ResultOfHDKeyXPrvFromMnemonic struct {
	Xprv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrv struct {
	Xprv       string `json:"xprv"`
	ChildIndex int    `json:"child_index"`
	Hardened   bool   `json:"hardened"`
}

type ResultOfHDKeyDeriveFromXPrv struct {
	Xprv string `json:"xprv"`
}

type ParamsOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
	Path string `json:"path"`
}

type ResultOfHDKeyDeriveFromXPrvPath struct {
	Xprv string `json:"xprv"`
}

type ParamsOfHDKeySecretFromXPrv struct {
	Xprv string `json:"xprv"`
}

type ResultOfHDKeySecretFromXPrv struct {
	Secret string `json:"secret"`
}

type ParamsOfHDKeyPublicFromXPrv struct {
	Xprv string `json:"xprv"`
}

type ResultOfHDKeyPublicFromXPrv struct {
	Public string `json:"public"`
}
