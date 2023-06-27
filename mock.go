package goakeneo

func MockClient() *Client {
	con := Connector{
		ClientID: "7_5ydf9uurmko4kok8s0cooco80g4gw8w8sc0ckccwcks4kswcwk",
		Secret:   "37vje48chrokww4wooowgk8gs4kwskcokg00w4cowk4ogggcc8",
		UserName: "newbella_ar_7172",
		Password: "8b562df94",
	}
	c, _ := con.NewClient(
		WithBaseURL("https://akeneo.aiogrowth.com/"))
	return c
}

func MockDLClient() *Client {
	con := Connector{
		ClientID: "1_6aoyrxge5ywwwgo80gwkgcs0w8o0kwkw4osskwoko4kcc8o0s8",
		Secret:   "1b8s3ie4hha844g8cwkkcg4480koc84gskow440owsg0s48k0o",
		UserName: "shoplaza_0650",
		Password: "1a49c4586",
	}
	c, _ := con.NewClient(
		WithBaseURL("http://pim.zdldove.top/"))
	return c
}
