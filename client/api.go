package client

import (
	"github.com/go-resty/resty/v2"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/user"
)

type API struct {
	User magic.User
}

func New(secret string, client *resty.Client) *API {
	return &API{
		User: user.NewUserClient(secret, client),
	}
}
