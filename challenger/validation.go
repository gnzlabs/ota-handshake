package challenger

import (
	"crypto/x509"

	"github.com/gnzlabs/identity/certificate"
	"github.com/gnzlabs/identity/errors"
	"github.com/gnzlabs/keyring"
	"github.com/gnzlabs/ota-handshake/challenge"
)

const (
	// TODO: Support more key types
	Ed25519SignatureLength int = 64
)

func (challenger *otaChallenger) getSignatureLength(cert *x509.Certificate) (signatureLength int, err error) {
	switch cert.SignatureAlgorithm {
	case x509.PureEd25519:
		signatureLength = Ed25519SignatureLength
	default:
		err = errors.ErrUnsupportedSigScheme
	}
	return
}

func (challenger *otaChallenger) getKeyringSignatureLength(slot keyring.KeySlot) (signatureLength int, err error) {
	if cert, e := challenger.configuration.Keyring.GetCertificate(slot); e == nil {
		signatureLength, err = challenger.getSignatureLength(cert)
	} else {
		err = e
	}
	return
}

func (challenger *otaChallenger) getAuthSignatureLength(cert certificate.Identity) (signatureLength int, err error) {
	if auth, e := cert.AuthenticationCertificate(); e == nil {
		signatureLength, err = challenger.getSignatureLength(auth.Certificate())
	} else {
		err = e
	}
	return
}

func (challenger *otaChallenger) getExpectedResponseSize() (responseSize int, err error) {
	if responseSize, err = challenger.getKeyringSignatureLength(keyring.AuthenticationKeySlot); err == nil {
		responseSize += challenge.Length
		for _, cert := range challenger.configuration.AuthorizedKeys() {
			if signatureLength, err := challenger.getAuthSignatureLength(cert); err == nil {
				responseSize += signatureLength
			} else {
				break
			}
		}
	}
	return
}
