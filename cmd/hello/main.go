package hello

import (
	"fmt"
	
	ghsource "github.com/kpcraig/gha-source"
	ghtarget "github.com/kpcraig/gha-target"
)

func main() {
	greeting := ghtarget.SayHello(ghsource.GetFake())
	fmt.Println(greeting)
}
