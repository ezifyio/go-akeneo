package goakeneo

func MockDLClient() *Client {
	con := Connector{
		ClientID: "1_18ef5k6abydc4og40osc0004cskswgw40gsow0sc00o0oc8880",
		Secret:   "64wgmrelgwsgg0o8k4ckkwsoookkcs0ccg4w00kc8cg8oc0os8",
		UserName: "shopline_1983",
		Password: "9b5468313",
	}
	c, _ := con.NewClient(
		WithBaseURL("http://pim.zdldove.top/"))
	return c
}
