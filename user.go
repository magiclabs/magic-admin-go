package magic

import (
	"encoding/json"

	"github.com/magiclabs/magic-admin-go/wallet"
)

type User interface {
	GetMetadataByIssuer(issuer string) (*UserInfo, error)
	GetMetadataByPublicAddress(pubAddr string) (*UserInfo, error)
	GetMetadataByToken(didToken string) (*UserInfo, error)

	GetMetadataByIssuerAndWallet(issuer string, walletType wallet.Type) (*UserInfo, error)
	GetMetadataByPublicAddressAndWallet(pubAddr string, walletType wallet.Type) (*UserInfo, error)
	GetMetadataByTokenAndWallet(didToken string, walletType wallet.Type) (*UserInfo, error)

	LogoutByIssuer(issuer string) error
	LogoutByPublicAddress(pubAddr string) error
	LogoutByToken(didToken string) error
}

type UserInfo struct {
	Email         string    `json:"email"`
	Issuer        string    `json:"issuer"`
	PublicAddress string    `json:"public_address"`
	Wallets       *[]Wallet `json:"wallets,omitempty"`
}

type ClientInfo struct {
	ClientId string `json:"client_id"`
	AppScope string `json:"app_scope"`
}

type Wallet struct {
	Network       string `json:"network"`
	PublicAddress string `json:"public_address"`
	Type          string `json:"wallet_type"`
}

func (m *UserInfo) String() string {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return string(data)
}
