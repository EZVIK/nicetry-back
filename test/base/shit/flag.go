package main

import (
	"flag"
	"fmt"
)

func main() {

	var module string

	flag.StringVar(&module, "module", "NOTHING", "assign run module")

	flag.Parse()

	fmt.Println(fmt.Sprintf("start run %s module", module))

}
