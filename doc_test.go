package badcerts_test

import (
	"net/http"

	"github.com/fishy/badcerts"
)

func Example() {
	// This is the cert fingerprint from https://self-signed.badssl.com/
	//
	// You can get it by:
	//
	//     go get -u github.com/fishy/badcerts/cmd/badcerts-fingerprint
	//     badcerts-fingerprint -url https://self-signed.badssl.com/
	myCertFingerprint := "9SLklscvzMYj8f+52lp5ze/hY0CFHyLSPQzSpYYIBm8="

	client := &http.Client{
		Transport: &http.Transport{
			DialTLS: badcerts.DialTLSWithWhitelistCerts(
				badcerts.IsSelfSignedError,
				myCertFingerprint,
			),
		},
	}

	// Now client can handle https://self-signed.badssl.com/ just fine:
	//
	// _, err := client.Get("https://self-signed.badssl.com/")
	// if err != nil {
	//	panic(err)
	// }
	// fmt.Println("Everything is awesome.")
	//
	// And it will still return error for other bad certificates.

	// Satisfy compiler
	_ = client
}
