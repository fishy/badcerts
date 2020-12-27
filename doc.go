// Package badcerts is a library to handle bad (e.g. self-signed) certificates.
//
// It provides a function you could use in your http.Client,
// to handle the case that you do not want to disable https certificate
// validation altogether,
// but you want to whitelist one (or more) bad (self-signed, expired, wrong
// common name, etc.) cert(s) because you have to. ¯\_(ツ)_/¯
//
// This library is inspired by https://github.com/tam7t/hpkp
package badcerts // import "go.yhsif.com/badcerts"
