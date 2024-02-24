package authorizer

type yb struct {
}

func (y yb) GetAuthInfo() (*Info, error) {
	// TODO implement me
	panic("implement me")
}

func init() {
	register("yb", &yb{})
}

func (y yb) Name() string {
	return "易班"
}

func (y yb) GetAccessToken(req GetAccessTokenRequest) (*GetAccessTokenReply, error) {
	return nil, nil
}
