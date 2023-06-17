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
