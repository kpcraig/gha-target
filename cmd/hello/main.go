package hello

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

func main() {
	fmt.Printf("hello %s!\n", getFake())
}

// contrive a dependency
func getFake() string {
	return faker.FirstName()
}
