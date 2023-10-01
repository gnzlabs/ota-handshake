package challenger

import (
	"io"

	"github.com/gnzlabs/identity/certificate"
	"github.com/gnzlabs/keyring"
)

type Configuration struct {

	// Rand is a an instance of a cryptographically secure
	// random number generator
	Rand io.Reader

	// Timeout defines the lifetime of a challenge in seconds
	Timeout uint

	// Biometric is a pointer to the identity certificate
	// authorized to provide biometric authentication
	Biometric certificate.Identity

	// SecurityKey is a pointer to the identity certificate
	// used to authenticate an authorized security key
	SecurityKey certificate.Identity

	// KnowledgeCheck is a pointer to the identity certificate
	// authorized to provide PIN/Password authentication
	KnowledgeCheck certificate.Identity

	// Keyring is a pointer to an instance of keyring.HardwareKeyRing
	// used to authenticate the Challenger. The Keyring instance must
	// be unlocked and ready for use
	Keyring keyring.HardwareKeyRing
}

func (c *Configuration) AuthorizedKeys() (certificates []certificate.Identity) {
	certificates = []certificate.Identity{
		c.Biometric,
		c.SecurityKey,
		c.KnowledgeCheck,
	}
	return
}
