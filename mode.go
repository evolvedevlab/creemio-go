package creemio

// Mode represents the environment
type Mode string

const (
	ModeTest       Mode = "test"
	ModeProduction Mode = "prod"
	ModeSandbox    Mode = "sandbox"
)
