package challenge

type Signed []byte

func (s Signed) Challenge() Challenge {
	return Challenge(s[0:ChallengeLength])
}

func (s Signed) Signature() []byte {
	return s[ChallengeLength:]
}

func NewSigned(c Challenge, signature []byte) (s Signed) {
	s = append(s, c...)
	s = append(s, signature...)
	return
}
