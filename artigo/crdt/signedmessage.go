package crdt

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"

	"github.com/pkg/errors"

	"golang.org/x/crypto/nacl/sign"
)

type (
	// SignedMessage contains a message that was signed using the given Identity
	SignedMessage struct {
		Identity Identity

		// Signed message contains the signature + message
		Message []byte
	}

	// Identity contains a public-key + signer id pair to identify any
	// signed message.
	//
	// The receiver is resposible to verify if the given signer is authorized
	// to present the given public-key.
	Identity struct {
		// SignerID is self-explanatory
		SignerID string

		// PublicKey is self-explanatory
		PublicKey PublicKey
	}

	// Signer signs messages
	Signer interface {
		// Public identity associated with the private identity
		Identity() Identity

		// Sign signs in bytes and return the signed message (using nacl/sig)
		Sign(in []byte) ([]byte, error)
	}
)

var (
	errInvalidKey      = errors.New("crdt:signedmsg: invalid key")
	errInvalidMessage  = errors.New("crdt:signedmsg: invalid msg")
	errInvalidIdentity = errors.New("crdt:signedmsg: invalid identity")
)

// IsInvalidKey returns true if err indicates that a key is invalid
func IsInvalidKey(err error) bool {
	return err == errInvalidKey
}

// IsInvalidMessage returns true if err indicates that a message is invalid,
func IsInvalidMessage(err error) bool {
	return err == errInvalidMessage
}

// IsInvalidIdentity indicates that a given identity is not present in the catalog
func IsInvalidIdentity(err error) bool {
	return err == errInvalidIdentity
}

// EncodeAndSign gob-encodes in and signs it using the given Signer
func EncodeAndSign(in interface{}, signer Signer) (SignedMessage, error) {
	buf, err := gobEncode(in)
	if err != nil {
		return SignedMessage{}, err
	}
	buf, err = signer.Sign(buf)
	if err != nil {
		return SignedMessage{}, errors.Wrapf(err, "crdt:signedmsg: signer error")
	}

	return SignedMessage{
		Identity: signer.Identity(),
		Message:  buf,
	}, nil
}

// VerifyAndDecode gob-decodes to "out" the bytes from the given message
// only if the identity is valid.
func VerifyAndDecode(out interface{}, catalog *Catalog, sm *SignedMessage) error {
	len := hex.DecodedLen(len(sm.Identity.PublicKey))
	if len != 32 {
		return errInvalidKey
	}

	if !catalog.VerifyIdentity(sm.Identity) {
		return errInvalidIdentity
	}

	var pubKey [32]byte
	_, err := hex.Decode(pubKey[:], []byte(sm.Identity.PublicKey))
	if err != nil {
		return err
	}

	buf, ok := sign.Open(nil, sm.Message, &pubKey)
	if !ok {
		return errInvalidMessage
	}

	return gobDecodeTo(out, buf)
}

func gobDecodeTo(out interface{}, buf []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	return dec.Decode(out)
}

func gobEncode(in interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(in)
	return buf.Bytes(), err
}
