package challenge

const (
	ChallengeTimestampLength = 8
	ChallengeNonceLength     = 32
)

var (
	ChallengeLength = ChallengeTimestampLength + ChallengeNonceLength
)
