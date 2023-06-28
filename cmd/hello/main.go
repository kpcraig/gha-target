package hello

import (
	"fmt"

	ghsource "github.com/kpcraig/gha-source"
)

func main() {
	fmt.Printf("hello %s!\n", ghsource.GetFake())
}
