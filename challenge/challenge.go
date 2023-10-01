package challenge

import (
	"io"
	"time"

	"github.com/gnzlabs/utilities/convert"
)

type Challenge []byte

func (c Challenge) CreationTime() uint64 {
	return convert.BytesToUInt64(c[0:ChallengeTimestampLength])
}

func (c Challenge) Nonce() []byte {
	return c[ChallengeTimestampLength:ChallengeNonceLength]
}

func New(rand io.Reader) (challenge Challenge, err error) {
	nonce := make([]byte, ChallengeNonceLength)
	if _, err = rand.Read(nonce); err == nil {
		challenge = append(challenge, convert.UInt64ToBytes(uint64(time.Now().Unix()))...)
		challenge = append(challenge, nonce...)
	}
	return
}
