package challenger

import (
	"errors"

	"github.com/gnzlabs/ota-handshake/challenge"
	"github.com/gnzlabs/ota-handshake/response"
)

// NewChallenge implements Challenger for otaChallenger
func (challenger *otaChallenger) NewChallenge() (challenge.Signed, error) {
	return nil, errors.New("not yet implemented")
}

// AuthenticateResponse implements Challenger for otaChallenger
func (challenger *otaChallenger) AuthenticateResponse(response.Signed) error {
	return errors.New("not yet implemented")
}
