package apiproto

import "context"

type KeyAuth struct {
	Key string
}

func (t KeyAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "apikey " + t.Key,
	}, nil
}

func (t KeyAuth) RequireTransportSecurity() bool {
	return false
}
