package client

import (
	"github.com/go-resty/resty/v2"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/user"
)

const (
	clientInfoV1 = "/v1/admin/client/get"
)

type API struct {
	User       magic.User
	ClientInfo magic.ClientInfo
}

func New(secret string, client *resty.Client) (*API, error) {
	clientInfo, err := getClientInfo(secret, client)
	if err != nil {
		return nil, err
	}
	return &API{
		User:       user.NewUserClient(secret, clientInfo.ClientId, client),
		ClientInfo: *clientInfo,
	}, nil
}

func getClientInfo(secret string, client *resty.Client) (*magic.ClientInfo, error) {
	meta := new(magic.ClientInfo)
	respData := new(magic.Response)
	respData.Data = meta

	r, err := client.R().
		SetHeader(magic.APISecretHeader, secret).
		SetResult(respData).
		Get(clientInfoV1)
	if err != nil {
		return nil, &magic.APIConnectionError{Err: err}
	}
	if r.IsError() {
		return nil, magic.WrapError(r, r.Error().(*magic.Error))
	}

	return meta, nil
}
