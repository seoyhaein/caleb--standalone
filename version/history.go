package version

import (
	"fmt"

	"github.com/blang/semver/v4"
)

func ToString() string {
	v, err := semver.Parse("0.0.1-alpha.preview")
	if err != nil {
		return fmt.Sprintf("Error while parsing (not valid): %q", err)
	}
	return fmt.Sprintf("Version: %q\n", v)
}
