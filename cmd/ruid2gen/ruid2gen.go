package main

import (
	"fmt"

	"github.com/glycerine/ruid"
)

// generate Ruid2 unique id.

func main() {
	myExternalIP := ruid.GetExternalIP()
	ruidGen := ruid.NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", string(ruidGen.Ruid2()))
}
