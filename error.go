package badcerts

import (
	"crypto/x509"
)

// ErrorFunc defines the function to determine whether an error is cert error
// that should retry to check against whitelisted cert fingerprints.
//
// It should return true to retry the error.
type ErrorFunc func(err error) bool

// IsSelfSignedError is an ErrorFunc returns true for self-signed certs.
func IsSelfSignedError(err error) bool {
	_, ok := err.(x509.UnknownAuthorityError)
	return ok
}
