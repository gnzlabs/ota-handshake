package handshake

type AuthenticationType byte

const (
	Biometric AuthenticationType = iota
	SecurityKey
	KnowledgeCheck
)
