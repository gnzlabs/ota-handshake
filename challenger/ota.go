package challenger

type otaChallenger struct {
	configuration *ChallengerConfiguration
}

// New instantiates otaChallenger with the specified configuration.
// It returns a pointer to the newly created instance
func New(c *ChallengerConfiguration) Challenger {
	challenger := otaChallenger{
		configuration: c,
	}
	return &challenger
}
