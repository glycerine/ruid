package main

import (
	"fmt"

	"github.com/glycerine/ruid"
)

// Ruid: a really unique id, Very fast to generate, opaque identifier.
// Huid: a really unique id, very fast to generate, decodable to be human readable.

func main() {
	myExternalIP := ruid.GetExternalIP()
	ruidGen := ruid.NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", string(ruidGen.Tuid()))
}
