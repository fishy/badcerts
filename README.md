[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/badcerts)](https://pkg.go.dev/go.yhsif.com/badcerts)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/badcerts)](https://goreportcard.com/report/go.yhsif.com/badcerts)

# BadCerts

BadCerts is a [Go](https://golang.org) library to deal with bad ssl cert(s)
(e.g. self-signed certificates).

**NOTE**: For self-signed certs, a better approach to deal with them is to use
[`x509.CertPool`](https://pkg.go.dev/crypto/x509#CertPool),
which is faster than using BadCerts library.
An example can be found
[here](https://github.com/fishy/blynk-proxy/blob/741ab221c0624d8b522428f7ac0958584c6d2a1f/main.go#L59-L77)
(Thanks to
[/u/loosecanonsandvich](https://www.reddit.com/r/golang/comments/8prc19/a_go_library_to_deal_with_bad_https_certs/e0dmnzp/)).
BadCerts library is still kinda useful to deal with other types of bad certs,
like expired certs or certs with wrong common names.

## Example

```go
// This is the cert fingerprint from https://self-signed.badssl.com/
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

_, err := client.Get("https://self-signed.badssl.com/")
if err != nil {
	panic(err)
}
fmt.Println("Everything is awesome.")

// And it will still return error for other bad certificates.
```

## FAQs

### But I could just disable certificate verification?

Yes you can,
but then you won't know if it's replaced by a different, malicious MITM cert.
Or you could use the same http client with sites with legit certs and now you
are losing protection.

BadCerts library still have all the normal certificate verification protections,
it just trust the whitelisted certificate(s) additionally, but nothing more.

### How do I get the fingerprint for my self-signed cert?

It comes with a command line tool
[badcerts-fingerprint](https://pkg.go.dev/go.yhsif.com/badcerts/cmd/badcerts-fingerprint).

### Aren't those certs bad?

Yes they are.
You should use [Let's Encrypt](https://letsencrypt.org/) on your site.
This is more for the sites you cannot control and have to deal with.

## Acknowledges

This library is inspired by [tam7t/hpkp](https://github.com/tam7t/hpkp)

## License

[BSD 3-Clause](LICENSE).
