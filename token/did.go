package token

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const nbfGracePeriod = 300

type Claim struct {
	Iat int64  `json:"iat"`
	Ext int64  `json:"ext"`
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Nbf int64  `json:"nbf"`
	Tid string `json:"tid"`
	Add string `json:"add,omitempty"`
}

// String returns string data of the claim in json format.
func (c *Claim) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// Token is representation of the DID token which contains proof (hash) and the serialized claim.
type Token struct {
	proof string
	claim Claim
}

// NewToken creates token decoder and validator.
func NewToken(tk string) (*Token, error) {
	token := new(Token)

	decoded, err := base64.URLEncoding.DecodeString(tk)
	if err != nil {
		return nil, err
	}

	// Decode list of proof and serialized in json claim.
	var decodePieces []string
	if err := json.Unmarshal(decoded, &decodePieces); err != nil {
		return nil, err
	}

	token.proof = decodePieces[0]

	if err := json.Unmarshal([]byte(decodePieces[1]), &token.claim); err != nil {
		return nil, err
	}

	return token, err
}

// GetIssuer returns the claim issuer.
func (t *Token) GetIssuer() string {
	return t.claim.Iss
}

// GetPublicAddress split issuer on parts and returns only public address.
func (t *Token) GetPublicAddress() (string, error) {
	parts := strings.Split(t.GetIssuer(), ":")
	if len(parts) != 3 {
		return "", ErrNotValidPublicAddr
	}

	return parts[2], nil
}

// GetProof returns the hash of the Ethereum message with claim serialized in json.
func (t *Token) GetProof() string {
	return t.proof
}

// GetClaim returns the claim structure with all data.
func (t *Token) GetClaim() Claim {
	return t.claim
}

// GetNbfGracePeriod returns nbf time with grace period.
func (t *Token) GetNbfGracePeriod() int64 {
	return t.claim.Nbf - nbfGracePeriod
}

// Validate validates DID token by trying to recover public key using signature data
// and the hash of the claim message.
func (t *Token) Validate() error {
	jsonClaim, err := json.Marshal(t.claim)
	if err != nil {
		return &DIDTokenError{err}
	}
	compactedBuffer := new(bytes.Buffer)
	if err := json.Compact(compactedBuffer, jsonClaim); err != nil {
		return err
	}

	proof, err := hexutil.Decode(t.proof)
	if err != nil {
		return &DIDTokenError{err}
	}
	addr, err := ecRecover(signHash(compactedBuffer.Bytes()).Bytes(), proof)
	if err != nil {
		return &DIDTokenError{err}
	}

	// Validate public address that is matched which is specified in proof and claim.
	pubAddr, err := t.GetPublicAddress()
	if err != nil {
		return &DIDTokenError{err}
	}
	if addr.String() != pubAddr {
		return ErrNotValidProof
	}

	// Check that current token is not expired.
	now := time.Now().Unix()
	if now > t.claim.Ext {
		return ErrExpired
	}
	if now < t.GetNbfGracePeriod() {
		return ErrNbfExpired
	}

	return nil
}

// signHash formats Ethereum signed message and takes keccak256 hash from it.
func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

// ecRecover returns the address for the Account that was used to create the signature.
//
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
//   hash = Keccak256Hash("\x19${byteVersion}Ethereum Signed Message:\n${message length}${message}")
//   addr = ecRecover(hash, signature)
func ecRecover(hash hexutil.Bytes, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27
	rpk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*rpk), nil
}
