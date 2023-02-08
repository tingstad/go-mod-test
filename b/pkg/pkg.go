package pkg

import "fmt"
import "github.com/tingstad/go-mod-test/c/pkg"

func B() string {
	return fmt.Sprintf("got: %s", pkg.Version())
}
