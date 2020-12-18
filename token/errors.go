package token

import "errors"

var (
	ErrNotValidPublicAddr = errors.New("not valid public address format")
	ErrNotValidProof      = errors.New("signature mismatch between 'proof' and 'claim'. Please " +
		"generate a new token with an intended issuer")
	ErrExpired            = errors.New("given DID token has expired. Please generate a new one")
	ErrNbfExpired         = errors.New("given DID token cannot be used at this time. Please " +
		"check the 'nbf' field and regenerate a new token with a suitable value")
)

type DIDTokenError struct {
	err error
}

func (e *DIDTokenError) Error() string {
	return e.err.Error()
}
