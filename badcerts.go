package badcerts

import (
	"crypto/tls"
	"net"
)

// DialTLSWithWhitelistCerts returns a DialTLS implementation.
//
// First it tries standard tls.Dial.
// If nothing is wrong, it returns the result directly.
//
// If the error satisfies errorFunc, it dials again without cert verification,
// then checks the fingerprint of the cert against the given certs.
// If the fingerprint matches it returns the connection without error,
// otherwise it returns the original self-signed cert error.
//
// As a result this function works with all the standard trusted root CAs plus
// the ones with matching cert fingerprint, and nothing else.
func DialTLSWithWhitelistCerts(
	certFingerprints []string,
	errorFunc ErrorFunc,
) func(network, addr string) (net.Conn, error) {
	fingerprints := make(map[string]bool)
	for _, fingerprint := range certFingerprints {
		fingerprints[fingerprint] = true
	}

	return func(network, addr string) (net.Conn, error) {
		conn, err := tls.Dial(network, addr, &tls.Config{})
		if err == nil {
			return conn, nil
		}
		if err != nil {
			if !errorFunc(err) {
				return conn, err
			}
		}

		newConn, newErr := tls.Dial(
			network,
			addr,
			&tls.Config{
				InsecureSkipVerify: true,
			},
		)
		if newErr != nil {
			return newConn, newErr
		}

		for _, cert := range newConn.ConnectionState().PeerCertificates {
			if fingerprints[Fingerprint(cert)] {
				return newConn, nil
			}
		}

		return conn, err
	}
}

// DialTLSWithWhitelistCert is the singular form of DialTLSWithWhitelistCerts.
func DialTLSWithWhitelistCert(
	certFingerprint string,
	errorFunc ErrorFunc,
) func(network, addr string) (net.Conn, error) {
	return DialTLSWithWhitelistCerts([]string{certFingerprint}, errorFunc)
}
