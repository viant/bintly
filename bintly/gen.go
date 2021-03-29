package main

import (
	"github.com/viant/bintly/codegen/cmd"
	"os"
)

var Version string = "1.0"

func main() {
	//args := []string{
	//	"main","-s=/xxx/message.go", "-d=/xxx","-t=Message",
	//}

	args := os.Args
	cmd.RunClient(Version, args)
}
