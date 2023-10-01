package challenge

type Signed []byte

func (s Signed) Challenge() Challenge {
	return Challenge(s[0:Length])
}

func (s Signed) Signature() []byte {
	return s[Length:]
}

func NewSigned(c Challenge, signature []byte) (s Signed) {
	s = append(s, c...)
	s = append(s, signature...)
	return
}
