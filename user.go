package magic

import (
	"encoding/json"
)

type User interface {
	GetMetadataByIssuer(issuer string) (*Metadata, error)
	GetMetadataByPublicAddress(pubAddr string) (*Metadata, error)
	GetMetadataByToken(didToken string) (*Metadata, error)

	LogoutByIssuer(issuer string) error
	LogoutByPublicAddress(pubAddr string) error
	LogoutByToken(didToken string) error
}

type Metadata struct {
	Email         string `json:"email"`
	Issuer        string `json:"issuer"`
	PublicAddress string `json:"public_address"`
}

func (m *Metadata) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return string(data)
}
