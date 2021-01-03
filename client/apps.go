package client

type AppDebotBrowser interface {
	Request(ParamsOfAppDebotBrowser) (ResultOfAppDebotBrowser, error)
	Notify(ParamsOfAppDebotBrowser)
}

type AppSigningBox interface {
	Request(ParamsOfAppSigningBox) (ResultOfAppSigningBox, error)
	Notify(ParamsOfAppSigningBox)
}
