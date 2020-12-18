package magic

import (
	"encoding/json"
)

type User interface {
	GetMetadataByIssuer(issuer string) (*UserInfo, error)
	GetMetadataByPublicAddress(pubAddr string) (*UserInfo, error)
	GetMetadataByToken(didToken string) (*UserInfo, error)

	LogoutByIssuer(issuer string) error
	LogoutByPublicAddress(pubAddr string) error
	LogoutByToken(didToken string) error
}

type UserInfo struct {
	Email         string `json:"email"`
	Issuer        string `json:"issuer"`
	PublicAddress string `json:"public_address"`
}

func (m *UserInfo) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return string(data)
}
