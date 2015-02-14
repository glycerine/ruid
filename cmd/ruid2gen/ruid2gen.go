package main

import (
	"fmt"
	"github.com/glycerine/ruid"
	"net"
	"regexp"
)

// generate Ruid2 unique id.

func main() {
	myExternalIP := GetExternalIP()
	ruidGen := ruid.NewRuidGen(myExternalIP)
	fmt.Printf("%s\n", string(ruidGen.Ruid2()))
}

func GetExternalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	valid := []string{}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			addr := ipnet.IP.String()
			match := validIPv4addr.FindStringSubmatch(addr)
			if match != nil {
				if addr != "127.0.0.1" {
					valid = append(valid, addr)
				}
			}
		}
	}
	switch len(valid) {
	case 0:
		return "127.0.0.1"
	default:
		return valid[0]
	}
}

var validIPv4addr = regexp.MustCompile(`^[0-9]+[.][0-9]+[.][0-9]+[.][0-9]+$`)
