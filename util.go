package creemio

import (
	"fmt"
	"strings"
)

func makeUrl(baseURL, path string, params ...string) string {
	joined := path
	if len(params) > 0 {
		joined = strings.TrimSuffix(path, "/") + "/" + strings.Join(params, "/")
	}
	if !strings.HasPrefix(joined, "/") {
		joined = "/" + joined
	}

	return fmt.Sprintf("%s/%s%s", strings.TrimSuffix(baseURL, "/"), APIVersion, joined)
}
