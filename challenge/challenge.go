package challenge

import (
	"io"
	"time"

	"github.com/gnzlabs/utilities/convert"
)

type Challenge []byte

func (c Challenge) CreationTime() uint64 {
	return convert.BytesToUInt64(c[0:TimestampLength])
}

func (c Challenge) Nonce() []byte {
	return c[TimestampLength:NonceLength]
}

func New(rand io.Reader) (challenge Challenge, err error) {
	nonce := make([]byte, NonceLength)
	if _, err = rand.Read(nonce); err == nil {
		challenge = append(challenge, convert.UInt64ToBytes(uint64(time.Now().Unix()))...)
		challenge = append(challenge, nonce...)
	}
	return
}
