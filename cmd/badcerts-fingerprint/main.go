package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/fishy/badcerts"
)

var url = flag.String("url", "", "url to get cert fingerprint(s)")

func dialer(network, addr string) (net.Conn, error) {
	conn, err := tls.Dial(
		network,
		addr,
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)
	if err != nil {
		return conn, err
	}

	for i, cert := range conn.ConnectionState().PeerCertificates {
		fmt.Printf("cert #%d: %q\n", i, badcerts.Fingerprint(cert))
	}

	return conn, err
}

func main() {
	flag.Parse()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Don't follow any redirects
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			DialTLS: dialer,
		},
	}

	_, err := client.Head(*url)
	if err != nil {
		panic(err)
	}
}
