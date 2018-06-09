package badcerts

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

// Fingerprint returns the sha256 of an x509 certificate signature,
// encoded with standard base64.
func Fingerprint(cert *x509.Certificate) string {
	digest := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
	return base64.StdEncoding.EncodeToString(digest[:])
}
