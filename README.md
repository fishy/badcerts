[![GoDoc](https://godoc.org/github.com/fishy/badcerts?status.svg)](https://godoc.org/github.com/fishy/badcerts)

# BadCerts

BadCerts is a [Go](https://golang.org) library to deal with bad ssl cert(s)
(e.g. self-signed certificates).

## Example

```go
// This is the cert fingerprint from https://self-signed.badssl.com/
myCertFingerprint := "9SLklscvzMYj8f+52lp5ze/hY0CFHyLSPQzSpYYIBm8="

client := &http.Client{
	Transport: &http.Transport{
		DialTLS: badcerts.DialTLSWithWhitelistCert(
			myCertFingerprint,
			badcerts.IsSelfSignedError,
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

badcerts library still have all the normal certificate verification protections,
it just trust the whitelisted certificate(s) additionally, but nothing more.

### How do I get the fingerprint for my self-signed cert?

It comes with a command line tool
[badcerts-fingerprint](https://godoc.org/github.com/fishy/badcerts/cmd/badcerts-fingerprint).

### Aren't those certs bad?

Yes they are.
You should use [Let's Encrypt](https://letsencrypt.org/) on your site.
This is more for the sites you cannot control and have to deal with.

## Acknowledges

This library is inspired by [tam7t/hpkp](https://github.com/tam7t/hpkp)
