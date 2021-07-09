// Package diffiehellman handles encryption using Diffie-Hellman algorithm
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey returns Diffie-Hellman private key
func PrivateKey(p *big.Int) *big.Int {
	min := big.NewInt(2)
	max := new(big.Int).Sub(p, min)
	pk, _ := rand.Int(rand.Reader, max)

	return pk.Add(pk, min)
}

// PublicKey returns the public key for a Diffie-Hellman algorithm
func PublicKey(a, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), a, p)
}

// SecretKey returns the secret key for a Diffie-Hellman algorithm
func SecretKey(a, B, p *big.Int) *big.Int {
	return new(big.Int).Exp(B, a, p)
}

// NewPair returns the pair of Public/Secret Key for a Diffie-Hellman algorithm
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	return a, PublicKey(a, p, g)
}
