// Package badcerts is a library to handle bad (e.g. self-signed) certificates.
//
// It provides a function you could use in your http.Client, to handle the case
// that you do not want to disable https certificate altogether, but you want to
// whitelist one (or more) self-signed cert(s) because you have to. ¯\_(ツ)_/¯
//
// This library is inspired by https://github.com/tam7t/hpkp
package badcerts
