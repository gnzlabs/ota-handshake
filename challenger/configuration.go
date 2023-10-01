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

	// Biometric is a pointer to the hardware-backed certificate
	// authorized to provide biometric authentication
	Biometric *certificate.HardwareBackedCertificate

	// SecurityKey is a pointer to the hardware-backed certificate
	// used to authenticate an authorized security key
	SecurityKey *certificate.HardwareBackedCertificate

	// KnowledgeCheck is a pointer to the hardware-backed certificate
	// authorized to provide PIN/Password authentication
	KnowledgeCheck *certificate.HardwareBackedCertificate

	// Keyring is a pointer to an instance of keyring.HardwareKeyRing
	// used to authenticate the Challenger. The Keyring instance must
	// be unlocked and ready for use
	Keyring *keyring.HardwareKeyRing
}
