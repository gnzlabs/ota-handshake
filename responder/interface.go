package responder

import (
	"github.com/gnzlabs/ota-handshake/challenge"
	"github.com/gnzlabs/ota-handshake/response"
)

type Responder interface {
	AuthenticateChallenge(challenge.Signed) error
	NewChallengeResponse(challenge.Signed) (response.Signed, error)
}
