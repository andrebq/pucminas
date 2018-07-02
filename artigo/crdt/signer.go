package crdt

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"golang.org/x/crypto/nacl/sign"
)

type (
	signer struct {
		identity Identity
		privKey  *[64]byte
	}
)

// NewSigner returns a new signer using a random generated
// key pair and a random generated ID
func NewSigner() (Signer, error) {
	pubKey, privKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		errors.Wrapf(err, "crdt:signer: unable to generate key pair")
	}
	return &signer{
		identity: Identity{
			SignerID:  xid.New().String(),
			PublicKey: PublicKey(hex.EncodeToString(pubKey[:])),
		},
		privKey: privKey,
	}, nil
}

// Identity implements Signer interface
func (s *signer) Identity() Identity {
	return s.identity
}

// Sign implements Signer interface
func (s *signer) Sign(in []byte) ([]byte, error) {
	return sign.Sign(nil, in, s.privKey), nil
}
