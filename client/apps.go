package client

type AppDebotBrowser interface {
	ApproveRequest(ApproveParamsOfAppDebotBrowser) (ApproveResultOfAppDebotBrowser, error)
	InputRequest(InputParamsOfAppDebotBrowser) (InputResultOfAppDebotBrowser, error)
	GetSigningBoxRequest(GetSigningBoxParamsOfAppDebotBrowser) (GetSigningBoxResultOfAppDebotBrowser, error)
	InvokeDebotRequest(InvokeDebotParamsOfAppDebotBrowser) (InvokeDebotResultOfAppDebotBrowser, error)

	LogNotify(LogParamsOfAppDebotBrowser)
	SwitchNotify(SwitchParamsOfAppDebotBrowser)
	SwitchCompletedNotify(SwitchCompletedParamsOfAppDebotBrowser)
	ShowActionNotify(ShowActionParamsOfAppDebotBrowser)
	SendNotify(SendParamsOfAppDebotBrowser)
}

type AppSigningBox interface {
	GetPublicKeyRequest(GetPublicKeyParamsOfAppSigningBox) (GetPublicKeyResultOfAppSigningBox, error)
	SignRequest(SignParamsOfAppSigningBox) (SignResultOfAppSigningBox, error)
}

type AppEncryptionBox interface {
	GetInfoRequest(GetInfoParamsOfAppEncryptionBox) (GetInfoResultOfAppEncryptionBox, error)
	EncryptRequest(EncryptParamsOfAppEncryptionBox) (EncryptResultOfAppEncryptionBox, error)
	DecryptRequest(DecryptParamsOfAppEncryptionBox) (DecryptResultOfAppEncryptionBox, error)
}
