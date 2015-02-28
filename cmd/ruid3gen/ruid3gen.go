package main

import (
	"fmt"

	"github.com/glycerine/ruid"
)

// generate Ruid3 unique id.

func main() {
	myExternalIP := ruid.GetExternalIP()
	ruidGen := ruid.NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", ruidGen.Ruid3())
}
