package bintly

import (
	"github.com/viant/bintly/codegen/cmd"
)

var Version string = "1.0"

func main() {
	args := []string{
		"main","-s=/Users/ppoudyal/go/src/github.com/viant/bintly/codegen/test_data/maps/message.go", "-d=/Users/ppoudyal/go/src/github.com/viant/bintly/codegen/test_data/maps","-t=Message",
	}

//	args := os.Args
	cmd.RunClient(Version, args)
}
