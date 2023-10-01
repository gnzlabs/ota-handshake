package challenger

import (
	"github.com/gnzlabs/ota-handshake/challenge"
	"github.com/gnzlabs/ota-handshake/response"
)

type Challenger interface {
	NewChallenge() (challenge.Signed, error)
	AuthenticateResponse(response.Signed) error
}
