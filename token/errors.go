package token

import "errors"

var (
	ErrNotValidPublicAddr = &DIDTokenError{errors.New("Invalid public address format.")}
	ErrNotValidProof      = &DIDTokenError{errors.New("Signature mismatch between 'proof' and 'claim'. Please " +
		"generate a new token with an intended issuer.")}
	ErrExpired    = &DIDTokenError{errors.New("Given DID token has expired. Please generate a new one.")}
	ErrNbfExpired = &DIDTokenError{errors.New("Given DID token cannot be used at this time. Please " +
		"check the 'nbf' field and regenerate a new token with a suitable value.")}
)

type DIDTokenError struct {
	err error
}

func (e *DIDTokenError) Error() string {
	return e.err.Error()
}
