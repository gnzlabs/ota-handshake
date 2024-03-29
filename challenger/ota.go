package challenger

type otaChallenger struct {
	configuration *Configuration
}

// New instantiates otaChallenger with the specified configuration.
// It returns a pointer to the newly created instance
func New(c *Configuration) Challenger {
	challenger := otaChallenger{
		configuration: c,
	}
	return &challenger
}
