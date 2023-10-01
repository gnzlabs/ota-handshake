package handshake

import "github.com/gnzlabs/identity"

type OtaHandshake interface {
	Authorize(factor AuthenticationType, certificates identity.CertificateManager) error
	Authenticate(factor AuthenticationType, challenge Challenge, response Response) error
}
