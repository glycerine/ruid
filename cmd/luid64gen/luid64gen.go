package main

import (
	"fmt"

	"github.com/glycerine/ruid"
)

func main() {
	myExternalIP := ruid.GetExternalIP()
	ruidGen := ruid.NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", string(ruidGen.Luid64()))
}
