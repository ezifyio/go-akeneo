package goakeneo

func MockClient() *Client {
	con := Connector{
		ClientID: "4_pnospzctqo04oock4kw0co84c404gkw08kckwsk0c00w00008",
		Secret:   "2167brjppc80gs8c8k4k0k0owsgok48gc4ogkggw0gks4wgss4",
		UserName: "shoplaza_8153",
		Password: "9088446db",
	}
	c, _ := con.NewClient(
		WithBaseURL("http://pim.zdldove.top/"))
	return c
}
