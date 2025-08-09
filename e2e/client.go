package e2e

import (
	"os"

	"github.com/evolvedevlab/creemio-go"
	_ "github.com/joho/godotenv/autoload"
)

// Only run tests inside e2e package via individually.

var client = creemio.New(
	creemio.WithBaseURL(creemio.TestAPIURL),
	creemio.WithAPIKey(os.Getenv("API_KEY")),
)
