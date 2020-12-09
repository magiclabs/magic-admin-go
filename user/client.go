package user

import (
	"fmt"
	"github.com/go-resty/resty/v2"

	"github.com/magiclabs/magic-admin-go"
	"github.com/magiclabs/magic-admin-go/token"
)

const (
	userInfoV1 = "/v1/admin/auth/user/get"
	userLogoutV2 = "/v2/admin/auth/user/logout"
)

type Client struct {
	secret string
	client *resty.Client
}

// NewUserClient constructor of user client api.
func NewUserClient(secret string, client *resty.Client) magic.User {
	return &Client{
		secret: secret,
		client: client,
	}
}

// GetMetadataByIssuer returns metadata by issuer.
func (u *Client) GetMetadataByIssuer(issuer string) (*magic.Metadata, error) {
	meta := new(magic.Metadata)
	respData := new(magic.Response)
	respData.Data = meta

	r, err := u.client.R().
		SetQueryParams(map[string]string{"issuer": issuer}).
		SetHeader(magic.APISecretHeader, u.secret).
		SetResult(respData).
		Get(userInfoV1)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, r.Error().(error)
	}

	return meta, nil
}

// GetMetadataByPublicAddress returns metadata by public address.
func (u *Client) GetMetadataByPublicAddress(pubAddr string) (*magic.Metadata, error) {
	return u.GetMetadataByIssuer(fmt.Sprintf("did:ethr:%s", pubAddr))
}

// GetMetadataByToken returns metadata by DID token with decoding and validating it.
func (u *Client) GetMetadataByToken(didToken string) (*magic.Metadata, error) {
	tk, err := token.NewToken(didToken)
	if err != nil {
		return nil, err
	}
	if err := tk.Validate(); err != nil {
		return nil, err
	}

	return u.GetMetadataByIssuer(tk.GetIssuer())
}

// LogoutByIssuer logout user from magic.link service by issuer.
func (u *Client) LogoutByIssuer(issuer string) error {
	r, err := u.client.R().
		SetBody(map[string]interface{}{"issuer": issuer}).
		SetHeader(magic.APISecretHeader, u.secret).
		Post(userLogoutV2)
	if err != nil {
		return err
	}
	if r.IsError() {
		return r.Error().(error)
	}

	return nil
}

// LogoutByPublicAddress logout user from magic.link service by public address.
func (u *Client) LogoutByPublicAddress(pubAddr string) error {
	return u.LogoutByIssuer(fmt.Sprintf("did:ethr:%s", pubAddr))
}

// LogoutByToken logout user from magic.link service by DID token with decoding and validating it.
func (u *Client) LogoutByToken(didToken string) error {
	tk, err := token.NewToken(didToken)
	if err != nil {
		return err
	}
	if err := tk.Validate(); err != nil {
		return err
	}

	return u.LogoutByIssuer(tk.GetIssuer())
}
