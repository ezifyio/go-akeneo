package goakeneo

func MockClient() *Client {
	con := Connector{
		ClientID: "1_3qs6i4sy5hescwkckkkc4c8ogg8s8ss4kcosko0k400wo088ws",
		Secret:   "3lvykhhaw7eo8s4osgo4s44kk8kc88cokgsgw84ss088k0k4s0",
		UserName: "shoplaza_0765",
		Password: "c064bd494",
	}
	c, err := con.NewClient(
		WithBaseURL("https://akeneo.aiogrowth.com/"))
	if err != nil {
		panic(err)
	}
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
